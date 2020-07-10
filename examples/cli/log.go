package main

import (
	"fmt"
	"log"
	"os"

	"github.com/c-bata/go-prompt"
)

func LogMode() func() {
	fileName := prompt.Input("Log Filename>> ", noneCompleter)
	closer := setLogFile(fileName)
	fmt.Printf("Log file config success. File: %s\n", fileName)

	return closer
}

func setLogFile(filename string) func() {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	return func() {
		f.Close()
	}
}
