// Code generated from query.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 41, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 36, 10, 6, 13, 
	6, 14, 6, 37, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 
	3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 41, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 
	2, 2, 5, 20, 3, 2, 2, 2, 7, 22, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 35, 
	3, 2, 2, 2, 13, 14, 7, 85, 2, 2, 14, 15, 7, 71, 2, 2, 15, 16, 7, 78, 2, 
	2, 16, 17, 7, 71, 2, 2, 17, 18, 7, 69, 2, 2, 18, 19, 7, 86, 2, 2, 19, 4, 
	3, 2, 2, 2, 20, 21, 7, 44, 2, 2, 21, 6, 3, 2, 2, 2, 22, 23, 7, 72, 2, 2, 
	23, 24, 7, 84, 2, 2, 24, 25, 7, 81, 2, 2, 25, 26, 7, 79, 2, 2, 26, 8, 3, 
	2, 2, 2, 27, 28, 7, 85, 2, 2, 28, 29, 7, 86, 2, 2, 29, 30, 7, 84, 2, 2, 
	30, 31, 7, 71, 2, 2, 31, 32, 7, 67, 2, 2, 32, 33, 7, 79, 2, 2, 33, 10, 
	3, 2, 2, 2, 34, 36, 9, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 
	37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 8, 
	6, 2, 2, 40, 12, 3, 2, 2, 2, 4, 2, 37, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'SELECT'", "'*'", "'FROM'", "'STREAM'",
}

var lexerSymbolicNames = []string{
	"", "SELECT", "STAR", "FROM", "STREAM", "WS",
}

var lexerRuleNames = []string{
	"SELECT", "STAR", "FROM", "STREAM", "WS",
}

type queryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewqueryLexer(input antlr.CharStream) *queryLexer {

	l := new(queryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// queryLexer tokens.
const (
	queryLexerSELECT = 1
	queryLexerSTAR = 2
	queryLexerFROM = 3
	queryLexerSTREAM = 4
	queryLexerWS = 5
)

