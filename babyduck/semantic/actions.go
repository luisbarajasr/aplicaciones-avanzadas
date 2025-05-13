package semantic

import (
	"fmt"
	"babyduck/token"
)

// ValidateAssign revisa si el tipo de la variable coincide con el tipo de la expresion
func (functionDir *FunctionDirectory) ValidateAssign(name interface{}) (interface{}, error) {
	
	varName := string(name.(*token.Token).Lit)
	_, err := functionDir.LookupVariable(varName)
	if err != nil {
		return nil, err
	}
	// resultType, err := CheckTypes(varDecl.Type, exprType, Assign)
	// if err != nil {
	// 	return fmt.Errorf("type mismatch in assignment: %s (%s) cannot be assigned %s", varName, varDecl.Type, exprType)
	// }
	// if resultType == Void {
	// 	return fmt.Errorf("type mismatch in assignment: %s (%s) cannot be assigned %s", varName, varDecl.Type, exprType)
	// }
	return nil, nil
}

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
