// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

package rdf4go

import (
	"fmt"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
	ntp "gobn.github.io/rdf4go/internal/nt_parser"
	"gobn.github.io/rdf4go/internal/utils"
)

var (
	svf = SimpleValueFactory{}
)

type TreeShapeListener struct {
	*ntp.BaseNTListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Fprintf(os.Stderr, "%s\n", ctx.GetText())
}

func (t *TreeShapeListener) EnterSubject(ctx *ntp.SubjectContext) {
	fmt.Fprintf(os.Stderr, "subject = %s\n", ctx.GetText())
	if iriT := ctx.GetToken(ntp.NTLexerIRIREF, 0); iriT != nil {
		t := iriT.GetText()
		iri, err := svf.NewIRI(t[1 : len(t)-1])
		if err != nil {
			utils.Flog().Errorf("err = %v", err)
		}
		fmt.Fprintf(os.Stderr, "subject.iri = %v -> %v\n", iriT, iri)
	}
}

func (t *TreeShapeListener) EnterPredicate(ctx *ntp.SubjectContext) {
	fmt.Fprintf(os.Stderr, "subject = %s\n", ctx.GetText())
	if iriT := ctx.GetToken(ntp.NTLexerIRIREF, 0); iriT != nil {
		t := iriT.GetText()
		iri, err := svf.NewIRI(t[1 : len(t)-1])
		if err != nil {
			utils.Flog().Errorf("err = %v", err)
		}
		fmt.Fprintf(os.Stderr, "subject.iri = %v -> %v\n", iriT, iri)
	}
	if bnodeT := ctx.GetToken(ntp.NTLexerBLANK_NODE_LABEL, 0); bnodeT != nil {
		bnode, err := svf.NewBNode(bnodeT.GetText()[2:])
		if err != nil {
			utils.Flog().Errorf("err = %v", err)
		}
		fmt.Fprintf(os.Stderr, "subject.bnode = %v -> %v\n", bnodeT, bnode)
	}
}

func TestNTParser(t *testing.T) {
	// https://github.com/antlr/antlr4/blob/master/doc/go-target.md
	input := antlr.NewInputStream(
		`<http://one.example/subject1> <http://one.example/predicate1> <http://one.example/object1> . # comments here
# or on a line by themselves
_:subject1 <http://an.example/predicate1> "object1" .
_:subject2 <http://an.example/predicate2> "object2" .`)
	lexer := ntp.NewNTLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := ntp.NewNTParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.NtriplesDoc()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	assert.Equal(t, 1, 1)
}
