package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/application"
	"time"
)

var sshPool sync.Map
var tokenIndex sync.Map
var (
	sshType    = "password" //password or key
	sshKeyPath = ""         //ssh id_rsa.id path"
)

func CreateApplicationInstance(params application.CreateApplicationInstanceParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)

	body := params.ReqBody

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)
	containerName := "jupyterlab-" + token
	containerPort := body.Port
	containerToken := token
	cmd := fmt.Sprintf("docker run -p %s:8888 -e JUPYTER_ENABLE_LAB=yes --rm -d --name %s jupyter/all-spark-notebook start-notebook.sh --NotebookApp.token='%s'", containerPort, containerName, containerToken)

	//run remote shell
	combo, err := execCmd(body.SSHHost, int(body.SSHPort), &models.Application{
		SshPassword: body.SSHPassword,
		SshUser:     body.SSHUser,
	}, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewCreateApplicationInstanceForbidden().WithPayload(&application.CreateApplicationInstanceForbiddenBody{
			Msg: err.Error(),
		})
	}
	log.Println("cmd output:", string(combo))
	sshPortStr := strconv.Itoa(int(body.SSHPort))
	row := map[string]string{
		"link":          fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		"containerID":   string(combo),
		"containerName": containerName,
		"instanceName":  body.InstanceName,
		"token":         token,
		"uid":           strconv.Itoa(int(uid)),
		"sshHost":       body.SSHHost,
		"sshUser":       body.SSHUser,
		"sshPassword":   body.SSHPassword,
		"sshType":       sshType,
		"sshKeyPath":    sshKeyPath,
		"sshPort":       sshPortStr,
		"appType":       body.AppType,
		"port":          containerPort,
	}

	var data models.Application

	if err := mapstructure.Decode(row, &data); err != nil {
		log.Println("map to struct error ", err)
	}
	db := config.GetDB()
	if err := db.Create(&data).Error; err != nil {
		log.Println("db install error", err)
	}
	log.Println("row Id ", data.ID)
	tokenIndex.Store(token, data.ID)

	result := &application.CreateApplicationInstanceOKBody{
		Link:  fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		Token: token,
	}
	return application.NewCreateApplicationInstanceOK().WithPayload(result)
}

func DeleteApplicationInstance(params application.DeleteApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}

	indexStock, has := tokenIndex.Load(params.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", params.Token).First(data)
	}

	port, err := strconv.Atoi(data.SshPort)
	if err != nil {
		log.Println("strconv error", err)
	}
	cmd := fmt.Sprintf("docker kill %s", data.ContainerID[:12])
	log.Println("cmd ", cmd)
	combo, err := execCmd(data.SshHost, port, data, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewDeleteApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	log.Println("index :", index)

	tokenIndex.Delete(params.Token)
	db.Delete(data)

	return application.NewDeleteApplicationInstanceOK().WithPayload(&application.DeleteApplicationInstanceOKBody{
		Msg: "success",
	})
}

func ListApplicationInstance(params application.ListApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var data []models.Application
	db.Where("uid = ?", fmt.Sprintf("%d", uid)).Find(&data)

	payload := make([]*application.ListApplicationInstanceOKBodyItems0, len(data), len(data))
	for i := range data {
		tmp := &application.ListApplicationInstanceOKBodyItems0{
			Link:         data[i].Link,
			Token:        data[i].Token,
			InstanceName: data[i].InstanceName,
			AppType:      data[i].AppType,
			SSHHost:      data[i].SshHost,
			SSHPassword:  data[i].SshPassword,
			SSHPort:      data[i].SshPort,
			SSHUser:      data[i].SshUser,
			Port:         data[i].Port,
		}
		payload[i] = tmp
	}
	return application.NewListApplicationInstanceOK().WithPayload(payload)
}

func UpdateApplicationInstance(params application.UpdateApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	indexStock, has := tokenIndex.Load(body.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", body.Token).First(data)
	}

	cmd := fmt.Sprintf("docker kill %s", data.ContainerID[:12])
	log.Println("cmd ", cmd)
	port, err := strconv.Atoi(data.SshPort)
	combo, err := execCmd(data.SshHost, port, data, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewUpdateApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	log.Println("index :", index)

	tokenIndex.Delete(body.Token)
	db.Delete(data)

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)
	containerName := "jupyterlab-" + token
	containerPort := body.Port
	containerToken := token
	cmd = fmt.Sprintf("docker run -p %s:8888 -e JUPYTER_ENABLE_LAB=yes --rm -d --name %s jupyter/all-spark-notebook start-notebook.sh --NotebookApp.token='%s'", containerPort, containerName, containerToken)

	//run remote shell
	combo, err = execCmd(body.SSHHost, int(body.SSHPort), &models.Application{
		SshPassword: body.SSHPassword,
		SshUser:     body.SSHUser,
	}, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewUpdateApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	sshPortStr := strconv.Itoa(int(body.SSHPort))
	row := map[string]string{
		"link":          fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		"containerID":   string(combo),
		"containerName": containerName,
		"instanceName":  body.InstanceName,
		"token":         token,
		"uid":           strconv.Itoa(int(uid)),
		"sshHost":       body.SSHHost,
		"sshUser":       body.SSHUser,
		"sshPassword":   body.SSHPassword,
		"sshType":       sshType,
		"sshKeyPath":    sshKeyPath,
		"sshPort":       sshPortStr,
		"appType":       body.AppType,
		"port":          containerPort,
	}
	data = new(models.Application)
	if err := mapstructure.Decode(row, &data); err != nil {
		log.Println("map to struct error ", err)
	}
	if err := db.Create(&data).Error; err != nil {
		log.Println("db install error", err)
	}
	log.Println("row Id ", data.ID)
	tokenIndex.Store(token, data.ID)

	result := &application.UpdateApplicationInstanceOKBody{
		Link:  fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		Token: token,
	}
	return application.NewUpdateApplicationInstanceOK().WithPayload(result)

}

func WatchApplicationInstanceLogs(params application.WatchApplicationInstanceLogsParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}

	indexStock, has := tokenIndex.Load(params.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", params.Token).First(data)
	}

	cmd := fmt.Sprintf("docker logs -n 200 %s", data.ContainerID[:12])
	port, _ := strconv.Atoi(data.SshPort)
	//run remote shell
	combo, err := execCmd(data.SshHost, port, data, cmd)
	if err != nil {
		log.Println("run remote shell error", err)
	}
	texts := strings.Split(string(combo), "\n")
	if len(texts) > 0 {
		texts = texts[:len(texts)-1]
	}
	payload := make([]*application.WatchApplicationInstanceLogsOKBodyItems0, len(texts), len(texts))
	for i, line := range texts {
		section := strings.Split(line, "] ")
		payload[i] = new(application.WatchApplicationInstanceLogsOKBodyItems0)
		if len(section) > 1 {
			payload[i].Content = section[1]
			row := strings.Split(section[0], " ")
			payload[i].Leve = strings.Replace(row[0], "[", "", 1)
			payload[i].Date = row[1]
			payload[i].Time = row[2]
			payload[i].Source = row[3]
		} else {
			payload[i].Content = section[0]
			payload[i].Leve = "None"
			payload[i].Date = "None"
			payload[i].Time = "None"
			payload[i].Source = "None"
		}

	}
	return application.NewWatchApplicationInstanceLogsOK().WithPayload(payload)
}

type WsHandler struct {
	Request *http.Request
}

func (ws WsHandler) WriteResponse(w http.ResponseWriter, r runtime.Producer) {

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(w, ws.Request, nil)
	defer conn.Close()
	if err != nil {
		log.Println("websocket error ", err)
	}
	for {
		data := `
 {
        "content": "/usr/local/bin/start-notebook.sh: running hooks in /usr/local/bin/before-notebook.d",
        "date": "None",
        "leve": "None",
        "source": "None",
        "time": "None"
    }
`
		t := rand.Int63n(3)
		log.Println(t)
		time.Sleep(time.Second * time.Duration(t))

		err = conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("websocket write error ", err)
		}
	}
}

func WsWatchApplicationInstanceLogs(params application.WsWatchApplicationInstanceLogsParams) middleware.Responder {
	wsStruct := &WsHandler{
		Request: params.HTTPRequest,
	}
	return wsStruct
}

func execCmd(sshHost string, sshPort int, data *models.Application, cmd string) ([]byte, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            data.SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshType == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(data.SshPassword)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(data.SshPassword)}
	}

	session := getSession(sshHost, sshPort, sshConfig)
	log.Println("cmd ", cmd)
	return session.CombinedOutput(cmd)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {

	key, err := ioutil.ReadFile(kPath)
	if err != nil {
		log.Println("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func getSession(sshHost string, sshPort int, config *ssh.ClientConfig) (session *ssh.Session) {

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	log.Println("ssh addr ", addr)
	sshClientStock, has := sshPool.Load(addr)
	var sshClient *ssh.Client

	if !has {
		sshClient, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Println("created ssh client failed", err)
			if err := sshClient.Close(); err != nil {
				log.Println("close ssh client failed", err)
			}
		} else {
			log.Println("existing sshClient")
			sshPool.Store(addr, sshClient)
		}
	} else {
		log.Println("map existing sshClient")
		sshClient = sshClientStock.(*ssh.Client)
	}

	defer func(addr string) {
		err := recover()
		if err != nil {
			log.Println("recover ", err)
			sshClient, err := ssh.Dial("tcp", addr, config)
			log.Println("defer", err)
			log.Println("defer existing sshClient")
			sshPool.Delete(addr)
			sshPool.Store(addr, sshClient)
			session, err = sshClient.NewSession()
			log.Println("defer", err)
		}
	}(addr)

	//create ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("created session failed", err)
		if err := session.Close(); err != nil {
			log.Println("close session failed", err)
		}
	}

	return session
}
