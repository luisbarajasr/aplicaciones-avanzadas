package semantic

import (
	"fmt"
	"babyduck/token"
)

// AppendVariable agrega una nueva variable a la lista de variables temporales
func (functionDir *FunctionDirectory) AppendVariable(name interface{}) (interface{}, error) {
	varName := string(name.(*token.Token).Lit)
	functionDir.AppendTempVariable(varName)
	return nil, nil
}

// SaveVariables guarda la lista de variables temporales en la tabla de variables de la function actual o global
func (functionDir *FunctionDirectory) SaveVariables(varType Type) (interface{}, error) {
	functionDir.SaveTempVarList(varType)
	return nil, nil
}

// RegisterFunction registra una nueva función en el directorio de funciones
func (cl *CuadruploList) RegisterFunction(name interface{}, ret Type) (interface{}, error) {
	functionName := string(name.(*token.Token).Lit)
	err := cl.AddFunction(functionName, ret)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// RegisterParam agrega un nuevo parámetro a la función actual
func (functionDir *FunctionDirectory) RegisterParam(name interface{}, typ Type) (interface{}, error) {
	varName := string(name.(*token.Token).Lit)
	if functionDir.CurrentFunction == nil {
		return nil, fmt.Errorf("no active function to add parameter %s", string(name.(*token.Token).Lit))
	}
	return nil, functionDir.AddParam(functionDir.CurrentFunction.Name, varName, typ)
}

// PrintCurrentFunctionParams prints the parameters of the current function
func (functionDir *FunctionDirectory) PrintCurrentFunctionParams() (interface{}, error) {
	for _, param := range functionDir.CurrentFunction.Params {
		fmt.Printf("Param: %s, Type: %s, Virtual Address: %d\n", param.Name, param.Type, param.Virtual_address)
	}

	return nil, nil
}

// ---------------- QUADRUPLE SECTION ----------------
func (cuadruplo *CuadruploList) AddOperatorAction(op Operator) (interface{}, error) {
	// Add the operator to the stack
	if _, err := cuadruplo.AddOperator(op); err != nil {
		return nil, err
	}

	return nil, nil
}

func (cuadruplo *CuadruploList) AddVariableAction(name interface{}) (interface{}, error) {
	// Add the variable to the stack
	varName := string(name.(*token.Token).Lit)

	// search the variable in the function directory by using its name
	variable, err := cuadruplo.functionDir.LookupVariable(varName)
	if err != nil { 
		return nil, fmt.Errorf("error: variable %s not found", varName, err)
	}
// Adding variable: {z 0 1003}
	cuadruplo.addVariable(variable.Virtual_address)
	return nil, nil
}

func (cuadruplo *CuadruploList) AddBeginIfAction() (interface{}, error) {
	cuadruplo.AddOperator(Semicolon)
	cuadruplo.BeginIf()
	return nil, nil
}
func (cuadruplo *CuadruploList) CompleteIfAction() (interface{}, error) {
	cuadruplo.CompleteIf()
	return nil, nil
}
func (cuadruplo *CuadruploList) CompleteIfEmptyAction() (interface{}, error) {
	cuadruplo.CompleteIfEmpty()
	return nil, nil
}

func (cuadruplo *CuadruploList) AddBeginElseAction() (interface{}, error) {
	cuadruplo.BeginElse()
	return nil, nil
}
func (cuadruplo *CuadruploList) CompleteElseAction() (interface{}, error) {
	cuadruplo.CompleteElse()
	return nil, nil
}

func (cuadruplo *CuadruploList) AddBeginWhileAction() (interface{}, error) {
	cuadruplo.AddOperator(Semicolon)
	cuadruplo.BeginWhile()
	return nil, nil
}
func (cuadruplo *CuadruploList) CompleteWhileAction() (interface{}, error) {
	cuadruplo.CompleteWhile()
	return nil, nil
}

func (cuadruplo *CuadruploList) PrintCuadruplosAction() (interface{}, error) {
	// Print the list of quadruples
	cuadruplo.PrintCuadruplos()
	return nil, nil
}


// ----- FUNCTION CALLS -----

/*
Cuando termina la declaracion de una funcion, se debe agregar un cuadruplo de tipo ENDFUNC, que indica el final de la funcion.
*/ 
func (cuadruplo *CuadruploList) EndFunctionAction() (interface{}, error) {
	cuadruplo.EndFunction()
	return nil, nil
}

func (cuadruplo *CuadruploList) AddEraAction(name interface{}) (interface{}, error) {
	functionName := string(name.(*token.Token).Lit)
	cuadruplo.AddFunctionCallERA(functionName)
	return nil, nil
}

func (cuadruplo *CuadruploList) AddGoSubAction() (interface{}, error) {
	cuadruplo.AddGoSub(cuadruplo.function_called)
	return nil, nil
}

func (cuadruplo *CuadruploList) AddGotoMainAction() (interface{}, error) {
	cuadruplo.AddGotoMain()
	return nil, nil
}

func (cuadruplo *CuadruploList) CompleteMainAction() (interface{}, error) {
	cuadruplo.CompleteMain()
	return nil, nil
}

func (cuadruplo *CuadruploList) AddParamAction() (interface{}, error) {
	cuadruplo.AddOperator(Comma)
	if err := cuadruplo.AddParameter(); err != nil {
		return nil, err
	}	
	return nil, nil
}

// ---------------- CONSTANTS SECTION ----------------

func (cl *CuadruploList) IntconstantAction(value interface {}) (interface {}, error) {
	cte, _ := value.(*token.Token).Int32Value()
	virtual_address := cl.functionDir.MemoryManager.Allocate(ConstantInt)
	cl.functionDir.MemoryManager.ConstInts = append(cl.functionDir.MemoryManager.ConstInts, int(cte))
	cl.addVariable(virtual_address)
	return nil, nil
}

func (cl *CuadruploList) FloatconstantAction(value interface {}) (interface {}, error) {
	cte, _ := value.(*token.Token).Float64Value()
	virtual_address := cl.functionDir.MemoryManager.Allocate(ConstantFloat)
	cl.functionDir.MemoryManager.ConstFloats = append(cl.functionDir.MemoryManager.ConstFloats, cte)
	cl.addVariable(virtual_address)
	return nil, nil
}

func (cl *CuadruploList) StringconstantAction(value *token.Token) (interface {}, error) {
	cte := value.CharLiteralValue()
	virtual_address := cl.functionDir.MemoryManager.Allocate(ConstantString)
	cl.functionDir.MemoryManager.ConstStrings = append(cl.functionDir.MemoryManager.ConstStrings, cte)
	cl.addVariable(virtual_address)
	return nil, nil
}

func (cl *CuadruploList) Done() (interface {}, error) {
	// Finalize the quadruple list
	// cl.PrintCuadruplos()
	cl.AddDone()
	fmt.Printf("=================== OUTPUT DE LA VM ==========================\n")
	err := cl.Finalize()
	if err != nil {
		return nil, err
	}
	return nil, nil
}