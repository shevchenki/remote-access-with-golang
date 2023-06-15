package main

import (
	"fmt"
	"remote-access-with-golang/RunCommandOverSsh"
)

func main() {

	var host, user, password, command string
	var port int

	fmt.Print("Enter host name: ")
	fmt.Scanln(&host)

	fmt.Print("Enter port number: ")
	fmt.Scanln(&port)

	fmt.Print("Enter user name: ")
	fmt.Scanln(&user)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	fmt.Print("Enter command: ")
	fmt.Scanln(&command)

	output, _ := RunCommandOverSsh.RunCommandOverSshByPass(host, port, user, password, command)
	fmt.Println(string(output))
}
