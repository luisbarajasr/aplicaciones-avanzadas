package semantic

import (
	"strconv"
	// "fmt"
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

func (CuadruploList *CuadruploList) addVariable(variable Variable){
	// fmt.Println("Adding variable to stack:", variable.Name)
	varStack.Push(variable)
	// varStack.Print()
}

func (CuadruploList *CuadruploList) addCuadruplo_dos(operator Operator) {
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

func (CuadruploList *CuadruploList) AddOperator(operator Operator) (Operator, error) {
	// fmt.Println("Adding operator to stack:", operator)

	if opStack.IsEmpty() {
		opStack.Push(operator)
		return operator, nil
	}

	switch operator {
		case NewPara:
			opStack.Push(operator)
			return operator, nil

		case ClosePara:
			for !opStack.IsEmpty() && opStack.Peek() != NewPara {
				topOp, _ := opStack.Pop()
				CuadruploList.addCuadruplo(topOp)	
			}
			if !opStack.IsEmpty() {
				opStack.Pop() // Remove the NewPara operator
			}

			return operator, nil

		default:
			currentPrecedence := operator.Precedence()
			topPrecedence := opStack.Peek().Precedence()

			if currentPrecedence > topPrecedence {
				opStack.Push(operator)
			} else {
				for !opStack.IsEmpty() && currentPrecedence <= opStack.Peek().Precedence() {
					topOp, _ := opStack.Pop()
					CuadruploList.addCuadruplo(topOp)
				}
				opStack.Push(operator)
			}
			return operator, nil
	}
} 

func (CuadruploList *CuadruploList) addCuadruplo(operator Operator) {

	switch operator {

	case Assign:
		result := varStack.Pop()
		varTarget := varStack.Pop()

		cuadruplo := Cuadruplo{
			Operator: operator,
			Arg1:     result,
			Result:   varTarget,
		}	
		CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)

	default:
		arg2 := varStack.Pop()
		arg1 := varStack.Pop()
		
		var opResult Type = SemanticCube[arg1.Type][arg2.Type][operator]
		if opResult == Error {
			panic("Error: Type mismatch in operation")
		}
		
		tempVar := NewTempVariable("t"+strconv.Itoa(len(TempVariables)), opResult)

		cuadruplo := Cuadruplo{
			Operator: operator,
			Arg1:     arg1,
			Arg2:     &arg2,
			Result:   tempVar,
		}

		varStack.Push(tempVar) // Push the result variable onto the stack
		CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)
		TempVariables[tempVar.Name] = Variable{
			Name: tempVar.Name,
			Type: opResult,
		}
	}
}