package main

import (
	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
)


func IsComment(statement string)bool{
	if statement[0]=='#'{
		return true
	}
	return false
}

func main(){
	cmd:=`file(anything.csv) | filter(age>30) | filter(age>30) | b1 = into();
          branch(b1) | agg counter = count() | stdout();
          file(anything.csv) | filter(age>30) | b1 = into();
          branch(b1) | filter(age>30) | stdout();;`

	//remove comments from command
	//filteredCommand := newPeg.RemoveStatement(cmd,IsComment)
	cleanedCommand := LanguageProcessor.PrepareQuery(cmd)
	stg,err := newPeg.Parse("", []byte(cleanedCommand))

	fmt.Println(stg,err)
}
