package semantic

import (
	"fmt"
)

// NewFunctionDirectory crea un nuevo directorio de funciones global
func NewFunctionDirectory() *FunctionDirectory {
	globalVars := NewVariableTable() // se crea una tabla de variables globales
	MemoryManager:= NewMemoryManager()

	return &FunctionDirectory{
		Functions:    make(map[string]*Function),
		GlobalVars:   globalVars,
		CurrentScope: globalVars,
		CurrentFunction: nil,
		MemoryManager: MemoryManager,
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

/* 
SE AGREGARA AL BNF EN LA SIGUIENTE VERSION (creacion y eliminacion de virtual address temporales)
PushScope cambia el scope actual a la tabla de variables de la funcion actual
y la funcion actual a la funcion que se pasa como argumento
*/
func (functionDir *FunctionDirectory) PushScope(function *Function) {
	functionDir.CurrentScope = function.Vars // variables actuales de la funcion 
	functionDir.CurrentFunction = function // la funcion actual en la que se encuentra
}	

/* 
SE AGREGARA AL BNF EN LA SIGUIENTE VERSION (creacion y eliminacion de virtual address temporales)
PopScope re-setea el scope actual a la tabla de variables globales
y la funcion actual a nil porque salimos de la funcion
*/
func (functionDir *FunctionDirectory) PopScope() {
	functionDir.CurrentScope = functionDir.GlobalVars
	functionDir.CurrentFunction = nil
}

func (functionDir *FunctionDirectory) ReturnFunctionDirectory() (interface {}, error) {
	return functionDir, nil
}

func (functionDir *FunctionDirectory) PrintFunctions() {
	for name, function := range functionDir.Functions {
		fmt.Printf("Function: %s\n", name)
		fmt.Printf("Return Type: %s\n", function.ReturnType)
		fmt.Printf("Parameters: ")
		for _, param := range function.Params {
			fmt.Printf("%s (%s) ", param.Name, param.Type)
		}

		for _, variable := range function.Vars.Variables {
			fmt.Printf("Variable: %s (%s) ", variable.Name, variable.Type)
			fmt.Printf("Virtual Address: %d ", variable.Virtual_address)
		}
		fmt.Println()
	}
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

func (functionDir *FunctionDirectory) SaveTempVarList(typ Type) error {
	
	var memType MemoryType

	for _, variable := range functionDir.TempVarList {

		// primero: establecer si es global o local para asignar la direccion virtual
		if functionDir.CurrentFunction == nil {
			switch typ {
			case Int: 
				memType = GlobalInt
			case Float: 
				memType = GlobalFloat
			default:
                return fmt.Errorf("tipo %v no soportado para variables globales", typ)
			}
		} else {
			switch typ {
			case Int: 
				memType = LocalInt
			case Float:
				memType = LocalFloat
			default:
				return fmt.Errorf("tipo %v no soportado para variables locales", typ)
			}
		}

		// segundo: asignar la direccion virtual en su respsectiva seccion
		varDir := functionDir.MemoryManager.Allocate(memType)
		newVar := Variable{
			Name: variable.Name,
			Type: typ,
			Virtual_address: varDir,
		}

		// tercero: guardarlo en la tabla de variables en local o global
        if functionDir.CurrentFunction == nil {
            functionDir.GlobalVars.Variables[variable.Name] = newVar
        } else {
            functionDir.CurrentFunction.Vars.Variables[variable.Name] = newVar
        }
	}
	functionDir.TempVarList = []Variable{}
    return nil
}

// func (functionDir *FunctionDirectory) SaveTempVarList(typ Type) {
// 	// si la funcion actual es nula, significa que estamos en global
// 	if functionDir.CurrentFunction == nil {
// 		for _, variable := range functionDir.TempVarList {			
// 			varDir := functionDir.MemoryManager.Allocate(Global)
// 			functionDir.GlobalVars.Variables[variable.Name] = Variable{Name: variable.Name, Type: typ, Virtual_address: varDir}
// 		}
// 	} else if functionDir.CurrentFunction != nil {
// 		// si la funcion actual no es nula, significa que estamos en una funcion
// 		for _, variable := range functionDir.TempVarList {
// 			varDir := functionDir.MemoryManager.Allocate(Local)
// 			functionDir.CurrentFunction.Vars.Variables[variable.Name] = Variable{Name: variable.Name, Type: typ, Virtual_address: varDir}
// 		}
// 	}
// 	functionDir.TempVarList = []Variable{}
// 	return
// }

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

func (functionDir *FunctionDirectory) GetTypeByAddress(address int) (Type, error) {
	if address >= 1000 && address < 2000 || address >= 3000 && address < 4000 || address >= 6000 && address < 7000 || address >= 8000 && address < 9000{
		return Int, nil
	} else if address >= 2000 && address < 3000 || address >= 4000 && address < 5000 || address >= 7000 && address < 8000 || address >= 9000 && address < 10000{
		return Float, nil
	} else if address >= 5000 && address < 6000 {
		return Bool, nil
	} else if address >= 10000 && address < 11000 {
		return String, nil
	}

	return Int, nil // por defecto, no debe pasar
}