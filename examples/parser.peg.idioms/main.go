package main

import (
	"fmt"
	"log"

	"github.com/raralabs/hike/parser/peg"
)

func main() {
	infos, err := peg.Parse("", []byte(`file(users.csv) | filter(age == "||" and (age%2 == 1 or age < 10)) | stdout()`))

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%+v\n", infos)
}
