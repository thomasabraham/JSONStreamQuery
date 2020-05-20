// Code generated from query.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // query

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 15, 4, 
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 2, 2, 4, 2, 4, 2, 2, 2, 12, 2, 6, 3, 2, 2, 2, 4, 9, 3, 2, 2, 2, 6, 7, 
	5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 7, 3, 2, 2, 10, 
	11, 7, 4, 2, 2, 11, 12, 7, 5, 2, 2, 12, 13, 7, 6, 2, 2, 13, 5, 3, 2, 2, 
	2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'SELECT'", "'*'", "'FROM'", "'STREAM'",
}
var symbolicNames = []string{
	"", "SELECT", "STAR", "FROM", "STREAM", "WS",
}

var ruleNames = []string{
	"start", "statement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type queryParser struct {
	*antlr.BaseParser
}

func NewqueryParser(input antlr.TokenStream) *queryParser {
	this := new(queryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "query.g4"

	return this
}

// queryParser tokens.
const (
	queryParserEOF = antlr.TokenEOF
	queryParserSELECT = 1
	queryParserSTAR = 2
	queryParserFROM = 3
	queryParserSTREAM = 4
	queryParserWS = 5
)

// queryParser rules.
const (
	queryParserRULE_start = 0
	queryParserRULE_statement = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = queryParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = queryParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(queryParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(queryListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(queryListener); ok {
		listenerT.ExitStart(s)
	}
}




func (p *queryParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, queryParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Statement()
	}
	{
		p.SetState(5)
		p.Match(queryParserEOF)
	}



	return localctx
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = queryParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = queryParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(queryParserSELECT, 0)
}

func (s *StatementContext) STAR() antlr.TerminalNode {
	return s.GetToken(queryParserSTAR, 0)
}

func (s *StatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(queryParserFROM, 0)
}

func (s *StatementContext) STREAM() antlr.TerminalNode {
	return s.GetToken(queryParserSTREAM, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(queryListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(queryListener); ok {
		listenerT.ExitStatement(s)
	}
}




func (p *queryParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, queryParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Match(queryParserSELECT)
	}
	{
		p.SetState(8)
		p.Match(queryParserSTAR)
	}
	{
		p.SetState(9)
		p.Match(queryParserFROM)
	}
	{
		p.SetState(10)
		p.Match(queryParserSTREAM)
	}



	return localctx
}


