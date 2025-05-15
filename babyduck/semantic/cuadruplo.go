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

func (CuadruploList *CuadruploList) addVariable(variable Variable){
	// fmt.Println("Adding variable to stack:", variable.Name)
	varStack.Push(variable)
	// varStack.Print()
}

func (cl *CuadruploList) PrintCuadruplos() {
    if cl == nil || cl.Cuadruplos == nil {
        fmt.Println("No quadruples generated")
        return
    }
    
    for _, quad := range cl.Cuadruplos {
        if quad.Arg2 != nil {
            fmt.Printf("%s %s %s -> %s\n", 
                quad.Arg1.Name, quad.Operator, quad.Arg2.Name, quad.Result.Name)
        } else {
            fmt.Printf("%s %s -> %s\n", 
                quad.Arg1.Name, quad.Operator, quad.Result.Name)
        }
    }
    fmt.Printf("Total Cuadruplos: %d\n", len(cl.Cuadruplos))
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
			if !opStack.IsEmpty() && opStack.Peek() == NewPara {
				opStack.Pop() // Remove the NewPara operator
			}	
			
			// fmt.Println("Closing parenthesis, popping operators until NewPara")
			// opStack.Print()

			return operator, nil

		case Semicolon:
			// fmt.Println("Semicolon detected, popping operators until empty stack")

			for !opStack.IsEmpty() {
				// opStack.Print()
				// fmt.Println("-----------------")
				topOp, _ := opStack.Pop()
				CuadruploList.addCuadruplo(topOp)
			}

			return operator, nil

		default:
			currentPrecedence := operator.Precedence()
			topPrecedence := opStack.Peek().Precedence()

			if currentPrecedence > topPrecedence {
				// fmt.Println("Adding operator to stack:", operator)
				opStack.Push(operator)
				// opStack.Print()
				// fmt.Println("-----------------")
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
	// fmt.Println("Adding operator to cuadruple list:", operator)

	// opStack.Print()
	// fmt.Println("-----------------")
	// varStack.Print()

	switch operator {

		case Assign:
			// fmt.Println("Adding assignment operator to stack:", operator)
			// fmt.Println("VarStack:", varStack)
			// fmt.Println("OpStack:", opStack)

			result := varStack.Pop()
			varTarget := varStack.Pop()
			// fmt.Println("Assigning", result.Name, "to", varTarget.Name)
			cuadruplo := Cuadruplo{
				Operator: operator,
				Arg1:     result,
				Result:   varTarget,
			}	
			// fmt.Println("Cuadruplo:", cuadruplo.Operator, cuadruplo.Arg1.Name, "->", cuadruplo.Result.Name)
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