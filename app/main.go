package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var buildins = []string{
	"echo",
	"exit",
	"type",
	"pwd",
	"cd",
}

func is_executable(argument string) (string, bool) {
	full_path, err := exec.LookPath(argument)
	if err != nil {
		return full_path, false
	} else {
		return full_path, true
	}
}

func check_type(type_argument string) string {
	if slices.Contains(buildins, type_argument) {
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

func directory_exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func echoParser(input string) []string {
	var args []string
	var current strings.Builder
	inQuote := false

	for i := 0; i < len(input); i++ {
		ch := input[i]

		switch {
		case ch == '\'' && !inQuote:
			inQuote = true
		case ch == '\'' && inQuote:
			inQuote = false
		case ch == ' ' && !inQuote:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteByte(ch) // write character as-is
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

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
		}

		//parts := strings.Split(command, " ")
		//parts := strings.Fields(command)
		args := echoParser(command)
		switch args[0] {
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "type":
			if len(args) == 2 {
				fmt.Print(check_type(args[1]))
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dir)
		case "cd":
			if len(args) == 2 {
				if directory_exists(args[1]) {
					err := os.Chdir(args[1])
					if err != nil {
						log.Fatal(err)
					}
				} else if args[1] == "~" {
					home := os.Getenv("HOME")
					err := os.Chdir(home)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					fmt.Printf("cd: %s: No such file or directory\n", args[1])
				}
			}
		case "cat":
			cmd := exec.Command("cat", args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		default:
			_, is_exec := is_executable(args[0])
			if is_exec {
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}
