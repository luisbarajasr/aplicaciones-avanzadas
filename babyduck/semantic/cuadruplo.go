package semantic

import (
	"fmt"
)

/* 
TODO:
- Manejo de constantes
- GOTOs
- Print*
*/

var (
    opStack       *OpStack
    varStack      *VarStack
    jumpStack     *JumpStack  // Nueva pila para manejar saltos
	BooleanTempVariables *VarStack
	TempVariables *VarStack
	whileIndex int
)

type JumpStack struct {
    stack []int
}

func NewJumpStack() *JumpStack {
    return &JumpStack{
        stack: make([]int, 0),
    }
}

func (js *JumpStack) Push(value int) {
    js.stack = append(js.stack, value)
}

func (js *JumpStack) Pop() (int, error) {
    if len(js.stack) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    value := js.stack[len(js.stack)-1]
    js.stack = js.stack[:len(js.stack)-1]
    return value, nil
}

func (js *JumpStack) Peek() (int, error) {
    if len(js.stack) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    return js.stack[len(js.stack)-1], nil
}


// cuadruplo 
type Cuadruplo struct {
	Arg1     int
	Arg2     int 
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
		return 5
	case Plus, Minus:
		return 4
	case Less, Greater, NotEqual:
		return 3
	case Assign:
		return 1
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

func NewCuadruploList(functionDir *FunctionDirectory) *CuadruploList {
    opStack = NewOpStack()
    varStack = NewVarStack()
    jumpStack = NewJumpStack()
	BooleanTempVariables = NewVarStack()
	TempVariables = NewVarStack()
	whileIndex = 0

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

    fmt.Println("\n=== Cuádruplos Generados ===")
    fmt.Printf("%-6s %-10s %-10s %-10s %-10s\n", "Pos", "Operador", "Arg1", "Arg2", "Resultado")
    fmt.Println("--------------------------------------------------")
    
    for i, quad := range cl.Cuadruplos {
        // Obtener información de cada argumento
        arg1Info := cl.getVariableInfo(quad.Arg1)
        arg2Info := ""
        if quad.Arg2 != 0 { // 0 indica que no hay segundo argumento
            arg2Info = cl.getVariableInfo(quad.Arg2)
        }
        resultInfo := cl.getVariableInfo(quad.Result)

        // Imprimir el cuádruplo formateado
        fmt.Printf("%-6d %-10v %-10s %-10s %-10s\n", 
            i, quad.Operator, arg1Info, arg2Info, resultInfo)
    }
    fmt.Printf("\nTotal Cuádruplos: %d\n", len(cl.Cuadruplos))
}

// Función auxiliar para obtener información de la variable
func (cl *CuadruploList) getVariableInfo(address int) string {
    // Buscar en variables globales
    for _, v := range cl.functionDir.GlobalVars.Variables {
        if v.Virtual_address == address {
            return v.Name
        }
    }

    // Buscar en variables locales de la función actual
    if cl.functionDir.CurrentFunction != nil {
        for _, v := range cl.functionDir.CurrentFunction.Vars.Variables {
            if v.Virtual_address == address {
                return v.Name
            }
        }
    }

    // Si no se encuentra, mostrar la dirección como tal
    return fmt.Sprintf("%d", address)
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
					// fmt.Println("Poping op from stack:", operator)
					opStack.Print()
					// fmt.Println("-----------------")
					// fmt.Println("WHOLE OP STACK:", opStack)
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
	// fmt.Println("temporal variable:", TempVariables)

	switch operator {
		
		case Assign:
			// fmt.Println("Adding assignment operator to stack:", operator)
			// fmt.Println("OpStack:", opStack)
			

			result := varStack.Pop()
			varTarget := varStack.Pop()
			

			// resultName := CuadruploList.getVariableInfo(result)
			// varTargetName := CuadruploList.getVariableInfo(varTarget)
			// fmt.Println("Assigning", resultName, "to", varTargetName)
			
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
		// fmt.Println("Operator:", operator, "Arg1:", arg1Type, "Arg2:", arg2Type, "Result Type:", opResult)

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

		if operator != Greater && operator != Less && operator != NotEqual {
			varStack.Push(resultAddress) // los temporales booleanos no se guardan en la pila de variables (causa problema en asignaciones)
		} else {
			BooleanTempVariables.Push(resultAddress) // los temporales booleanos se guardan en la pila de booleanos
		}
		TempVariables.Push(resultAddress)
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

// ----------- IF-ELSE ---------------
func (cl *CuadruploList) BeginIf() {
    jumpStack.Push(len(cl.Cuadruplos))
	// Genera el GOTOF (salto si falso)
	condition := BooleanTempVariables.Pop()
	cuadruplo := Cuadruplo{
		Operator: GOTOF,
		Arg1:     condition, // dirección del booleano temporal dentro del If( )
		Arg2:     0, // Se completará después
		Result:   0, 
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	jumpStack.Push(len(cl.Cuadruplos) - 1) // Guarda la posición del GOTOF
}

func (cl *CuadruploList) CompleteIf() {
	// Completa el GOTOF con la posición actual (fin del bloque if)
    gotoFIndex, _ := jumpStack.Pop()
	cl.Cuadruplos[gotoFIndex].Arg1 = BooleanTempVariables.Pop() // direccion del booleano temporal dentro del If( ) 
    cl.Cuadruplos[gotoFIndex].Arg2 = len(cl.Cuadruplos) + 1 // Completa el GOTOF con la posición actual (fin del bloque if)
}

func (cl *CuadruploList) BeginElse() int {
    // Agrega GOTO para el else
	jumpStack.Push(len(cl.Cuadruplos))
    cl.Cuadruplos = append(cl.Cuadruplos, Cuadruplo{
        Operator: "GOTO",
        Result:   -1,
    })
    return len(cl.Cuadruplos) - 1
}

func (cl *CuadruploList) CompleteElse() {
    // Completa el GOTO del else
	elsePos, _ := jumpStack.Pop()
    cl.Cuadruplos[elsePos].Result = len(cl.Cuadruplos)
}

// ------------ WHILE ----------------
func (cl *CuadruploList) BeginWhile()  {
	whileIndex = len(cl.Cuadruplos)

    // Guarda posición de inicio del while
    jumpStack.Push(len(cl.Cuadruplos))

	condition := BooleanTempVariables.Pop()
	cuadruplo := Cuadruplo{
		Operator: GOTOF,
		Arg1:     condition, // dirección del booleano temporal dentro del If( )
		Arg2:     0, // Se completará después
		Result:   0, 
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	jumpStack.Push(len(cl.Cuadruplos) - 1) // Guarda la posición del GOTOF
}

func (cl *CuadruploList) CompleteWhile() {
	
	gotoFIndex, _ := jumpStack.Pop()
	cl.Cuadruplos[gotoFIndex].Arg1 = BooleanTempVariables.Pop() // direccion del booleano temporal dentro del If( )
	cl.Cuadruplos[gotoFIndex].Arg2 = len(cl.Cuadruplos) + 1 // Completa el GOTOF con la posición actual (fin del bloque if)

    // Agrega GOTOT para volver al inicio
    cl.Cuadruplos = append(cl.Cuadruplos, Cuadruplo{
        Operator: "GOTO",
        Arg2: whileIndex,
    })
}