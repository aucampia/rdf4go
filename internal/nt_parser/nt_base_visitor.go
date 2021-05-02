// Code generated from NT.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nt_parser // NT
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseNTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNTVisitor) VisitNtriplesDoc(ctx *NtriplesDocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitTriple(ctx *TripleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitSubject(ctx *SubjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNTVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
