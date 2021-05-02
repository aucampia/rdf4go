// Code generated from NT.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nt_parser // NT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNTListener is a complete listener for a parse tree produced by NTParser.
type BaseNTListener struct{}

var _ NTListener = &BaseNTListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNTListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNTListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNTListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNTListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNtriplesDoc is called when production ntriplesDoc is entered.
func (s *BaseNTListener) EnterNtriplesDoc(ctx *NtriplesDocContext) {}

// ExitNtriplesDoc is called when production ntriplesDoc is exited.
func (s *BaseNTListener) ExitNtriplesDoc(ctx *NtriplesDocContext) {}

// EnterLine is called when production line is entered.
func (s *BaseNTListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseNTListener) ExitLine(ctx *LineContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseNTListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseNTListener) ExitComment(ctx *CommentContext) {}

// EnterTriple is called when production triple is entered.
func (s *BaseNTListener) EnterTriple(ctx *TripleContext) {}

// ExitTriple is called when production triple is exited.
func (s *BaseNTListener) ExitTriple(ctx *TripleContext) {}

// EnterSubject is called when production subject is entered.
func (s *BaseNTListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BaseNTListener) ExitSubject(ctx *SubjectContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseNTListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseNTListener) ExitPredicate(ctx *PredicateContext) {}

// EnterObject is called when production object is entered.
func (s *BaseNTListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseNTListener) ExitObject(ctx *ObjectContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseNTListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseNTListener) ExitLiteral(ctx *LiteralContext) {}
