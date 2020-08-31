package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/raralabs/hike/parser/single"
	"log"
	"os"
	"time"

	"github.com/raralabs/hike/utils/faker"
)

const TmpPath = "./"
const TmpFile = TmpPath + "users.csv"

const (
	defaultCommand = ``
	usage          = "command for the cli tool"
)

var commandFlag = flag.String("cmd", defaultCommand, usage)

func main() {
	// File Generation

	//if err := os.MkdirAll(TmpPath, os.ModePerm); err != nil {
	//	log.Panic(err)
	//}
	//generateCsv(TmpFile, 10)

	flag.Parse()

	closeFile := setLogFile("./tmp/test.log")
	defer closeFile()

	start := time.Now()
	cmd := *commandFlag
	if cmd == defaultCommand {
		panic("Command not passed")
	}

	builder := single.NewParser()

	// Initialize the builder
	builder.Init()

	// Build the command
	starter, ppln, closer := builder.Build(1, cmd)
	defer closer()

	fmt.Printf("Time taken to build the Command: %v\n", time.Since(start))

	// Create context to run the pipeline
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	// Run the starter to initialize all stages of a pipeline
	starter()

	// Run the pipeline
	ppln.Start(c, cancel)
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

func generateCsv(filename string, rows int) {

	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	// Generate Stream CSV Data
	choices := map[string][]interface{}{
		"first_name": {"Madhav", "Shambhu", "Pushpa", "Kumar", "Hero"},
		"last_name":  {"Mashima", "Dahal", "Poudel", "Rimal", "Amatya", "Shrestha", "Bajracharya"},
		"age":        {10, 20, 30, 40, 50, 60, 70, 15, 25, 35, 45, 55, 65, 75, 100, 6, 33, 47},
	}
	faker.FakeCsvData(file, choices, rows)
	file.Close()

	log.Println("CSV File Generation Complete")
	log.Printf("Total Rows: \t%v", rows)
}
