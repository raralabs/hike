package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/raralabs/canal/utils/cast"
	"github.com/raralabs/hike/utils/faker"
	"log"
	"os"
)

func FakerMode() {
	fileName := prompt.Input("Filename>> ", noneCompleter)
	numRows := prompt.Input("NumRows>> ", noneCompleter)
	if rows, ok := cast.TryInt(numRows); ok {
		generateCsv(fileName, int(rows))
		fmt.Printf("Fake csv file generation success. File: %s\n", fileName)
	} else {
		fmt.Println("Fake csv file generation failed.")
	}
}

func generateCsv(filename string, rows int) {

	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	// Generate Fake CSV Data
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