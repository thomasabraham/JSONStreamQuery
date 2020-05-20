// Code generated from query.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// queryListener is a complete listener for a parse tree produced by queryParser.
type queryListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)
}
