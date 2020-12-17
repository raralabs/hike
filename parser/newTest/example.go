package main

import (
	"context"
	"fmt"
	"github.com/raralabs/hike/parser/newMulti"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
	"time"
)


func IsComment(statement string)bool{
	if statement[0]=='#'{
		return true
	}
	return false
}



func main(){
	cmd:=`fake(100) | select(age) | stdout();;`
	//remove comments from command
	//parsed,_ := newPeg.Parse("",[]byte(sql_query))
	//fmt.Println("query",parsed)
	qryPreprocessor := LanguageProcessor.InitPreProcessor(cmd)
	cleanedCommand := qryPreprocessor.PrepareQuery(IsComment)

	fmt.Println("cleaned",cleanedCommand)
	stg,err := newPeg.Parse("", []byte(cleanedCommand))
	if err!= nil{
		panic(err)
	}
	fmt.Println("Stage",stg)
	//
	builder := newMulti.NewATBuilder()
	absTree := builder.Build(stg.([]interface{}))
	fmt.Println(absTree)
	nb := newMulti.NewNetBuilder()
	//
	p := nb.Build(1, absTree)
	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	//
	p.Start(c, cancel)
}

