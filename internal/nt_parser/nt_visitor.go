// Code generated from NT.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nt_parser // NT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by NTParser.
type NTVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NTParser#ntriplesDoc.
	VisitNtriplesDoc(ctx *NtriplesDocContext) interface{}

	// Visit a parse tree produced by NTParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by NTParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by NTParser#triple.
	VisitTriple(ctx *TripleContext) interface{}

	// Visit a parse tree produced by NTParser#subject.
	VisitSubject(ctx *SubjectContext) interface{}

	// Visit a parse tree produced by NTParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by NTParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by NTParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
