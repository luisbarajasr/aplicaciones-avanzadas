package main

import (
	"testing"

	// "babyduck/semantic"
	"babyduck/lexer"
	"babyduck/parser"
)

// func Test1(t *testing.T) {
// 	src :=
// 		`program demoEleven;

// 		var a, b, c, s : int;

// 		main {
// 			s = (a + b) * c;
// 			a = b + c;
// 			b = c;

// 			if(a<b*c){
// 				a = b;
// 			}else{
// 				b = c;	
// 			};
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

// func Test2(t *testing.T) {
// 	src :=
// 	`program demoTwelve;

// 	var a, b, c, z : int; nose : float;

// 	void uno(x : int, y : int) [{
// 		y = x + b;
// 		a = x + y;
// 		nose = a+b; 
// 	}];

// 	void dos(dos : float, tres : float) [{
// 		tres = dos + a;
// 	}];

// 	main {
// 		uno(a*b-z, b);

// 		dos(nose, nose);

// 		while(a < b)
// 		do {
// 			a = a + b;
// 		};

// 		c = a + b;
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

func Test3(t *testing.T) {
	src :=
	`program myprog;
	var id1 : int; id2 : float;
	main {
		print("hola", id2, id1);
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