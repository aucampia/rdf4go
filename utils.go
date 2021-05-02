// SPDX-FileCopyrightText: 2021 Iwan Aucamp
//
// SPDX-License-Identifier: CC0-1.0 OR Apache-2.0

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
