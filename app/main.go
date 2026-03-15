package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func is_executable(argument string) (string, bool) {
	full_path, err := exec.LookPath(argument)
	if err != nil {
		return full_path, false
	} else {
		return full_path, true
	}
}

func check_type(type_argument string) string {
	if type_argument == "echo" || type_argument == "type" || type_argument == "exit" {
		return fmt.Sprintf("%s is a shell builtin\n", type_argument)
	} else {
		full_path, is_executable := is_executable(type_argument)
		if !is_executable {
			return fmt.Sprintf("%s: not found\n", type_argument)
		} else {
			return fmt.Sprintf("%s is %s\n", type_argument, full_path)
		}
	}
}

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
			if len(parts) == 2 {
				fmt.Print(check_type(parts[1]))
			}
		default:
			executable, is_exec := is_executable(parts[0])
			if is_exec {
				cmd := exec.Command(executable, parts[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				//if err != nil {
				//	log.Fatal(err)
				//}
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}
