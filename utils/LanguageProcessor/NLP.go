package LanguageProcessor

import "strings"

//tokenizes the sentence or statement based on the
//separator provided and returns the slice of stripped
//tokens
func Tokenize(sentences string,separator string )[]string{
	var cleanedTokens []string
	tokens := strings.Split(sentences, separator)
	for _,token :=  range tokens{
		cleanedTokens = append(cleanedTokens,strings.TrimSpace(token))
	}
	return cleanedTokens
}
//It takes the query after all comments have been removed
//and then strips all the extra preceding and trailing white
//spaces
func PrepareQuery(command string)string{
	var cleanedStatements []string
	cmdTokens := Tokenize(command,";")
	for _,statement := range cmdTokens{
		cleanedStatement := strings.TrimSpace(statement)
		cleanedStatements = append(cleanedStatements,cleanedStatement)
	}
	joinedCommand :=  strings.Join(cleanedStatements,";")
	return joinedCommand
}
