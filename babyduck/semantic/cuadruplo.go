package semantic
// cuadruplo 
type Cuadruplo struct {
	Arg1     Variable
	Arg2     *Variable // siendo un puntero, permite ser nil, en caso de que no se necesite, eg. asignar valores
	Operator Operator
	Result   Variable
}

type CuadruploList struct {
	Cuadruplos []Cuadruplo
}

var TempVariables = map[string]Variable

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

func (op *Operator) IsRightAssociative() bool {
    return op == Operator.Assign
}

func NewTempVariable(name string, typ Type) Variable {
    return Variable{
        Name:          name,
        Type:        typ,
    }
}

// ------------ CUADRUPLO LIST ------------ 


func NewCuadruploList() *CuadruploList {
	var (
		OpStack  *OpStack  // Global OpStack instance
		VarStack *VarStack // Global VarStack instance
	)
	
	OpStack = NewOpStack()   // Initialize OpStack
    VarStack = NewVarStack() // Initialize VarStack

    return &CuadruploList{
        Cuadruplos: []Cuadruplo{},
    }
}

func (CuadruploList *CuadruploList) AddOperator(operator Operator) (Operator, error){

	if OpStack.IsEmpty() {
		// agregar operador a la pila
		OpStack.Push(operator)
		return operator, nil

	} else if(operator == NewPara){
		OpStack.Push(operator)
		return operator, nil

	} else if (operator == ClosePara) {

		for OpStack.Peek() != NewPara {
			topOperator := OpStack.Pop()
			CuadruploList.addCuadruplo(topOperator)
			OpStack.Pop()
		}

		OpStack.Pop() // sacar el falso stack
		return OpStack.Peek(), nil

	}else if (operator.IsRightAssociative() && operator.Precedence() > OpStack[len(OpStack)-1].Precedence()) || (!operator.IsRightAssociative() && operator.Precedence() >= OpStack[len(OpStack)-1].Precedence()) {
		OpStack.Push(operator)
		return operator, nil

	} else if operator.Precedence() < OpStack[len(OpStack)-1].Precedence() {
		// sacar el operador de la cima de la pila
		topOperator := OpStack.Pop()

		// agergar el operador a la lista de cuadruplos
		CuadruploList.addCuadruplo(topOperator)

		// agregar el nuevo operador a la pila
		OpStack.Push(operator)
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
	VarStack.VarPush(variable)
}

func (CuadruploList *CuadruploList) addCuadruplo(operator Operator) {
	var1, var2 := VarStack[len(VarStack)-2], VarStack[len(VarStack.Stack)-1]

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

	CuadruploList = append(CuadruploList.Cuadruplos, cuadruplo)
	VarStack.VarPop() // quitar las 2 variables de la pila
	VarStack.VarPop() 
	VarStack.VarPush(cuadruplo.Result) // agregar el resultado al stack
	TempVariables[cuadruplo.Result.Name] = opResult
}

func (CuadruploList *CuadruploList) PrintCuadruplos() {
	for _, cuadruplo := range CuadruploList.Cuadruplos {
		println(cuadruplo.Arg1.Name, cuadruplo.Operator, cuadruplo.Arg2.Name, "->", cuadruplo.Result.Name)
	}
	println("Total Cuadruplos:", len(CuadruploList.Cuadruplos))
}