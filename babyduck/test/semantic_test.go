package main

import (
	"testing"

	// "babyduck/semantic"
	"babyduck/lexer"
	"babyduck/parser"
)

func Test1(t *testing.T) {
	src :=
		`program demoEleven;

		var a, b, c, s : int;

		main {
			s = (a + b) * c;
			a = b + c;
			b = c;

			if(a<b){
				a = b;
			};
		}
		end`

	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()

	tree, perr := p.Parse(l)

	if perr != nil {
		t.Fatalf("parse failed: %v", perr)
	}

	t.Logf("parse OK %#v", tree)
}

// func Test2(t *testing.T) {
// 	src :=
// 	`program demoTwelve;

// 	var a, b, c, z : int;

// 	void anotherFunction(p : int, q : float) [
// 		var o : int;

// 		{
// 			o = p + q;
// 			print(o);
// 		}
// 	];

// 	main {

// 	}

// 	end`

// 	l := lexer.NewLexer([]byte(src))
// 	p := parser.NewParser()

// 	tree, perr := p.Parse(l)

// 	if perr != nil {
// 		t.Fatalf("parse failed: %v", perr)
// 	}

// 	t.Logf("parse OK %#v", tree)
// }

// func Test3(t *testing.T) {
// 	src :=
// 	`program myprog;
// 	var id1 : int; id2 : float;
// 	main {
// 		id2 = 5;
// 		id1 = 5 + 10;
// 	}
// 	end`
	
// 	l := lexer.NewLexer([]byte(src))
// 	p := parser.NewParser()

// 	tree, perr := p.Parse(l)

// 	if perr != nil {
// 		t.Fatalf("parse failed: %v", perr)
// 	}

// 	t.Logf("parse OK %#v", tree)
// }

// func Test4(t *testing.T) {
// 	src :=
// 		`program demoEleven;

// 		var a, b, c, d, e : int; z : float;
// 		main {
// 			z = (a - (d / e)) * c;
// 		}
// 		end`

// 	l := lexer.NewLexer([]byte(src))
// 	p := parser.NewParser()

// 	tree, perr := p.Parse(l)

// 	if perr != nil {
// 		t.Fatalf("parse failed: %v", perr)
// 	}

// 	t.Logf("parse OK %#v", tree)
// }