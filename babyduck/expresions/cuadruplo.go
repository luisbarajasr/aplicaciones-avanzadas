package expresions

import (
	"babyduck/semantic"
	"babyduck/structures"
)

// cuadruplo 
type Cuadruplo struct {
	Arg1     semantic.Variable
	Arg2     *semantic.Variable // siendo un puntero, permite ser nil, en caso de que no se necesite, eg. asignar valores
	Operator semantic.Operator
	Result   semantic.Variable
}

type CuadruploList struct {
	Cuadruplos []Cuadruplo
}

var OpStack = structures.OpStack{}

var VarStack = structures.Stack{}

var TempVariables = map[string]semantic.Variable

// regresa las precednecias de los operadores
func (o Operator) Precedence() int {
	switch o {
	case NewPara:
		return 4
	case Times, Divide:
		return 3
	case Plus, Minus:
		return 2
	case Less, Greater, NotEqual:
		return 1
	case Assign:
		return 0
	default:
		return -1 // operador invalido
	}
}

func (operator Operator) IsRightAssociative() bool {
    return operator == Assign
}

func NewTempVariable(id string, typ semantic.Type) semantic.Variable {
    return semantic.Variable{
        Id:          id,
        Type:        typ
    }
}

// ------------ CUADRUPLO LIST ------------ 

func NewCuadruploList() CuadruploList {
	return CuadruploList{
		Cuadruplos: []Cuadruplo{},
	}
}

func AddOperator(opertator semantic.Operator) (semantic.Operator, error){
	
	if OpStack.IsEmpty() {
		// agregar operador a la pila
		OpStack.Stack.Push(opertator)
		return operator, nil

	} 
	else if(operador == semantic.NewPara){
		OpStack.Stack.Push(opertator)
		return operator, nil

	} else if (operator == semantic.ClosePara) {

		for OpStack.Stack.Peek() != semantic.NewPara {
			topOperator := OpStack.Stack.Pop()
			AddQuadruple(topOperator)
			OpStack.Stack.Pop()
		}

		OpStack.Stack.Pop() // sacar el falso stack
		return OpStack.Peek(), nil

	}else if (operator.IsRightAssociative() && operator.Precedence() > OpStack.Stack[len(OpStack.Stack)-1].Precedence()) || (!operator.IsRightAssociative() && operator.Precedence() >= OpStack.Stack[len(OpStack.Stack)-1].Precedence()) {
		OpStack.Stack.Push(opertator)
		return operator, nil

	} else if operator.Precedence() < OpStack.Stack[len(OpStack.Stack)-1].Precedence() {
		// sacar el operador de la cima de la pila
		topOperator := OpStack.Stack.Pop()

		// agergar el operador a la lista de cuadruplos
		AddQuadruple(topOperator)

		// agregar el nuevo operador a la pila
		OpStack.Stack.Push(opertator)
		return operator, nil
	}
	
	return operator, nil

	/* 
	revisar si esta vacio 
		vacio, agregar operador
	
	revisar si operador es parentesis
		crear nuevo stack falso
		agregar operador
	
	revisar si operador es cierre de parentesis
		mientras el operador de la cima del stack falso no sea el de apertura:
			sacar los operadores dentro del stack falso
			agregar a la lista de cuadruplos
		ya encontro el operdaor de apertura (sale del loop)
		sacar el operador de apertura del stack falso

	no vacio, revisar si el nuevo operador tiene mayor o igual precedencia que el de la cima
		menor, agregar operador
		igual, agregar operador
		mayor, sacar de la pila el operador de la cima y agregarlo a la lista de cuadruplos
			agregar el nuevo operador a la pila
		
	*/
}

func addVariable(variable semantic.Variable){
	VarStack.Stack.Push(variable)
}

func addCuadruplo(operator semantic.Operator) {
	var1, var2 := VarStack.Stack[len(VarStack.Stack)-2], VarStack.Stack[len(VarStack.Stack)-1]

	var opResult semantic.Type = SemanticCube[var1.Type][var2.Type][operator]
	if opResult == semantic.Error {
		panic("Error: Type mismatch in operation")
	}

	// que pasa si es un asignacion?

	var cuadruplo Cuadruplo = Cuadruplo{
		Arg1:     var1,
		Arg2:     &var2,
		Operator: operator,
		Result:   NewTempVariable("t"+strconv.Itoa(len(TempVariables)), opResult),
	}

	CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)
	VarStack.Pop() // quitar las 2 variables de la pila
	VarStack.Pop() 
	VarStack.Push(cuadruplo.Result) // agregar el resultado al stack
	TempVariables[cuadruplo.Result.Name] = opResult
}

func PrintCuadruplos() {
	for _, cuadruplo := range CuadruploList.Cuadruplos {
		println(cuadruplo.Arg1.Name, cuadruplo.Operator, cuadruplo.Arg2.Name, "->", cuadruplo.Result.Name)
	}
	println("Total Cuadruplos:", len(CuadruploList.Cuadruplos))
}