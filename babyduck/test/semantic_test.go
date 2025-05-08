package main

import (
	"babyduck/lexer"
	"babyduck/parser"
	"testing"
)

// testData maps input programs to a boolean indicating whether parsing should succeed.
var testData = map[string]bool{
	// Test case 1: testing empty program
	`program demoEight;

		var x, y, z : int;

		void anotherFunction(a : int, b : float) [
			{
				x = 2;
			}
		];

		void anotherFunction(abc : float, bca : int) [
			var c : float;

			{
				c = 3;
			}
		];

		main {
			print(x);
		}

		end`: false,
	// Test case 2: testing variable declaration and assignment
	// "program p; var x: int; main { x = 5.2; } end": true,
	// // // Test case 3: testing float
	// "program p; var y: float; main { y = 3.14; } end": true,
	// // // Test case 4: testing if
	// "program p; var z: int; main { if (z < 10) { z = 2 + 1; }; } end": true,
	// // // Test case 6: testing print
	// "program p; var nueva: int; main { nueva = 5; print(nueva); } end": true,
	// // Test case 8: testing while
	// `program p; var r: int; main 
	// { while (r < 10) do { r = r + 1; }; } 
	// end`: true,
	// // Test case 9: missing semicolon after identifier
	// "program p main { } end": false,
	// // Test case 10: missing 'end'
	// "program p; main { }": false,
	// // Test case 11: missing semicolon after assignment
	// "program p; var e: int; main { e = 5 } end": false,
	// // Test case 13: testing function declaration
	// "program p; var m: int; void f(a: int) [{ m = a + 2; }]; main { } end": true,
	// // Test case 12: extra text after 'end'
	// "program p; var z: int; main { z = 5; } end extra": false,
}

func TestParse(t *testing.T) {
	i := 1
	for input, ok := range testData {
		// Log the test input and output
        t.Logf("\n")
		t.Logf("=== Parsing Test #%d", i)
        t.Logf("Input:\n%s", input)
        t.Logf("\n")
		l := lexer.NewLexer([]byte(input))
		p := parser.NewParser()
		_, err := p.Parse(l)

		// Check expectation
		if (err == nil) != ok {
			if ok {
				t.Errorf("unexpected error parsing valid input:\n%s\nerror: %v", input, err)
				t.Logf("Parse error: %v", err)
			} else {
				t.Errorf("expected error parsing invalid input, but got none:\n%s", input)
				t.Logf("Parse error: %v", err)
			}
		}
		i++
	}
}