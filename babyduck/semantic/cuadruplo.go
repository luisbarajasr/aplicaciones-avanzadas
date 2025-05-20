package semantic

import ()

// cuadruplo 
type Cuadruplo struct {
	Arg1     int
	Arg2     int // siendo un puntero, permite ser nil, en caso de que no se necesite, eg. asignar valores
	Operator Operator
	Result   int
}

type CuadruploList struct {
	Cuadruplos []Cuadruplo
	functionDir *FunctionDirectory
}

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

func NewTempVariable(name string, typ Type, address int) Variable {
    return Variable{
        Name:          name,
        Type:        typ,
		Virtual_address: address,
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

// fmt.Println("Adding variable to stack:", variable.Name)
func (CuadruploList *CuadruploList) addVariable(variable Variable){
	varStack.Push(variable.Virtual_address)
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

// AddOperator agrega un operador al stack de operadores y genera un cuadruplo si es necesario 
// dependiendo de la precedencia del operador y el operador en la cima del stack
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

// AddCuadruplo agrega un cuadruplo a la lista de cuadruplos
// y genera un nuevo temporal si es necesario
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
			
			resultType,_ := CuadruploList.functionDir.GetTypeByAddress(result)
			varTargetType,_ := CuadruploList.functionDir.GetTypeByAddress(varTarget)

			if validateAssignment(resultType, varTargetType) != nil {
				panic("Error: Type mismatch in assignment")
			}

			cuadruplo := Cuadruplo{
				Operator: operator,
				Arg1:     result,
				Result:   varTarget,
			}

			CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)
			// fmt.Println("Cuadruplo:", cuadruplo.Operator, cuadruplo.Arg1.Name, "->", cuadruplo.Result.Name)
			
		default:
		arg2 := varStack.Pop()
		arg1 := varStack.Pop()

		arg1Type,_:= CuadruploList.functionDir.GetTypeByAddress(arg1)
		arg2Type,_ := CuadruploList.functionDir.GetTypeByAddress(arg2)
		
		var opResult,_ = CheckTypes(arg1Type, arg2Type, operator)
		// var opResult Type = SemanticCube[arg1.Type][arg2.Type][operator]
		if opResult == Error {
			panic("Error: Type mismatch in operation")
		}

		resultVarType := varTypetoMemoryType(opResult) // convierte el tipo de variable a un tipo de memoria para la asignacion de dir correcta 
		resultAddress := CuadruploList.functionDir.MemoryManager.Allocate(resultVarType) // sera un temporal de int, float o bool
		
		cuadruplo := Cuadruplo{
			Operator: operator,
			Arg1:     arg1,
			Arg2:     arg2,
			Result:   resultAddress,
		}

		varStack.Push(resultAddress) // Push the result variable onto the stack
		CuadruploList.Cuadruplos = append(CuadruploList.Cuadruplos, cuadruplo)
	}
}

func varTypetoMemoryType(varType Type) MemoryType {
	switch varType {
	case Int:
		return TemporalInt
	case Float:
		return TemporalFloat
	case Bool:
		return TemporalBool
	}
	return GlobalInt // Default case, should not happen
}