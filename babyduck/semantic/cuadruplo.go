package semantic

import (
	"strconv"
	"fmt"
)

// cuadruplo 
type Cuadruplo struct {
	Arg1     Variable
	Arg2     *Variable // siendo un puntero, permite ser nil, en caso de que no se necesite, eg. asignar valores
	Operator Operator
	Result   Variable
}

type CuadruploList struct {
	Cuadruplos []Cuadruplo
	functionDir *FunctionDirectory
}

var TempVariables = map[string]Variable{}

// regresa las precednecias de los operadores
func (o Operator) Precedence() int {
	switch o {
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

func (op Operator) IsRightAssociative() bool {
    return op == Assign
}

func NewTempVariable(name string, typ Type) Variable {
    return Variable{
        Name:          name,
        Type:        typ,
    }
}

// ------------ CUADRUPLO LIST ------------ 

var (
	opStack  *OpStack  // Global OpStack instance
	varStack *VarStack // Global VarStack instance
)

func NewCuadruploList(functionDir *FunctionDirectory) *CuadruploList {

	opStack = NewOpStack()   // Initialize OpStack
    varStack = NewVarStack() // Initialize VarStack

    return &CuadruploList{
        Cuadruplos: []Cuadruplo{},
		functionDir: functionDir,
    }
}

func (CuadruploList *CuadruploList) AddOperator(operator Operator) (Operator, error){

	fmt.Println("Adding operator to stack:", operator)

	if opStack.IsEmpty() {
		// agregar operador a la pila
		opStack.Push(operator)
		return operator, nil

	} else if(operator == NewPara){
		opStack.Push(operator)
		return operator, nil

	} else if (operator == ClosePara) {

		for opStack.Peek() != NewPara {
			opOperator,_ := opStack.Pop()
			CuadruploList.addCuadruplo(opOperator)
			opStack.Pop()
		}

		opStack.Pop() // sacar el falso stack
		return opStack.Peek(), nil

	} else if ( operator.Precedence() > opStack.Peek().Precedence() || operator.Precedence() == opStack.Peek().Precedence()) {
		opStack.Push(operator)
		return operator, nil

	} else if (operator.Precedence() < opStack.Peek().Precedence()) {
		// sacar el operador de la cima de la pila
		topOperator,_ := opStack.Pop()

		// agergar el operador a la lista de cuadruplos
		CuadruploList.addCuadruplo(topOperator)

		// agregar el nuevo operador a la pila
		opStack.Push(operator)
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

func (CuadruploList *CuadruploList) addVariable(variable Variable){
	fmt.Println("Adding variable to stack:", variable.Name)
	varStack.Push(variable)
	varStack.Print()
}

func (CuadruploList *CuadruploList) addCuadruplo(operator Operator) {
	var1, verifier := varStack.PeekDouble()
	if !verifier {
		panic("Error: Not enough variables in stack")
	}
	var2 := varStack.Peek() // tengo que corregir esto
	/* 
	entra una variable
	se lee un operador
	trata de tomar 2 variables del stack
		hay error porque solo esta una variable
	*/

	var opResult Type = SemanticCube[var1.Type][var2.Type][operator]
	if opResult == Error {
		panic("Error: Type mismatch in operation")
	}

	var cuadruplo Cuadruplo

	// que pasa si es un asignacion?
	if operator == Assign {
		cuadruplo = Cuadruplo{
			Arg1:     var2,
			Arg2:     nil,
			Operator: operator,
			Result:   var1,
		}
	} else {
		cuadruplo  = Cuadruplo{
			Arg1:     var1,
			Arg2:     &var2,
			Operator: operator,
			Result:   NewTempVariable("t"+strconv.Itoa(len(TempVariables)), opResult),
		}
	}

	CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)
	varStack.Pop() // quitar las 2 variables de la pila
	varStack.Pop() 
	varStack.Push(cuadruplo.Result) // agregar el resultado al stack
	TempVariables[cuadruplo.Result.Name] = Variable{
		Name: cuadruplo.Result.Name,
		Type: opResult,
	}
}

func (CuadruploList *CuadruploList) PrintCuadruplos() {
	for _, cuadruplo := range CuadruploList.Cuadruplos {
		println(cuadruplo.Arg1.Name, cuadruplo.Operator, cuadruplo.Arg2.Name, "->", cuadruplo.Result.Name)
	}
	println("Total Cuadruplos:", len(CuadruploList.Cuadruplos))
}