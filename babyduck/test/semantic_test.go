package main

import (
	"testing"
	// "fmt"
	"babyduck/lexer"
	"babyduck/parser"
)

// func Test1(t *testing.T) {
// 	src :=
// 	`
// 	program fibonacci;
// 	var input: int;
// 	void funcion2(nose: int) [
// 		{
// 			print("imprimiendo desde funcion2", nose);
// 		}
// 	];

// 	void funcion1(n: int)[{
// 		funcion2(n);
// 	}];

// 	main {
// 		input = 10;
// 		funcion1(input);
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

// func Test4(t *testing.T) {
// 	src :=
// 		`program demoEleven;
// 		var e : int; a, z : float;

// 		main {
// 			a = 3 + 2.2 ;
// 			print(a);
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

func Test5(t *testing.T) {
	src :=
	`	program factorial;
		var input: int;

		void factorial(n: int) [
			var result, i: int;
			{
				if (n < 2) {
					print("Factorial of", n, "is", 1);
				} else {
					result = 1;
					i = 2;
					while (i < n+1) do {
						result = result * i;
						i = i + 1;
					};
					print("Factorial of", n, "is", result);
				};
			}
		];

		main {
			input = 5; 
			factorial(input);
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

// func Test6(t *testing.T) {
// 	src :=
// 	`
// 	program fibonacci;
// 	var input: int;

// 	void fibonacci(n: int) [
// 	var a,b,count,temp : int;
// 	{
// 		if (n < 2) {
// 			print("Fibonacci of", n, "is", b);
// 		} else {
// 			a = 0;
// 			b = 1;
// 			count = 2;
// 			while (count < n+1) do {
// 				temp = a+b;
// 				a = b;
// 				b = temp;
// 				count = count + 1;
// 			};
// 			print("Fibonacci of", n, "is", b);
// 		};
// 	}];
// 	main {
// 		input = 10;
// 		fibonacci(input);
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


// func Test6(t *testing.T) {
// 	src :=
// 	`
// 	program fibonacci;
// 	var input: int;
// 	void print_fibonacci(num: int, b:int) [
// 		{
// 			print("Fibonacci of", num, "is", b);
// 		}
// 	];

// 	void fibonacci(n: int) [
// 	var a,b,count,temp : int;
// 	{
// 		if (n < 2) {
// 			print_fibonacci(n,1);
// 		} else {
// 			a = 0;
// 			b = 1;
// 			count = 2;
// 			while (count < n+1) do {
// 				temp = a+b;
// 				a = b;
// 				b = temp;
// 				count = count + 1;
// 			};
// 			print_fibonacci(n,b);
// 		};
// 	}];
// 	main {
// 		input = 10;
// 		fibonacci(input);
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