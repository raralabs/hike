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
	cmd:=`# transform
file(anything.csv) | filter(age>30) | filter(age>30) | b1 = into();
          branch(b1) | agg counter = count() | stdout();
          file(anything.csv) | filter(age>30) | b1 = into();
          branch(b1) | filter(age>30) | stdout();
branch(b1) | agg by last_name last_name_count = count() | stdout();;`

	//remove comments from command
	filteredCommand := newPeg.RemoveStatement(cmd,IsComment)
	fmt.Println("comment filtered",filteredCommand)
	cleanedCommand := LanguageProcessor.PrepareQuery(filteredCommand)
	stg,err := newPeg.Parse("", []byte(cleanedCommand))

	fmt.Println(stg,err)
}
