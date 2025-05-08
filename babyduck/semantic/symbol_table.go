package semantic

import (
	"fmt"
)

// NewFunctionDirectory crea un nuevo directorio de funciones global
func NewFunctionDirectory() *FunctionDirectory {
	globalVars := NewVariableTable() // se crea una tabla de variables globales
	return &FunctionDirectory{
		Functions:    make(map[string]*Function),
		GlobalVars:   globalVars,
		CurrentScope: globalVars,
		CurrentFunction: nil,
	}
}

// AddFunction agrega una nueva función al directorio de funciones global
func (functionDir *FunctionDirectory) AddFunction(name string, returnType Type) error {
	if value, exists := functionDir.Functions[name]; exists {
		return fmt.Errorf("error: function %s already exists with return type %s", name, value.ReturnType)
	}
	functionDir.Functions[name] = &Function{
		Name:       name,
		ReturnType: returnType,
		Params:     []Variable{},
		ParamsType: []Type{},
		Vars:       NewVariableTable(),
	}
	functionDir.PushScope(functionDir.Functions[name])
	return nil
}

// AddParam agrega un nuevo parámetro a la función
func (functionDir *FunctionDirectory) AddParam(funcName, paramName string, tipo Type) error {
	if function, exists := functionDir.Functions[funcName]; exists {
		function.Params = append(function.Params, Variable{Name: paramName, Type: tipo})
		function.ParamsType = append(function.ParamsType, tipo)
		return function.Vars.AddVariable(paramName, tipo)
	}
	return fmt.Errorf("function %s not found", funcName)
}

// PushScope sets the current scope to a new variable table
func (functionDir *FunctionDirectory) PushScope(function *Function) {
	functionDir.CurrentScope = function.Vars // variables actuales de la funcion 
	functionDir.CurrentFunction = function // la funcion actual en la que se encuentra
}	

// PopScope resets the current scope to the global scope
func (functionDir *FunctionDirectory) PopScope() {
	functionDir.CurrentScope = functionDir.GlobalVars
	functionDir.CurrentFunction = nil
}

func (functionDir *FunctionDirectory) ReturnFunctionDirectory() (interface {}, error) {
	return functionDir, nil
}

/* --- SECCION DE VARIABLES --- */

func NewVariableTable() *VariableTable {
	return &VariableTable{Variables: make(map[string]Variable)}
}

// AddVariable adds a variable to the table
func (varTable *VariableTable) AddVariable(name string, varType Type) error {
	if value, exists := varTable.Variables[name]; exists {
		return fmt.Errorf("error: variable %s already exists with type %s", name, value.Type)
	}
	varTable.Variables[name] = Variable{Name: name, Type: varType}
	return nil
}

// LookupVariable revisa si la variable existe en la tabla de variables locales o globales
func (functionDir *FunctionDirectory) LookupVariable(name string) (Variable, error) {
    // revisar en la tabla de variables locales
    if value, exists := functionDir.CurrentScope.Variables[name]; exists {
        return value, nil
    }
    return Variable{}, fmt.Errorf("error: variable %s does not exist", name)
}