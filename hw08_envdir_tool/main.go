package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func main() {
	envDir := os.Args[1]
	if envDir == "" {
		log.Fatalf("Empty path")
	}
	if !FileExists(envDir) {
		log.Fatal("Directory does not exists or not enough permissions")
	}
	env, err := ReadDir(envDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	RunCmd(os.Args[2:], env)
}
