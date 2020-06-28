package jsq

import (
	"github.com/thomasabraham/JSONStreamQuery/query/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type JSONStreamQuerier interface {
	Query(input []interface{}) []interface{}
}

type JSONStreamQuery struct {
}

type queryListener struct {
	*parser.BasequeryListener
}

func BuildJSONStreamQuerier(query string) JSONStreamQuerier {
	is := antlr.NewInputStream(query)
	// Create the Lexer
	lexer := parser.NewqueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewqueryParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&queryListener{}, p.Start())

	return JSONStreamQuery{}
}

func (JSONStreamQuery) Query(input []interface{}) []interface{} {
	return input
}
