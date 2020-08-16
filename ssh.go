package main

import (
	"bytes"
	// "fmt"
	"io/ioutil"
	"log"
	// "net"
	"os"
	"os/user"
	"strconv"

	// "strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	// "golang.org/x/crypto/ssh/knownhosts"
)

var (
	sshKey = getLocalSSHKey()
	usr = getUser()
)

type (
	host struct {
		Host string
		Port int
	}
)

func runScript(script string, args []string, host host, sudo bool) (string, string, error) {
	hkcb, err := knownhosts.New(getHome()+"/.ssh/known_hosts")
	if err != nil {
        return "", "", err
    }
	signer, err := ssh.ParsePrivateKey(sshKey)
    if err != nil {
        return "", "", err
    }
	cfg := &ssh.ClientConfig{
		User: getUser(),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSHPASS")),
			ssh.PublicKeys(signer),
			// ssh.PublicKeys(signers ...ssh.Signer)
		},
		HostKeyCallback: hkcb,
	}
	conn, err := ssh.Dial("tcp", host.String(), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}
	session.Stdin = bytes.NewReader([]byte(script))
	session.Stderr = stderr
	session.Stdout = stdout
	cmd := "env bash -s --"
	if sudo {
		cmd = "sudo " + cmd
	}
	if len(args) > 0 {
		for _, arg := range args {
			cmd = cmd + " " + arg
		}
	}
	err = session.Run(cmd)
	return stdout.String(), stderr.String(), err
}

func getLocalSSHKey() []byte {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadFile(u.HomeDir + "/.ssh/id_rsa")
    if err != nil {
        log.Fatalln(err)
    }
	return b
}

func getUser() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return u.Username
}

func getHome() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return u.HomeDir
}

func (h host) String() string {
	return h.Host + ":" + strconv.Itoa(h.Port)
}