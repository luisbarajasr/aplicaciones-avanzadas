// cubo semantico
// This file contains the semantic cube for BabyDuck, which defines the
// valid operations and their resulting types for the BabyDuck programming language.
package semantic

import (
	"fmt"
)

var SemanticCube = map[Type]map[Type]map[Operator]Type{
	Int: {
		Int: {
			Plus:        Int,
			Minus:       Int,
			Times:       Int,
			Divide:      Float, // int / int -> float (to handle division)
			Less:        Bool,
			Greater:     Bool,
			NotEqual:    Bool,
			Assign:      Int,
		},
		Float: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Bool,
			Greater:     Bool,
			NotEqual:    Bool,
			Assign:      Error,
		},
		Bool: {
			Plus:     Error,
			Minus:    Error,
			Times:    Error,
			Divide:   Error,
			Less:     Error,
			Greater:  Error,
			NotEqual: Error,
			Assign:      Error,
		},
	},
	Float: {
		Int: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Bool,
			Greater:     Bool,
			NotEqual:    Bool,
			Assign:      Error,
		},
		Float: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Bool,
			Greater:     Bool,
			NotEqual:    Bool,
			Assign:      Float,
		},
		Bool: {
			Plus:     Error,
			Minus:    Error,
			Times:    Error,
			Divide:   Error,
			Less:     Error,
			Greater:  Error,
			NotEqual: Error,
			Assign:      Error,
		},
	},
	Bool: {
		Int: {
			Plus:     Error,
			Minus:    Error,
			Times:    Error,
			Divide:   Error,
			Less:     Error,
			Greater:  Error,
			NotEqual: Bool,
			Assign:      Error,
		},
		Float: {
			Plus:     Error,
			Minus:    Error,
			Times:    Error,
			Divide:   Error,
			Less:     Error,
			Greater:  Error,
			NotEqual: Error,
			Assign:      Error,
		},
		Bool: {
			Plus:     Error,
			Minus:    Error,
			Times:    Error,
			Divide:   Error,
			Less:     Error,
			Greater:  Error,
			NotEqual: Error,
			Assign:   Error,
		},
	},
}

// CheckTypes returns the result type for an operation or error if invalid
func CheckTypes(t1, t2 Type, op Operator) (Type, error) {
	if t1Map, exists := SemanticCube[t1]; exists {
		if t2Map, exists := t1Map[t2]; exists {
			if resultType, exists := t2Map[op]; exists {
				if resultType == Void {
					return Void, fmt.Errorf("type mismatch: %s %s %s is invalid", t1, op, t2)
				}
				return resultType, nil
			}
		}
	}
	return Void, fmt.Errorf("unsupported operation: %s %s %s", t1, op, t2)
}