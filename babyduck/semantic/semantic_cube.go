/ cubo semantico
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
			Less:        Int,   // boolean, but BabyDuck uses int for simplicity
			Greater:     Int,
			NotEqual:    Int,
			Assign:      Int,
		},
		Float: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Int,
			Greater:     Int,
			NotEqual:    Int,
			Assign:      Error, // int = float is invalid
		},
	},
	Float: {
		Int: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Int,
			Greater:     Int,
			NotEqual:    Int,
			Assign:      Error, // float = int is invalid
		},
		Float: {
			Plus:        Float,
			Minus:       Float,
			Times:       Float,
			Divide:      Float,
			Less:        Int,
			Greater:     Int,
			NotEqual:    Int,
			Assign:      Float,
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