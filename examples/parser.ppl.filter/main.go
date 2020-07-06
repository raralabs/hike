package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/raralabs/hike/parser/ppl"
)

func main() {

	filter := &ppl.Filter{}
	err := ppl.HikeParser.Parse(strings.NewReader(`age > 30 || a == 10 && age < 50`), filter)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(filter.Functions.First.Function)
	for _, flt := range filter.Functions.Second {
		fmt.Printf("%+v  %+v\n", flt.Op, flt.Filter.Function)
	}
}
