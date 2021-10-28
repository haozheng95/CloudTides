package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"strconv"
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

	//create sshp config
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            body.SSHUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(body.SSHPassword)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)
	session := getSession(body.SSHHost, int(body.SSHPort), sshConfig)

	containerName := "jupyterlab-" + token
	containerPort := body.Port
	containerToken := token
	cmd := fmt.Sprintf("docker run -p %s:8888 -e JUPYTER_ENABLE_LAB=yes --rm -d --name %s jupyter/all-spark-notebook start-notebook.sh --NotebookApp.token='%s'", containerPort, containerName, containerToken)

	//run remote shell
	combo, err := session.CombinedOutput(cmd)

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

	//create sshp config
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

	port, err := strconv.Atoi(data.SshPort)
	if err != nil {
		log.Println("strconv error", err)
	}
	session := getSession(data.SshHost, port, sshConfig)
	cmd := fmt.Sprintf("docker kill %s", data.ContainerID[:12])
	log.Println("cmd ", cmd)
	combo, err := session.CombinedOutput(cmd)

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
	log.Println("data", data)

	payload := make([]*application.ListApplicationInstanceOKBodyItems0, len(data), len(data))
	for i := range data {
		tmp := &application.ListApplicationInstanceOKBodyItems0{
			Link:         data[i].Link,
			Token:        data[i].Token,
			InstanceName: data[i].InstanceName,
			AppType:      data[i].AppType,
		}
		payload[i] = tmp
	}
	return application.NewListApplicationInstanceOK().WithPayload(payload)
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
