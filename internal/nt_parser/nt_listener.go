// Code generated from NT.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nt_parser // NT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// NTListener is a complete listener for a parse tree produced by NTParser.
type NTListener interface {
	antlr.ParseTreeListener

	// EnterNtriplesDoc is called when entering the ntriplesDoc production.
	EnterNtriplesDoc(c *NtriplesDocContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterTriple is called when entering the triple production.
	EnterTriple(c *TripleContext)

	// EnterSubject is called when entering the subject production.
	EnterSubject(c *SubjectContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitNtriplesDoc is called when exiting the ntriplesDoc production.
	ExitNtriplesDoc(c *NtriplesDocContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitTriple is called when exiting the triple production.
	ExitTriple(c *TripleContext)

	// ExitSubject is called when exiting the subject production.
	ExitSubject(c *SubjectContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
