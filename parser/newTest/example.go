package main

import (

	"fmt"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/hike/utils/LanguageProcessor"
	"github.com/raralabs/hike/parser/newMulti"
)


func IsComment(statement string)bool{
	if statement[0]=='#'{
		return true
	}
	return false
}

func main(){
	cmd:=`fake(10) | filter(age>30) | agg alias = sum(age) | s1=into();
		s1,s2 | join alias=rightouter( select * on age == age) | stdout();;`
	//remove comments from command
	//parsed,_ := newPeg.Parse("",[]byte(sql_query))
	//fmt.Println("query",parsed)
	qryPreprocessor := LanguageProcessor.InitPreProcessor(cmd)
	cleanedCommand := qryPreprocessor.PrepareQuery(IsComment)

	//fmt.Println("cleaned",cleanedCommand)
	stg,_ := newPeg.Parse("", []byte(cleanedCommand))
	fmt.Println("Stage",stg)
	//
	builder := newMulti.NewATBuilder()
	absTree := builder.BuildAT(stg.([]interface{}))
	fmt.Println(absTree)
	//nb := newMulti.NewNetBuilder()
	//
	//p := nb.Build(1, absTree)
	//c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	//
	//p.Start(c, cancel)
}

