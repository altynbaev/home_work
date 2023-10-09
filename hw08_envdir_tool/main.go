package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func dirExists(dir string) bool {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func main() {
	envDir := os.Args[1]
	if envDir == "" {
		log.Fatalf("Empty path")
	}
	if !dirExists(envDir) {
		log.Fatal("Directory does not exists or not enough permissions")
	}
	env, err := ReadDir(envDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	RunCmd(os.Args[2:], env)
}
