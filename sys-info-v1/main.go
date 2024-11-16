package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("hello, enter your input: 1 to run ls, 2 to run pwd, 3 to run df")

	var myOption string
	_, err := fmt.Scan(&myOption)
	if err != nil {
		fmt.Println("error during read: ", err)
		return
	}

	fmt.Println("you entered", myOption)

	var cmd *exec.Cmd
	var stdout []byte
	var errExec error
	switch myOption {
	case "1":
		fmt.Println("you chose to run 'ls'")
		cmd = exec.Command("ls")
		stdout, errExec = cmd.Output()
	case "2":
		fmt.Println("you chose to run 'pwd'")
		cmd = exec.Command("pwd")
		stdout, errExec = cmd.Output()
	case "3":
		fmt.Println("you chose to run 'df'")
		cmd = exec.Command("df")
		stdout, errExec = cmd.Output()
	default:
		fmt.Printf("option %s not supported\n", myOption)
		errExec = errors.New("option unsupported")
	}
	if errExec != nil {
		fmt.Println("error during os exec:", errExec)
		return
	}

	fmt.Printf("output:\n-----------\n%s----------\n", string(stdout))
}
