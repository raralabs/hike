package main

import (
	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
	"reflect"
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

	cleanedCommand := LanguageProcessor.PrepareQuery(filteredCommand)
	stg,_ := newPeg.Parse("", []byte(cleanedCommand))
	for _,stage := range (stg.([]interface{})){
		for _,erin := range(stage.([]interface{})){
			fmt.Println("subodh loves erin",reflect.TypeOf(erin))
		}
	}

}
