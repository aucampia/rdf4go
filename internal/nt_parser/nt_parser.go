// Code generated from NT.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nt_parser // NT
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 88, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2, 25,
	11, 2, 3, 2, 5, 2, 28, 10, 2, 3, 3, 5, 3, 31, 10, 3, 3, 3, 5, 3, 34, 10,
	3, 3, 4, 7, 4, 37, 10, 4, 12, 4, 14, 4, 40, 11, 4, 3, 4, 3, 4, 3, 5, 7,
	5, 45, 10, 5, 12, 5, 14, 5, 48, 11, 5, 3, 5, 3, 5, 7, 5, 52, 10, 5, 12,
	5, 14, 5, 55, 11, 5, 3, 5, 3, 5, 7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11,
	5, 3, 5, 3, 5, 7, 5, 66, 10, 5, 12, 5, 14, 5, 69, 11, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 80, 10, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 86, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2,
	3, 4, 2, 8, 8, 10, 10, 2, 92, 2, 18, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6,
	38, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12, 74, 3, 2, 2,
	2, 14, 79, 3, 2, 2, 2, 16, 81, 3, 2, 2, 2, 18, 23, 5, 4, 3, 2, 19, 20,
	7, 7, 2, 2, 20, 22, 5, 4, 3, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2,
	23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3,
	2, 2, 2, 26, 28, 7, 7, 2, 2, 27, 26, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28,
	3, 3, 2, 2, 2, 29, 31, 5, 8, 5, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2,
	2, 31, 33, 3, 2, 2, 2, 32, 34, 5, 6, 4, 2, 33, 32, 3, 2, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 37, 7, 17, 2, 2, 36, 35, 3, 2, 2, 2,
	37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3,
	2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 42, 7, 5, 2, 2, 42, 7, 3, 2, 2, 2, 43,
	45, 7, 17, 2, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2,
	2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 53,
	5, 10, 6, 2, 50, 52, 7, 17, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2,
	2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53,
	3, 2, 2, 2, 56, 60, 5, 12, 7, 2, 57, 59, 7, 17, 2, 2, 58, 57, 3, 2, 2,
	2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63,
	3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 67, 5, 14, 8, 2, 64, 66, 7, 17, 2,
	2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68,
	3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 71, 7, 3, 2, 2,
	71, 9, 3, 2, 2, 2, 72, 73, 9, 2, 2, 2, 73, 11, 3, 2, 2, 2, 74, 75, 7, 8,
	2, 2, 75, 13, 3, 2, 2, 2, 76, 80, 7, 8, 2, 2, 77, 80, 7, 10, 2, 2, 78,
	80, 5, 16, 9, 2, 79, 76, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2,
	2, 2, 80, 15, 3, 2, 2, 2, 81, 85, 7, 9, 2, 2, 82, 83, 7, 4, 2, 2, 83, 86,
	7, 8, 2, 2, 84, 86, 7, 6, 2, 2, 85, 82, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2,
	85, 86, 3, 2, 2, 2, 86, 17, 3, 2, 2, 2, 13, 23, 27, 30, 33, 38, 46, 53,
	60, 67, 79, 85,
}
var literalNames = []string{
	"", "'.'", "'^^'",
}
var symbolicNames = []string{
	"", "", "", "COMMENT", "LANGTAG", "EOL", "IRIREF", "STRING_LITERAL_QUOTE",
	"BLANK_NODE_LABEL", "UCHAR", "ECHAR", "PN_CHARS_BASE", "PN_CHARS_U", "PN_CHARS",
	"HEX", "WS",
}

var ruleNames = []string{
	"ntriplesDoc", "line", "comment", "triple", "subject", "predicate", "object",
	"literal",
}

type NTParser struct {
	*antlr.BaseParser
}

// NewNTParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *NTParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNTParser(input antlr.TokenStream) *NTParser {
	this := new(NTParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NT.g4"

	return this
}

// NTParser tokens.
const (
	NTParserEOF                  = antlr.TokenEOF
	NTParserT__0                 = 1
	NTParserT__1                 = 2
	NTParserCOMMENT              = 3
	NTParserLANGTAG              = 4
	NTParserEOL                  = 5
	NTParserIRIREF               = 6
	NTParserSTRING_LITERAL_QUOTE = 7
	NTParserBLANK_NODE_LABEL     = 8
	NTParserUCHAR                = 9
	NTParserECHAR                = 10
	NTParserPN_CHARS_BASE        = 11
	NTParserPN_CHARS_U           = 12
	NTParserPN_CHARS             = 13
	NTParserHEX                  = 14
	NTParserWS                   = 15
)

// NTParser rules.
const (
	NTParserRULE_ntriplesDoc = 0
	NTParserRULE_line        = 1
	NTParserRULE_comment     = 2
	NTParserRULE_triple      = 3
	NTParserRULE_subject     = 4
	NTParserRULE_predicate   = 5
	NTParserRULE_object      = 6
	NTParserRULE_literal     = 7
)

// INtriplesDocContext is an interface to support dynamic dispatch.
type INtriplesDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNtriplesDocContext differentiates from other interfaces.
	IsNtriplesDocContext()
}

type NtriplesDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNtriplesDocContext() *NtriplesDocContext {
	var p = new(NtriplesDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_ntriplesDoc
	return p
}

func (*NtriplesDocContext) IsNtriplesDocContext() {}

func NewNtriplesDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NtriplesDocContext {
	var p = new(NtriplesDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_ntriplesDoc

	return p
}

func (s *NtriplesDocContext) GetParser() antlr.Parser { return s.parser }

func (s *NtriplesDocContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *NtriplesDocContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *NtriplesDocContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(NTParserEOL)
}

func (s *NtriplesDocContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(NTParserEOL, i)
}

func (s *NtriplesDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NtriplesDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NtriplesDocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterNtriplesDoc(s)
	}
}

func (s *NtriplesDocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitNtriplesDoc(s)
	}
}

func (p *NTParser) NtriplesDoc() (localctx INtriplesDocContext) {
	localctx = NewNtriplesDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NTParserRULE_ntriplesDoc)
	var _la int

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Line()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(17)
				p.Match(NTParserEOL)
			}
			{
				p.SetState(18)
				p.Line()
			}

		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NTParserEOL {
		{
			p.SetState(24)
			p.Match(NTParserEOL)
		}

	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Triple() ITripleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *NTParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NTParserRULE_line)
	var _la int

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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(27)
			p.Triple()
		}

	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NTParserCOMMENT || _la == NTParserWS {
		{
			p.SetState(30)
			p.Comment()
		}

	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(NTParserCOMMENT, 0)
}

func (s *CommentContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(NTParserWS)
}

func (s *CommentContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(NTParserWS, i)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *NTParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NTParserRULE_comment)
	var _la int

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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NTParserWS {
		{
			p.SetState(33)
			p.Match(NTParserWS)
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(NTParserCOMMENT)
	}

	return localctx
}

// ITripleContext is an interface to support dynamic dispatch.
type ITripleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTripleContext differentiates from other interfaces.
	IsTripleContext()
}

type TripleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripleContext() *TripleContext {
	var p = new(TripleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_triple
	return p
}

func (*TripleContext) IsTripleContext() {}

func NewTripleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleContext {
	var p = new(TripleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_triple

	return p
}

func (s *TripleContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleContext) Subject() ISubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubjectContext)
}

func (s *TripleContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *TripleContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *TripleContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(NTParserWS)
}

func (s *TripleContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(NTParserWS, i)
}

func (s *TripleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TripleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterTriple(s)
	}
}

func (s *TripleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitTriple(s)
	}
}

func (p *NTParser) Triple() (localctx ITripleContext) {
	localctx = NewTripleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NTParserRULE_triple)
	var _la int

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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NTParserWS {
		{
			p.SetState(41)
			p.Match(NTParserWS)
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Subject()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NTParserWS {
		{
			p.SetState(48)
			p.Match(NTParserWS)
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Predicate()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NTParserWS {
		{
			p.SetState(55)
			p.Match(NTParserWS)
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Object()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NTParserWS {
		{
			p.SetState(62)
			p.Match(NTParserWS)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(NTParserT__0)
	}

	return localctx
}

// ISubjectContext is an interface to support dynamic dispatch.
type ISubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubjectContext differentiates from other interfaces.
	IsSubjectContext()
}

type SubjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubjectContext() *SubjectContext {
	var p = new(SubjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_subject
	return p
}

func (*SubjectContext) IsSubjectContext() {}

func NewSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubjectContext {
	var p = new(SubjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_subject

	return p
}

func (s *SubjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SubjectContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(NTParserIRIREF, 0)
}

func (s *SubjectContext) BLANK_NODE_LABEL() antlr.TerminalNode {
	return s.GetToken(NTParserBLANK_NODE_LABEL, 0)
}

func (s *SubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterSubject(s)
	}
}

func (s *SubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitSubject(s)
	}
}

func (p *NTParser) Subject() (localctx ISubjectContext) {
	localctx = NewSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NTParserRULE_subject)
	var _la int

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
		p.SetState(70)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NTParserIRIREF || _la == NTParserBLANK_NODE_LABEL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(NTParserIRIREF, 0)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *NTParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NTParserRULE_predicate)

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
		p.SetState(72)
		p.Match(NTParserIRIREF)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(NTParserIRIREF, 0)
}

func (s *ObjectContext) BLANK_NODE_LABEL() antlr.TerminalNode {
	return s.GetToken(NTParserBLANK_NODE_LABEL, 0)
}

func (s *ObjectContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *NTParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NTParserRULE_object)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NTParserIRIREF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(NTParserIRIREF)
		}

	case NTParserBLANK_NODE_LABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(NTParserBLANK_NODE_LABEL)
		}

	case NTParserSTRING_LITERAL_QUOTE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NTParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NTParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_LITERAL_QUOTE() antlr.TerminalNode {
	return s.GetToken(NTParserSTRING_LITERAL_QUOTE, 0)
}

func (s *LiteralContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(NTParserIRIREF, 0)
}

func (s *LiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(NTParserLANGTAG, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NTListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *NTParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NTParserRULE_literal)

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
		p.SetState(79)
		p.Match(NTParserSTRING_LITERAL_QUOTE)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NTParserT__1:
		{
			p.SetState(80)
			p.Match(NTParserT__1)
		}
		{
			p.SetState(81)
			p.Match(NTParserIRIREF)
		}

	case NTParserLANGTAG:
		{
			p.SetState(82)
			p.Match(NTParserLANGTAG)
		}

	case NTParserT__0, NTParserWS:

	default:
	}

	return localctx
}
