package semantic

import "fmt"

type Type int

const (
    Int Type = iota
    Float
    Error
)

var syntaxCube = map[Type]map[Type]map[string]Type{
    Int: {
        Int: {
            "+":  Int,
            "-":  Int,
            "*":  Int,
            "/":  Float, // int / int -> float
            ">":  Int,
            "<":  Int,
            "!=": Int,
            "=":  Int,
        },
        Float: {
            "+":  Float,
            "-":  Float,
            "*":  Float,
            "/":  Float,
            ">":  Int,
            "<":  Int,
            "!=": Int,
            "=":  Error, // int = float is invalid
        },
    },
    Float: {
        Int: {
            "+":  Float,
            "-":  Float,
            "*":  Float,
            "/":  Float,
            ">":  Int,
            "<":  Int,
            "!=": Int,
            "=":  Error, // float = int is invalid
        },
        Float: {
            "+":  Float,
            "-":  Float,
            "*":  Float,
            "/":  Float,
            ">":  Int,
            "<":  Int,
            "!=": Int,
            "=":  Float,
        },
    },
}

func stringToType(s string) Type {
    switch s {
    case "int":
        return Int
    case "float":
        return Float
    default:
        return Error
    }
}

func CheckTypes(t1, t2, op string) (string, error) {
    t1Type := stringToType(t1)
    t2Type := stringToType(t2)
    if t1Map, exists := syntaxCube[t1Type]; exists {
        if t2Map, exists := t1Map[t2Type]; exists {
            if resultType, exists := t2Map[op]; exists {
                if resultType == Error {
                    return "", fmt.Errorf("type mismatch: %s %s %s", t1, op, t2)
                }
                return typeToString(resultType), nil
            }
        }
    }
    return "", fmt.Errorf("unsupported operation: %s %s %s", t1, op, t2)
}

func typeToString(t Type) string {
    switch t {
    case Int:
        return "int"
    case Float:
        return "float"
    default:
        return "error"
    }
}