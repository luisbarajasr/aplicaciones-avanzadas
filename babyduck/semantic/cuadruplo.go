package semantic

import (
	"fmt"
)

/* 
TODO:
- Manejo de constantes
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
		function_called: "",
    }
}

func (CuadruploList *CuadruploList) addVariable(variable Variable){
	varStack.Push(variable.Virtual_address)
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

	// Si el operador es Comma, no se agrega al stack de operadores porque se usa para separar parámetros en funciones
	if opStack.IsEmpty() && operator != Comma { 
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

			return operator, nil

		case Comma:
			// Pop operators until parenthesis or op stack is empty
			for !opStack.IsEmpty() && opStack.Peek() != NewPara {
				topOp, _ := opStack.Pop()
				CuadruploList.addCuadruplo(topOp)
			}
			return operator, nil

		case Semicolon:
			for !opStack.IsEmpty() {
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

	switch operator {
		
		case Assign:
			result := varStack.Pop()
			varTarget := varStack.Pop()
			
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
		
		default:
		arg2 := varStack.Pop()
		arg1 := varStack.Pop()

		arg1Type,_:= CuadruploList.functionDir.GetTypeByAddress(arg1)
		arg2Type,_ := CuadruploList.functionDir.GetTypeByAddress(arg2)
		
		var opResult,_ = CheckTypes(arg1Type, arg2Type, operator)

		if opResult == Error {
			panic("Error: Type mismatch in operation")
		}

		resultVarType := varTypetoMemoryType(opResult) // convierte el tipo de variable a un tipo de memoria para la asignacion de dir correcta 
		resultAddress := CuadruploList.functionDir.MemoryManager.Allocate(resultVarType) // sera un temporal de int, float o bool
		
		if CuadruploList.functionDir.CurrentFunction != nil {
			CuadruploList.functionDir.CurrentFunction.vars_needed += 1 // incrementa el contador de variables necesarias para la funcion actual
		}
		
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
    // jumpStack.Push(len(cl.Cuadruplos))

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
        Operator: GOTO,
        Arg2: whileIndex,
    })
}

// ------------- GOTO MAIN ----------------
func (cl *CuadruploList) AddGotoMain()  {
    jumpStack.Push(len(cl.Cuadruplos))

	cuadruplo := Cuadruplo{
		Operator: GOTO,
		Arg1:     0,
		Arg2:     0, // Se completará después
		Result:   0, 
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	jumpStack.Push(len(cl.Cuadruplos) - 1) // Guarda la posición del GOTO
}

func (cl *CuadruploList) CompleteMain() {
	gotoIndex, _ := jumpStack.Pop()
	cl.Cuadruplos[gotoIndex].Arg2 = len(cl.Cuadruplos) + 1 // Completa el GOTOF con la posición actual (fin del bloque if)
}

// ------------ FUNCTION CALLS ----------------

func (cl *CuadruploList) EndFunction() {
	cuadruplo := Cuadruplo{
		Operator: ENDFUNC,
		Arg1:     0, 
		Arg2:     0,
		Result:   0, 
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	cl.functionDir.MemoryManager.ResetLocalVariables()
	/* 
	resetear el directorio de variables locales, algo asi: 
	memory_address: map[MemoryType]int{
			LocalInt: 3000,
			LocalFloat: 4000,
			TemporalBool: 5000,
			TemporalInt: 6000,
			TemporalFloat: 7000,
		}

	!!!! variables globales y constantes no se resetean !!!!!
	*/
}

func (cl *CuadruploList) AddFunctionCallERA(funcName string) error {
	varsNeeded := cl.functionDir.Functions[funcName].vars_needed
	cuadruplo := Cuadruplo{
		Operator: ERA,
		Arg1:     0,
		Arg2:     0,
		Result:   varsNeeded, 
		// !!! numero de variables necesarias para la funcion !!!
		// debe ser el nombre de la funcion pero no se puede usar como argumento string
	}
	cl.function_called = funcName
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	return nil
}

func (cl *CuadruploList) AddGoSub(funcName string) error {
	varsNeeded := cl.functionDir.Functions[funcName].vars_needed
	cuadruplo := Cuadruplo{
		Operator: GOSUB,
		Arg1:     0,
		Arg2:     0,
		Result:   varsNeeded, 
		// !!! numero de variables necesarias para la funcion !!!
		// debe ser el nombre de la funcion pero no se puede usar como argumento string
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	return nil
}

func (cl *CuadruploList) AddParameter() error {
	paramAddress := varStack.Peek()
	varStack.Pop() // Sacamos la dirección de la variable del stack

	cuadruplo := Cuadruplo{
		Operator: PARAM,
		Arg1:     0,
		Arg2:     0,
		Result:   paramAddress, 
	}
	cl.Cuadruplos = append(cl.Cuadruplos, cuadruplo)
	return nil
}