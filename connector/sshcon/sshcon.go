package sshcon

import (
	"bytes"
	"fmt"
	"log"
	"time"

	// we need this "hacked" ssh
	// in order to connect to Nimbles
	"github.com/bored-engineer/ssh"
	"github.com/xjj1/StorageReporter/devices"
)

type MySSH struct {
	*ssh.Client
	Name string
}

func NewSSH(a *devices.Device) (MySSH, error) {
	log.Printf("Connecting to %s", a.Name)
	arr_ip := fmt.Sprintf("%s:22", a.Name)
	cfg := ssh.Config{}
	cfg.SetDefaults()
	cfg.KeyExchanges = append(cfg.KeyExchanges,
		"diffie-hellman-group-exchange-sha256",
		"diffie-hellman-group-exchange-sha1",
	)
	config := &ssh.ClientConfig{
		User: a.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(a.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Minute,
		Config:          cfg,
	}
	client, err := ssh.Dial("tcp", arr_ip, config)
	if err != nil {
		log.Printf("Error connecting to %s : %v", a.Name, err)
		return MySSH{}, err
	}
	return MySSH{client, a.Name}, nil
}

// needs question/answer SSH connection to the Nimbles
func NewSSHNimble(a *devices.Device) (MySSH, error) {
	log.Printf("Connecting to %s", a.Name)
	arr_ip := fmt.Sprintf("%s:22", a.Name)
	cfg := ssh.Config{}
	cfg.SetDefaults()
	cfg.KeyExchanges = append(cfg.KeyExchanges,
		"diffie-hellman-group-exchange-sha256",
		"diffie-hellman-group-exchange-sha1",
	)
	config := &ssh.ClientConfig{
		User: a.Username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				// Just send the password back for all questions
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = a.Password // replace this
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Minute,
		Config:          cfg,
	}
	client, err := ssh.Dial("tcp", arr_ip, config)
	if err != nil {
		log.Printf("Error connecting to %s : %v", a.Name, err)
		return MySSH{}, err
	}
	return MySSH{client, a.Name}, nil
}

func (c MySSH) ExecCmd(cmd string) (string, error) {
	var session *ssh.Session
	var b bytes.Buffer
	session, err := c.NewSession()
	if err != nil {
		log.Println("Failed to create session: " + err.Error())
		return "", err
	}
	defer session.Close()
	session.Stdout = &b
	log.Printf("Runnning %s on %s", cmd, c.Name)
	if err = session.Run(cmd); err != nil {
		return "", err
	} else {
		return b.String(), nil
	}
}
