package main

import (
	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
)

func main(){
	cmd:=`file(something.csv) | filter(age>30) | filter(age>20) | p1 = into();
          file(anything.csv) | filter(age > 30) | stdout();
          branch(b1) | filter(age > 30) | stdout();;`

	cleanedCommand := LanguageProcessor.PrepareQuery(cmd)
	fmt.Println("cleaned",cleanedCommand)
	stg,err := newPeg.Parse("", []byte(cleanedCommand))

	fmt.Println(stg,err)
}
