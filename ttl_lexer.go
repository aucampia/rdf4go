package rdf4go

import (
	"bufio"
	"io"
	"unicode/utf8"
)

type itemType int

// item represents a token or text string returned from the scanner.
type item struct {
	typ  itemType
	pos  int
	line int
	col  int
	val  string
}

type stateFn func(*lexer) stateFn

const (
	itemError itemType = iota
	itemEOF
	itemEOL
	itemIRI
	itemBNode
	itemLiteral
)

type lexer struct {
	name  string
	rdr   *bufio.Reader
	items chan item
	input []byte
	pos   int
	start int
	line  int
	width int
}

const eof = -1

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

func (l *lexer) lex() {
	go l.run()
	return l
}

func (l *lexer) feed() (bool, error) {
	bytes := byte[1024]
	count, err := l.rdr.Read(bytes)
	if err == io.EOF {
		return false
	}
	if err != nil {
		return nil, err
	}
	l.input = append(l.input, bytes)
	return true
}

func (l *lexer) run() {
	for {
		if !l.feed(false) {
			l.emit(itemEOF)
			break
		}

		for l.state = lexAny; l.state != nil; {
			l.state = l.state(l)
		}
	}

	// No more input to lex, emit final EOF token and terminate.
	// The value of the closed tokens channel is tokenEOF.
	close(l.tokens)
}

func newLexer() (*lexer, error) {
	return &lexer{
		name: name,
		rdr:  rdr,
	}, nil
}
