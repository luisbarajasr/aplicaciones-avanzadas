package semantic

import (
    "babyduck/token"
    "fmt"
)

// CreateFunction adds a new function or program to ProgramFunctions
func CreateFunction(name interface{}, isProgram bool) (interface{}, error) {
    id := string(name.(*token.Token).Lit)
    if _, ok := ProgramFunctions[id]; ok {
        return nil, fmt.Errorf("function %s already exists at line %d, column %d",
            id, name.(*token.Token).Line, name.(*token.Token).Column)
    }
    newFunc := Function{
        Name: id,
        Vars: make(map[string]Variable),
    }
    if isProgram {
        GlobalProgramName = id
    }
    ProgramFunctions[id] = newFunc
    CurrentModule = id
    return nil, nil
}

// AddVarToQueue queues a variable ID for later type assignment
func AddVarToQueue(name interface{}) (interface{}, error) {
    token, ok := name.(*token.Token)
    if !ok {
        return nil, fmt.Errorf("expected token, got %T", name)
    }
    id := string(token.Lit)
    varsQueue.Enqueue(id)
    if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
        return nil, fmt.Errorf("variable %s already exists in module %s at line %d, column %d",
            id, CurrentModule, token.Line, token.Column)
    }
    return nil, nil
}

// SetCurrentType sets the type for queued variables
func SetCurrentType(varType string) (interface{}, error) {
    if varType != "int" && varType != "float" {
        return nil, fmt.Errorf("invalid type: %s", varType)
    }
    CurrentType = varType
    return nil, nil
}

// AddVarsToTable assigns types to queued variables
func AddVarsToTable(varType string) (interface{}, error) {
    if varsQueue.IsEmpty() {
        return nil, fmt.Errorf("no variables queued for type %s", varType)
    }
    for !varsQueue.IsEmpty() {
        id := varsQueue.Dequeue()
        if _, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
            return nil, fmt.Errorf("variable %s already exists in module %s", id, CurrentModule)
        }
        if CurrentModule != GlobalProgramName {
            if _, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
                return nil, fmt.Errorf("global variable %s already exists in program %s", id, GlobalProgramName)
            }
        }
        newVar := Variable{
            Id:   id,
            Type: varType,
        }
        ProgramFunctions[CurrentModule].Vars[id] = newVar
    }
    return nil, nil
}

// AssignVarValue assigns a value to a variable with type checking
func AssignVarValue(name interface{}, exprType interface{}) (interface{}, error) {
    token, ok := name.(*token.Token)
    if !ok {
        return nil, fmt.Errorf("expected token, got %T", name)
    }
    id := string(token.Lit)
    exprTypeStr, ok := exprType.(string)
    if !ok {
        return nil, fmt.Errorf("expected string expression type, got %T", exprType)
    }
    v, ok := lookupVariable(id)
    if !ok {
        return nil, fmt.Errorf("variable %s not defined at line %d, column %d",
            id, token.Line, token.Column)
    }
    _, err := CheckTypes(v.Type, exprTypeStr, "=")
    if err != nil {
        return nil, fmt.Errorf("type mismatch at line %d, column %d: cannot assign %s to %s variable %s",
            token.Line, token.Column, exprTypeStr, v.Type, id)
    }
    v.Value = "assigned"
    ProgramFunctions[CurrentModule].Vars[id] = v
    return nil, nil
}

// GetVarValue retrieves a variable's type for expression evaluation
func GetVarValue(name interface{}) (interface{}, error) {
    token, ok := name.(*token.Token)
    if !ok {
        return nil, fmt.Errorf("expected token, got %T", name)
    }
    id := string(token.Lit)
    v, ok := lookupVariable(id)
    if !ok {
        return nil, fmt.Errorf("variable %s not defined at line %d, column %d",
            id, token.Line, token.Column)
    }
    if v.Value != "assigned" {
        return nil, fmt.Errorf("variable %s not assigned at line %d, column %d",
            id, token.Line, token.Column)
    }
    return v.Type, nil
}

// BinaryExpression computes the type of a binary expression
func BinaryExpression(leftType, op, rightType interface{}) (interface{}, error) {
    leftTypeStr, ok := leftType.(string)
    if !ok {
        return nil, fmt.Errorf("expected string left type, got %T", leftType)
    }
    rightTypeStr, ok := rightType.(string)
    if !ok {
        return nil, fmt.Errorf("expected string right type, got %T", rightType)
    }
    opToken, ok := op.(*token.Token)
    if !ok {
        return nil, fmt.Errorf("expected token for operator, got %T", op)
    }
    opStr := string(opToken.Lit)
    resultType, err := CheckTypes(leftTypeStr, rightTypeStr, opStr)
    if err != nil {
        return nil, fmt.Errorf("%s at line %d, column %d", err.Error(), opToken.Line, opToken.Column)
    }
    return resultType, nil
}

// lookupVariable checks for a variable in the current or global scope
func lookupVariable(id string) (Variable, bool) {
    if v, ok := ProgramFunctions[CurrentModule].Vars[id]; ok {
        return v, true
    }
    if CurrentModule != GlobalProgramName {
        if v, ok := ProgramFunctions[GlobalProgramName].Vars[id]; ok {
            return v, true
        }
    }
    return Variable{}, false
}