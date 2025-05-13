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
		TempVarList: []Variable{},
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
	// cambiamos el scope actual a la funcion que acabamos de crear
	functionDir.PushScope(functionDir.Functions[name])
	return nil
}

// AddParam agrega un nuevo parámetro a la función
func (functionDir *FunctionDirectory) AddParam(funcName, paramName string, tipo Type) error {
	if function, exists := functionDir.Functions[funcName]; exists {
		function.Params = append(function.Params, Variable{Name: paramName, Type: tipo})
		function.ParamsType = append(function.ParamsType, tipo)
		function.Vars.Variables[paramName] = Variable{Name: paramName, Type: tipo}
		return nil
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

func (functionDir *FunctionDirectory) AppendTempVariable(name string) error {

	// Verificar si la variable ya existe en el scope actual o global
	if _, err := functionDir.LookupVariable(name); err == nil {
		return fmt.Errorf("error: variable %s already exists", name)
	}

	//agregar la variable a la tabla de variables de la funcion actual
	functionDir.TempVarList = append(functionDir.TempVarList, Variable{Name: name})
	return nil
}

func (functionDir *FunctionDirectory) SaveTempVarList(typ Type) {
	// si la funcion actual es nula, significa que estamos en global
	if functionDir.CurrentFunction == nil {
		for _, variable := range functionDir.TempVarList {
			functionDir.GlobalVars.Variables[variable.Name] = Variable{Name: variable.Name, Type: typ}
		}
	} else if functionDir.CurrentFunction != nil {
		// si la funcion actual no es nula, significa que estamos en una funcion
		for _, variable := range functionDir.TempVarList {
			functionDir.CurrentFunction.Vars.Variables[variable.Name] = Variable{Name: variable.Name, Type: typ}
		}
	}
	functionDir.TempVarList = []Variable{}
	return
}

func (functionDir *FunctionDirectory) LookupVariable(name string) (Variable, error) {
	// primero revisamos si la variable existe en la tabla de variables de la funcion actual
	if variable, exists := functionDir.CurrentScope.Variables[name]; exists {
		return variable, nil
	}
	// si no existe, revisamos si existe en la tabla de variables globales
	if variable, exists := functionDir.GlobalVars.Variables[name]; exists {
		return variable, nil
	}
	return Variable{}, fmt.Errorf("error: variable %s not found", name)
}