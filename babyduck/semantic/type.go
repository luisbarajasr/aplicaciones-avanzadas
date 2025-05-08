package semantic

type Type string

const (
	Int   Type = "int"
	Float Type = "float"
	Void  Type = "void"
	Bool  Type = "bool"
	Error Type = "error"
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
)

// Variable represents a variable with a name and type
type Variable struct {
	Name string
	Type Type
}

// VariableTable holds variables in a scope
type VariableTable struct {
	Variables map[string]Variable
}

// Function represents a function with its variables
type Function struct {
	Name       string
	ReturnType Type
	Params     []Variable
	ParamsType []Type
	Vars       *VariableTable
}

// FunctionDirectory holds all functions and global variables
type FunctionDirectory struct {
	Functions    map[string]*Function
	GlobalVars   *VariableTable
	CurrentScope *VariableTable
	CurrentFunction *Function
}