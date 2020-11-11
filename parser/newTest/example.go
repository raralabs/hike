package main

import (
	"context"
	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
	"github.com/raralabs/hike/parser/newMulti"
	"time"
)


func IsComment(statement string)bool{
	if statement[0]=='#'{
		return true
	}
	return false
}

func main(){
	cmd:=`# transform
          fake(10) | filter(age>30) | filter(age>30) | stdout();;`

	//remove comments from command
	qryPreprocessor := LanguageProcessor.InitPreProcessor(cmd)
	cleanedCommand := qryPreprocessor.PrepareQuery(IsComment)

	fmt.Println("cleaned",cleanedCommand)
	stg,_ := newPeg.Parse("", []byte(cleanedCommand))
	fmt.Println("Stage",stg)

	builder := newMulti.NewATBuilder()
	absTree := builder.BuildAT(stg.([]interface{}))
	fmt.Println(absTree)
	nb := newMulti.NewNetBuilder()

	p := nb.Build(1, absTree)
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	p.Start(c, cancel)
}

