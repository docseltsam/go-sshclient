package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	// https://skarlso.github.io/2019/02/17/go-ssh-with-host-key-verification/

	// Private Key auth
	// https://gist.github.com/stefanprodan/2d20d0c6fdab6f14ce8219464e8b4b9a

	// var hostKey ssh.PublicKey
	// hostKeyCallback, err := knownhosts.New("~/.ssh/known_hosts")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	config := &ssh.ClientConfig{
		User: "usr",
		Auth: []ssh.AuthMethod{
			ssh.Password("pwd"),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
		// HostKeyCallback: hostKeyCallback,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.178.39:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	var b bytes.Buffer
	session.Stdout = &b
	//command := "/usr/bin/whoami"
	command := "docker container ls"
	if err := session.Run(command); err != nil {
		log.Fatal("Failed to run: ", err.Error())
	}
	fmt.Println(b.String())
}
