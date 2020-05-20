// Code generated from query.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasequeryListener is a complete listener for a parse tree produced by queryParser.
type BasequeryListener struct{}

var _ queryListener = &BasequeryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasequeryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasequeryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasequeryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasequeryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasequeryListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasequeryListener) ExitStart(ctx *StartContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasequeryListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasequeryListener) ExitStatement(ctx *StatementContext) {}
