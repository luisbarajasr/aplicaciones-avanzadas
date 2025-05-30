package semantic

type Operator string

const (
	Plus        Operator = "+"
	Minus       Operator = "-"
	Times       Operator = "*"	
	Divide      Operator = "/"
	Less        Operator = "<"
	Greater     Operator = ">"
	NotEqual    Operator = "!="
	Assign	  	Operator = "="
	NewPara   	Operator = "("
	ClosePara 	Operator = ")"
	Semicolon   Operator = ";"
	Comma 	 	Operator = ","
	Print 		Operator = "print"
	GOTO 		Operator = "GOTO"
	GOTOF 		Operator = "GOTOF"
	GOTOV 		Operator = "GOTOV"
	ERA			Operator = "ERA"
	GOSUB 		Operator = "GOSUB"
	ENDFUNC 	Operator = "ENDFUNC"
	RETURN 		Operator = "RETURN"
	PARAM 		Operator = "PARAM"
	DONE 		Operator = "DONE"
)

type Type int

const (
	Int   Type = iota
	Float 
	Void
	Bool
	String
	Error
)

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
	function_called string
}

type JumpStack struct {
    stack []int
}

// Variable nombre y tipo
type Variable struct {
	Name string
	Type Type
	Virtual_address int
}

// VariableTable guarda las variables de una función
type VariableTable struct {
	Variables map[string]Variable
}

// Function representa una función con su nombre, tipo de retorno, parámetros y variables
type Function struct {
	Name       string
	ReturnType Type
	Params     []Variable
	ParamsType []Type
	Vars       *VariableTable
	vars_needed int // contador de variables locales que la funcion necesita para la memoria de ejecucion
	virtual_address int // direccion virtual de la funcion en la memoria de ejecucion
	cuad_init int // direccion virtual del primer cuadruplo de la funcion
}

// FunctionDirectory guarda todas las funciones y variables globales
type FunctionDirectory struct {
	Functions    map[string]*Function
	GlobalVars   *VariableTable
	CurrentScope *VariableTable
	CurrentFunction *Function
	MemoryManager *MemoryManager
	TempVarList []Variable // lista para variables temporales que se agregan desde el BNF (lista de parametros, variables, etc)
}