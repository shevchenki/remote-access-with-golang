package RunCommandOverSsh

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func runCommand(host string, port int, command string, config *ssh.ClientConfig) (string, error) {

	// Connect to remote server via SSH
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Create a new SSH session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// Execute command remotely
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func RunCommandOverSshByKey(host string, port int, user, privateKeyPath, command string) (string, error) {
	// Read the private key file
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	// Create a signer for the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return "", err
	}

	// Set up SSH client configuration
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return runCommand(host, port, command, config)
}

func RunCommandOverSshByPass(host string, port int, user, password, command string) (string, error) {
	// Set up SSH client configuration
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return runCommand(host, port, command, config)
}
