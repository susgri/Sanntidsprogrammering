package main

import (
	"os/exec"
)



func main() {

	cmd := exec.Command("cmd", "/C", "start", "powershell", "go", "run", "program/program.go")

	cmd.Run()

}


