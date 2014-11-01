package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {
	var pw string
	print(" tell me the password: ")

	// cmd := exec.Command("bash", "-c", "stty", "-echo")
	// cmd.Start()

	fmt.Scanf("%s", &pw)
	// cmd = exec.Command("bash", "-c", "stty", "echo")
	// cmd.Start()
	println("you typed ", pw)
}

func Getpasswd(prompt string) (passwd string, err error) {
	fmt.Print(prompt)
	const stty_arg0 = "/bin/stty"
	stty_argv_e_off := []string{"stty", "-echo"}
	stty_argv_e_on := []string{"stty", "echo"}
	const exec_cwdir = ""
	fd := []*os.File{os.Stdin, os.Stdout, os.Stderr}
	pid, err := syscall.ForkExec
	if err != nil {
		return passwd, error.NewError(fmt.Sprintf("Failed turning off console echo for password entry:\n\t%s", err))
	}
	rd := bufio.NewReader(os.Stdin)
	os.Wait(pid, 0)
	line, err := rd.ReadString('\n')
	if err == nil {
		passwd = str.TrimSpace(line)
	} else {
		err = os.NewError(fmt.Sprintf("Failed during password entry:%s", err))
	}
	pid, e := os.ForkExec(stty_arg0, stty_argv_e_on, nil, exec_cwdir, fd)
	if e == nil {
		os.Wait(pid, 0)
	} else if err == nil {
		err = os.NewError(fmt.Sprintf("Failed turning on console echo post password entry:\n\t%s", e))
	}
	return passwd, err
}
