package main

import (
	"babyduck/lexer"
	"babyduck/parser"
	"testing"
)

var testData = map[string]bool{
	// caso 1 : funcion doble declarada
	`program demoEleven;

		var a, b, c, z : int;

		main {
			z = (a + b) * c;
		}
		end`: true,
	// caso 2 : variable no declarada
    // `program demo2;

    //     var x, y, z : int;

    //     main {
    //         noesta = 2;
    //     }

    // end`: false,
	// // caso 3: variable declarada
    // `program demo3;

    // var t, u, i : int;

    // main {
    //     print(t);
    // }
    // end`: true,
	// // caso 4: re-declaracion de variable global
	// `program demo4;

	// 	var x, y, z : int;

	// 	void anotherFunction(a : int, b : float) [
	// 		var x : int;

	// 		{
	// 			d = a + b;
	// 			print(d);
	// 		}
	// 	];

	// 	main {
	// 		print(x);
	// 	}

	// 	end`: false,
	// // caso 5: variable no declarada
	// `program demo5;

	// 	var x, y, z : int;

	// 	void anotherFunction(a : int, b : float) [
	// 		var nueva : int;

	// 		{
	// 			nose = 1 + 2;
	// 			print(nose);
	// 		}
	// 	];

	// 	main {
	// 		x = 1;
	// 		anotherFunction(1, 2.0);
	// 	}

	// 	end`: false,
	// // caso 6: token no registrado
    // `program demo6;

    // var  x, y, z : string;

    // main {
    //     print(1 + 2);
    // }`: false,
	// // caso 7: variable no declarada
	// "program p main { } end": false,
}

func TestParse(t *testing.T) {
	i := 1
	for input, ok := range testData {
		// Log the test input and output
        t.Logf("\n")
		t.Logf("=== Parsing Test #%d", i)
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