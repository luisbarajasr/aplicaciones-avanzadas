package main

import (
	"testing"
	// "fmt"
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

// func Test3(t *testing.T) {
// 	src :=
// 	`program demoTwelve;

// 	var a, b, c, z : int; nose, idk : float;

// 	void uno(x : int, y : int) [{
// 		y = x + 2;
// 		a = x + y;
// 	}];

// 	void dos(dos : float, tres : float) [{
// 		tres = dos + a;
// 	}];

// 	main {
// 		uno(a*b-z, b);

// 		dos(nose, idk+3.12);

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

// func Test3(t *testing.T) {
// 	src :=
// 	`program myprog;
// 	var id1 : int; id2 : float;
// 	main {
// 		print("hola", id2, id1);
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

func Test4(t *testing.T) {
	src :=
		`program demoEleven;
		var a, e : int; z : float;

		main {
			print("Hola mundo:", a);
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

// func Test5(t *testing.T) {
// 	src :=
// 	`	program primerintento; 
	
// 	var x: int;
	
// 	void uno(a: int) 
// 	[ 
// 		{
// 			while (a < 10) do
// 			{
// 				print("El valor de a es: ", a);
// 				a = a + 1;
// 			};
// 		}
// 	]; 

// 	void dos() 
// 	[ 
// 		var y: int;
// 		{
// 			y = 10;
// 			print(y);
// 			if (x > 5) 
// 			{
// 				print ( "X es mayor que 5" );
// 			}
// 			else
// 			{
// 				print ( "X es menor a 5" );
// 			};
// 		}
// 	]; 
	
// 	main { 
// 		x = 3;
// 		print('Hola mundo!');
// 		uno(x);
// 		dos();
// 		print(21/3);
// 		print('Fin del programa');
// 	} end`
	
// 	l := lexer.NewLexer([]byte(src))
// 	p := parser.NewParser()

// 	tree, perr := p.Parse(l)

// 	if perr != nil {
// 		t.Fatalf("parse failed: %v", perr)
// 	}

// 	t.Logf("parse OK %#v", tree)
// }