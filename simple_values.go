package rdf4go

import (
	"errors"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// type SimpleBlank
////////////////////////////////////////////////////////////////////////////////

type SimpleBlank struct {
	id string
}

func (*SimpleBlank) isSubjectType() {}
func (*SimpleBlank) isObjectType()  {}

func (s *SimpleBlank) Id() string {
	return s.id
}

func (s *SimpleBlank) String() string {
	return s.id
}

func NewSimpleBlank(id string) (*SimpleBlank, error) {
	if len(strings.TrimSpace(id)) == 0 {
		return nil, errors.New("blank id")
	}
	return &SimpleBlank{id: id}, nil
}

////////////////////////////////////////////////////////////////////////////////
// type SimpleIRI
////////////////////////////////////////////////////////////////////////////////

type SimpleIRI struct {
	value string
}

func (*SimpleIRI) isSubjectType()   {}
func (*SimpleIRI) isPredicateType() {}
func (*SimpleIRI) isObjectType()    {}

func (s *SimpleIRI) Value() string {
	return s.value
}

func (s *SimpleIRI) String() string {
	return s.value
}

func NewSimpleIRI(value string) (*SimpleIRI, error) {
	err := ValidateIRI(value)
	if err != nil {
		return nil, err
	}
	return &SimpleIRI{value: value}, nil
}

////////////////////////////////////////////////////////////////////////////////
// type SimpleLiteral
////////////////////////////////////////////////////////////////////////////////

// Literal ...
type SimpleLiteral struct {
	value    string
	language *string
	dataType IRI
}

func (*SimpleLiteral) isObjectType() {}

func (s *SimpleLiteral) Value() string {
	return s.value
}

func (s *SimpleLiteral) Language() *string {
	return s.language
}

func (s *SimpleLiteral) DataType() IRI {
	return s.dataType
}

func NewSimpleLiteral(value string, language *string, dataType IRI) (*SimpleLiteral, error) {
	return &SimpleLiteral{value: value, language: language, dataType: dataType}, nil
}
