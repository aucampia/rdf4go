// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

package rdf4go

////////////////////////////////////////////////////////////////////////////////
// type SimpleBnode
////////////////////////////////////////////////////////////////////////////////

type SimpleBnode struct {
	id string
}

func (*SimpleBnode) isSubjectType() {}
func (*SimpleBnode) isObjectType()  {}

func (s *SimpleBnode) Id() string {
	return s.id
}

func (s *SimpleBnode) String() string {
	return s.id
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

////////////////////////////////////////////////////////////////////////////////
// type SimpleLiteral
////////////////////////////////////////////////////////////////////////////////

// Literal ...
type SimpleLiteral struct {
	value    string
	language *string
	dataType *IRI
}

func (*SimpleLiteral) isObjectType() {}

func (s *SimpleLiteral) Value() string {
	return s.value
}

func (s *SimpleLiteral) Language() *string {
	return s.language
}

func (s *SimpleLiteral) DataType() *IRI {
	return s.dataType
}

////////////////////////////////////////////////////////////////////////////////
// type SimpleValueFactory
////////////////////////////////////////////////////////////////////////////////

type SimpleValueFactory struct {
}

func (f *SimpleValueFactory) NewBNode(id string) (*SimpleBnode, error) {
	return &SimpleBnode{id: id}, nil
}

func (f *SimpleValueFactory) NewIRI(value string) (*SimpleIRI, error) {
	err := ValidateIRI(value)
	if err != nil {
		return nil, err
	}
	return &SimpleIRI{value: value}, nil
}

func (f *SimpleValueFactory) NewLiteral(value string, language *string, dataType *IRI) (*SimpleLiteral, error) {
	return &SimpleLiteral{value: value, language: language, dataType: dataType}, nil
}
