package cli

import (
	"fmt"
	"log"
	"os"
)

func LogMode(logFunc func(string) string) func() {
	fileName := logFunc("Log Filename>> ")
	closer := SetLogFile(fileName)
	fmt.Printf("Log file config success. File: %s\n", fileName)

	return closer
}

func SetLogFile(filename string) func() {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	return func() {
		f.Close()
	}
}
