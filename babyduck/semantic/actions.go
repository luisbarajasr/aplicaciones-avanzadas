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
func (functionDir *FunctionDirectory) RegisterFunction(name interface{}, ret Type) (interface{}, error) {
	functionName := string(name.(*token.Token).Lit)
	err := functionDir.AddFunction(functionName, ret)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// RegisterParam agrega un nuevo parámetro a la función actual
func (functionDir *FunctionDirectory) RegisterParam(name interface{}, typ Type) (interface{}, error) {
	if functionDir.CurrentFunction == nil {
		return nil, fmt.Errorf("no active function to add parameter %s", string(name.(*token.Token).Lit))
	}
	return nil, functionDir.AddParam(functionDir.CurrentFunction.Name, string(name.(*token.Token).Lit), typ)
}

// ---------------- QUADRUPLE SECTION ----------------
func (cuadruplo *CuadruploList) AddOperatorAction(op Operator) (interface{}, error) {
	// Add the operator to the stack
	cuadruplo.AddOperator(op)
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

	cuadruplo.addVariable(variable)
	return nil, nil
}

func (cuadruplo *CuadruploList) AddBeginIfAction() (interface{}, error) {
	cuadruplo.BeginIf()
	return nil, nil
}
func (cuadruplo *CuadruploList) CompleteIfAction() (interface{}, error) {
	cuadruplo.CompleteIf()
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
