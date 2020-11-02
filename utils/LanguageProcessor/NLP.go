package LanguageProcessor

import "strings"

func Tokenize(sentences string,separator string )[]string{
	var cleanedTokens []string
	tokens := strings.Split(sentences, separator)
	for _,token :=  range tokens{
		cleanedTokens = append(cleanedTokens,strings.TrimSpace(token))
	}
	return cleanedTokens
}

