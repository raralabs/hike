package ppl

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

type Filter struct {
	Functions *SecondaryFlt `@@`
}

type SecondaryFlt struct {
	First  *PrimaryFlt     `@@`
	Second []*FilterSecond `{ @@ }`
}

type FilterSecond struct {
	Op     string      ` @SecondaryOps`
	Filter *PrimaryFlt ` @@`
}

type PrimaryFlt struct {
	Function string ` @PrimaryFlt`
}

var (
	basicGrammar = `
		Value = (Int | Float | String) .
		Ident = (alpha | "_") { "_" | alpha | digit } .
		String = "\"" { "\u0000"…"\uffff"-"\""-"\\" | "\\" any } "\"" .
		Int = [ "-" | "+" ] digit { digit } .
		Float = ("." | digit) {"." | digit} .
		Whitespace = " " | "\t" | "\n" | "\r" .
   		Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .

		alpha = "a"…"z" | "A"…"Z" .
		digit = "0"…"9" .
		any = "\u0000"…"\uffff" .
	`

	filterGrammar = `
		PrimaryFlt = Ident {Whitespace} PrimaryOps {Whitespace} Value .
		PrimaryOps = ">" | "<" | "<=" | ">=" | "==" .
		SecondaryOps = "&&" | "||" .
	`

	grammar = filterGrammar + basicGrammar

	hikeLexer = lexer.Must(ebnf.New(grammar))

	HikeParser = participle.MustBuild(&Filter{},
		participle.Lexer(hikeLexer),
		//ppl.Unquote("String"),
		participle.Elide("Whitespace"),
	)
)
