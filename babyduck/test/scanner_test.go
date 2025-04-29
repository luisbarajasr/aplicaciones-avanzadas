package main

import (
	"fmt"
	"testing"
	"babyduck/lexer"
	"babyduck/parser"
)

func TestParser(t *testing.T) {
	ex_array := []string{
		`program demo;

		var  variable, y, z : float;

		main {
			print(1 + 2);
		}
		end`,
		`program programa_demo;

		var  x, y, z : int; p, e, o : float;

		void funcion_uno(a : int, b : float) [
			var c : int;

			{
				d = a + b;
				print("el resultado es d");
			}
		];

		main {
			funcion1(3,5);

			while (x < 10) do {
				print(x);
				x = x + 1;
			};
		}
		end`,
		`
		program while_test;
		var count: int;
		main {
			count = 0;
			while (count < 5) do {
				print("looping");
				count = count + 1;
			};
		}
		end
		`,
		`
		program if_test;
		var x: int;
		main {
			x = 15;
			if (x > 20) {
				print("x is big");
			}
			else {
				print("x is small");
			};		
		}
		end
		`,

	}

	for idx, input := range ex_array {
		t.Run(fmt.Sprintf("TestCase%d", idx+1), func(t *testing.T) {
			l := lexer.NewLexer([]byte(input))
			p := parser.NewParser()
			_, err := p.Parse(l)
			if err != nil {
				t.Fatalf("TestCase%d - parser error: %v", idx+1, err)
			}
		})
	}
}