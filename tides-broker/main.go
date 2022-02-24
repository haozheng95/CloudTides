package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var sshPool sync.Map

type Operate struct {
	Op      string
	Token   string
	Data    string
	CMD     string
	SshType string

	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}

type Host struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}

var url string
var mqHost string
var mqTopic string

func init() {
	mqHost = os.Getenv("MQ_HOST")
	mqTopic = os.Getenv("MQ_TOPIC")
	url = os.Getenv("SERVER_URL")
}

func main() {
	hosts := make([]*Host, 0, 10)
	yamlFile, err := ioutil.ReadFile("hosts.yaml")
	log.Print(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &hosts)
	failOnError(err, "Failed to Unmarshal")
	log.Print(hosts[0])

	log.Print("MQ host ", mqHost)
	conn, err := amqp.Dial(mqHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		mqTopic, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			op := &Operate{}
			err = json.Unmarshal(d.Body, op)
			failOnError(err, "Failed convert JSON string to struct")

			host := hosts[0]
			combo, err := execCmd(host.Host, host.User, host.Pass, op.SshType, host.Port, op.CMD)
			failOnError(err, "Failed to connect host by ssh")
			log.Printf("Received a combo: %s", combo)

			//SSHHost     string
			//SSHUser     string
			//SSHPassword string
			//SSHPort     string
			values := map[string]string{"combo": string(combo),
				"token":        op.Token,
				"ssh_host":     host.Host,
				"ssh_user":     host.User,
				"ssh_password": host.Pass,
				"ssh_port":     host.Port,
			}

			if err != nil {
				values["error"] = err.Error()
			} else {
				values["error"] = ""
			}

			log.Printf("Received a values: %v", values)

			jsonData, err := json.Marshal(values)
			failOnError(err, "Failed to convent json")
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			failOnError(err, "Failed to connect host")
			log.Printf("Received a resp: %s", resp.Status)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func execCmd(sshHost, SshUser, SshPassword, sshType, sshPort, cmd string) ([]byte, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshType == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(SshPassword)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(SshPassword)}
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

func getSession(sshHost string, sshPort string, config *ssh.ClientConfig) (session *ssh.Session) {

	addr := fmt.Sprintf("%s:%s", sshHost, sshPort)
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
