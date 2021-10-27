package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"tides-server/pkg/restapi/operations/application"
	"time"
)

var sshPool sync.Map
var tokenPool sync.Map
var tokenIndex sync.Map
var (
	sshHost     = "120.133.15.12"
	sshUser     = "root"
	sshPassword = "ca$hc0w"
	sshType     = "password" //password or key
	sshKeyPath  = ""         //ssh id_rsa.id path"
	sshPort     = 20023
)

func CreateApplicationInstance(params application.CreateApplicationInstanceParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody

	//create sshp config
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)
	session := getSession(sshHost, sshPort, config)

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
	sshPortStr := strconv.Itoa(sshPort)
	row := map[string]string{
		"link":          fmt.Sprintf("%s:%s/lab?token=%s", sshHost, containerPort, token),
		"containerID":   string(combo),
		"containerName": containerName,
		"instanceName":  body.InstanceName,
		"token":         token,
		"uid":           string(uid),
		"sshHost":       sshHost,
		"sshUser":       sshUser,
		"sshPassword":   sshPassword,
		"sshType":       sshType,
		"sshKeyPath":    sshKeyPath,
		"sshPort":       sshPortStr,
	}

	var col []map[string]string
	colStock, has := tokenPool.Load(uid)
	if has {
		col = colStock.([]map[string]string)
		tokenIndex.Store(token, len(col))
		col = append(col, row)
		tokenPool.Store(uid, col)
	} else {
		col = []map[string]string{row}
		tokenPool.Store(uid, col)
		tokenIndex.Store(token, 0)
	}

	result := &application.CreateApplicationInstanceOKBody{
		Link:  fmt.Sprintf("%s:%s/lab?token=%s", sshHost, containerPort, token),
		Token: token,
	}
	return application.NewCreateApplicationInstanceOK().WithPayload(result)
}

func DeleteApplicationInstance(params application.DeleteApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)

	colStock, has := tokenPool.Load(uid)
	if !has {
		return application.NewDeleteApplicationInstanceForbidden()
	}

	colList := colStock.([]map[string]string)

	indexStock, has := tokenIndex.Load(params.Token)

	if !has {
		return application.NewDeleteApplicationInstanceForbidden()
	}
	index := indexStock.(int)
	row := colList[index]

	//create sshp config
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            row["sshUser"],
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	log.Println(row)

	port, err := strconv.Atoi(row["sshPort"])
	if err != nil {
		log.Println("strconv error", err)
	}
	session := getSession(row["sshHost"], port, config)
	cmd := fmt.Sprintf("docker kill %s", row["containerID"][:12])
	log.Println("cmd ", cmd)
	combo, err := session.CombinedOutput(cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewDeleteApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	log.Println("index :", index)
	if index == 0 {
		log.Println("len colList", len(colList))
		if len(colList) > 1 {
			colList = colList[1:]
		} else {
			colList = make([]map[string]string, 0, 0)
		}
	} else if index+1 == len(colList) {
		colList = colList[:index]
	}
	if len(colList) != 0 {
		colList = append(colList[:index], colList[index+1:]...)
	}
	if len(colList) != 0 {
		tokenPool.Store(uid, colList)
	} else {
		tokenPool.Delete(uid)
	}
	tokenIndex.Delete(params.Token)

	token := params.Token
	log.Println("will delete token", token)

	return application.NewDeleteApplicationInstanceOK().WithPayload(&application.DeleteApplicationInstanceOKBody{
		Msg: "success",
	})
}

func ListApplicationInstance(params application.ListApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	colStock, has := tokenPool.Load(uid)
	if !has {
		return application.NewListApplicationInstanceOK()
	}
	colList := colStock.([]map[string]string)
	payload := make([]*application.ListApplicationInstanceOKBodyItems0, len(colList), len(colList))
	for i := range colList {
		tmp := &application.ListApplicationInstanceOKBodyItems0{
			Link:         colList[i]["link"],
			Token:        colList[i]["token"],
			InstanceName: colList[i]["instanceName"],
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

func getSession(sshHost string, sshPort int, config *ssh.ClientConfig) *ssh.Session {

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
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
