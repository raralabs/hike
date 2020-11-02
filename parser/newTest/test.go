package main

import (
	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
)

func main(){
	cmd:=`file(something.csv) | filter(age > 30) | stdout();
          file(something.csv) | filter(age > 30) | stdout();
          fake(10) | filter(age>20) | stdout();;`
	stg,err := newPeg.Parse("", []byte(cmd))
	fmt.Println(stg,err)
}
