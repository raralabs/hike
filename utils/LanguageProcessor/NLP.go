package LanguageProcessor

import (
	"strings"
)

//it preprocesses the command like removing the leading and
//trailing whitespaces and other unwanted comments
//before feeding to the peg parser to parse the stages.
type preProcessor struct{
	command 	string
}

//initialize the preprocessor with the required command to be
//preprocessed
func InitPreProcessor(command string)*preProcessor{
	return &preProcessor{command:command}
}

//tokenizes the sentence or statement based on the
//separator provided and returns the slice of stripped
//tokens
func (p *preProcessor)Tokenize(separator string )[]string{
	//holds the cleaned tokens after the removal of extra white spaces if any
	var cleanedTokens []string
	tokens := strings.Split(p.command, separator)
	for _,token :=  range tokens{
		cleanedTokens = append(cleanedTokens,strings.TrimSpace(token))
	}
	return cleanedTokens
}

//removes the whole statement if the criteria is met
func (p *preProcessor)RemoveStatement(criteria func(string)bool)string{
	//container to hold the commands after the removal of comments
	var filteredStatements []string
	tokenizedStatements := p.Tokenize("\n")
	for _,statement  := range tokenizedStatements{
		if !criteria(statement){
			filteredStatements = append(filteredStatements,statement)
		}
	}
	joinedStatements := strings.Join(filteredStatements,"\n")
	return joinedStatements
}

//It takes the query after all comments have been removed
//and then strips all the extra preceding and trailing white
//spaces
func (p *preProcessor)PrepareQuery(IsComment func(string)bool)string{
	commentsRemovedCmd:= p.RemoveStatement(IsComment)
	return commentsRemovedCmd
}


