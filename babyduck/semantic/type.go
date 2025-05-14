package semantic

type Type int

const (
	Int   Type = iota
	Float 
	Void
	Bool
	String
	Error
)

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
)

// Variable nombre y tipo
type Variable struct {
	Name string
	Type Type
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
}

// FunctionDirectory guarda todas las funciones y variables globales
type FunctionDirectory struct {
	Functions    map[string]*Function
	GlobalVars   *VariableTable
	CurrentScope *VariableTable
	CurrentFunction *Function
	TempVarList []Variable
}