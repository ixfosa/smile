package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("calc")
	if err != nil {
		panic(err)
	}
	fmt.Println(path) // D:\env\go\bin\go.exe

	// func Command(name string, arg ...string) *Cmd
	cmd := exec.Command(path)

	cmd.Run()
}
