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
		var builder strings.Builder

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		parts := strings.Split(command, " ")

		switch parts[0] {
		case "echo":
			for i := 1; i < len(parts); i++ {
				builder.WriteString(parts[i] + " ")
			}
			fmt.Println(strings.TrimSpace(builder.String()))
		case "type":
			if len(parts) == 2 && (parts[1] == "echo" || parts[1] == "type" || parts[1] == "exit") {
				fmt.Printf("%s is a shell builtin\n", parts[1])
			} else {
				fmt.Printf("%s: not found\n", parts[1])
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
