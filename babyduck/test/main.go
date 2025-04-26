package main

import (
	"fmt"
	"os"

	"babyduck/lexer"
	"babyduck/parser"
)

func main() {
	source := `
	program myProgram;
	var x, y : int;
	void myFunc(x : int, y : float)[
		var z : float;
		{
			x = y;
			print(x, y, z);
		}
	];
	end
	`

	l := lexer.NewLexer([]byte(source))
	p := parser.NewParser()

	ast, err := p.Parse(l)
	if err != nil {
		fmt.Println("Parse error:", err)
		os.Exit(1)
	}
	fmt.Println("Parse successful!")
	fmt.Printf("AST: %#v\n", ast)
}
