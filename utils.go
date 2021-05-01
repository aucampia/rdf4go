package rdf4go

import "fmt"

func ValidateIRI(iri string) error {
	for _, r := range iri {
		if r >= '\x00' && r <= '\x20' {
			return fmt.Errorf("disallowed character: %q", r)
		}
		switch r {
		case '<', '>', '"', '{', '}', '|', '^', '`', '\\':
			return fmt.Errorf("disallowed character: %q", r)
		}
	}
	return nil
}
