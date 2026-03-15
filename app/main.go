package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else if strings.HasPrefix(command, "type ") {
			if command[5:5+len("echo")] == "echo" {
				fmt.Println("echo is a shell builtin")
			} else if command[5:5+len("exit")] == "exit" {
				fmt.Println("exit is a shell builtin")
			} else if command[5:5+len("type")] == "type" {
				fmt.Println("type is a shell builtin")
			} else {
				fmt.Println(command[5:] + " is a shell builtin")
			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}

}
