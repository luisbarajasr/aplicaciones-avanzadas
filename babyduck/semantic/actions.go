package semantic

import (
	"fmt"
	"babyduck/token"
)

// ValidateAssign checks if an assignment is type-compatible
func (functionDir *FunctionDirectory) ValidateAssign(name interface{}, exprType Type) error {
	varName := string(name.(*token.Token).Lit)
	varDecl, err := functionDir.LookupVariable(varName)
	if err != nil {
		return err
	}
	resultType, err := CheckTypes(varDecl.Type, exprType, Assign)
	if err != nil {
		return fmt.Errorf("type mismatch in assignment: %s (%s) cannot be assigned %s", varName, varDecl.Type, exprType)
	}
	if resultType == Void {
		return fmt.Errorf("type mismatch in assignment: %s (%s) cannot be assigned %s", varName, varDecl.Type, exprType)
	}
	return nil
}

// RegisterGlobalVars adds multiple global variables of the same type
func (functionDir *FunctionDirectory) RegisterGlobalVars(names []string, typ Type) (interface{}, error) {
	for _, name := range names {
		if err := functionDir.GlobalVars.AddVariable(name, typ); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// RegisterFunction registers a function in the directory
func (functionDir *FunctionDirectory) RegisterFunction(name interface{}, ret Type, params []Type) (interface{}, error) {
	functionName := string(name.(*token.Token).Lit)
	err := functionDir.AddFunction(functionName, ret)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// MakeVarList builds a slice with a single identifier
func (functionDir *FunctionDirectory) MakeVarList(name interface{}) (interface{}, error) {
	return []string{string(name.(*token.Token).Lit)}, nil
}

// ConcatVarList adds an identifier to the front of an existing list
func (functionDir *FunctionDirectory) ConcatVarList(head interface{}, tail []string) (interface{}, error) {
	return append([]string{string(head.(*token.Token).Lit)}, tail...), nil
}

// MakeParamList builds a slice of parameter types with a single element
func (functionDir *FunctionDirectory) MakeParamList(param Type) (interface{}, error) {
	return []Type{param}, nil
}

// ConcatParamList adds a parameter type to the front of an existing list
func (functionDir *FunctionDirectory) ConcatParamList(head Type, tail []Type) (interface{}, error) {
	return append([]Type{head}, tail...), nil
}

// RegisterParam adds a parameter to the active function
func (functionDir *FunctionDirectory) RegisterParam(name interface{}, typ Type) (interface{}, error) {
	if functionDir.CurrentFunction == nil {
		return nil, fmt.Errorf("no active function to add parameter %s", string(name.(*token.Token).Lit))
	}
	return nil, functionDir.AddParam(functionDir.CurrentFunction.Name, string(name.(*token.Token).Lit), typ)
}

// ResolveVarType looks up a variable's type in the current scope
func (functionDir *FunctionDirectory) ResolveVarType(name interface{}) (interface{}, error) {
	varName := string(name.(*token.Token).Lit)
	varDecl, err := functionDir.LookupVariable(varName)
	if err != nil {
		return nil, err
	}
	return varDecl.Type, nil
}

// BinaryExpression computes the type of a binary expression
func (functionDir *FunctionDirectory) BinaryExpression(left Type, op interface{}, right Type) (interface{}, error) {
	opStr := string(op.(*token.Token).Lit)
	return CheckTypes(left, right, Operator(opStr))
}

// ConcatOperator handles operator precedence for MoreExp and MoreTerm
func (functionDir *FunctionDirectory) ConcatOperator(term Type, op string, prev Type) (interface{}, error) {
	if prev == "" {
		return term, nil
	}
	return CheckTypes(prev, term, Operator(op))
}

func (functionDir *FunctionDirectory) DeleteFuncDirectory() (interface{}, error) {
	functionDir.Functions = make(map[string]*Function)
	functionDir.GlobalVars = NewVariableTable()
	functionDir.CurrentScope = functionDir.GlobalVars
	functionDir.CurrentFunction = nil
	return nil, nil
}