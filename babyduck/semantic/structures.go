package semantic


type VarStack struct {
    data []Variable
}

type OpStack struct {
    data []Operator
}

func NewVarStack() *VarStack {
	return &VarStack{
		data: []Variable{},
	}
}

func (stack *VarStack) Push(val Variable) {
	stack.data = append(stack.data, val)
}

func (stack *VarStack) Pop() Variable {

	if stack.IsEmpty() {
		return Variable{}
	}

	var ultimo Variable = stack.Peek()
	stack.data = stack.Reduce()

	return ultimo
}

func (stack *VarStack) Print() {
	for i := len(stack.data) - 1; i >= 0; i-- {
		println(stack.data[i].Name)
	}
}

// Peek
func (stack *VarStack) Peek() Variable {
	return stack.data[len(stack.data)-1]
}

func (stack *VarStack) PeekDouble() (Variable, bool) {
	if len(stack.data) < 2 {
        return Variable{}, false
    }
	return stack.data[len(stack.data)-2], true
}

// IsEmpty
func (stack *VarStack) IsEmpty() bool {
	return len(stack.data) == 0
}

// reduce
func (stack *VarStack) Reduce() []Variable {
	return stack.data[:len(stack.data)-1]
}

// ----------------------------------------------------

func NewOpStack() *OpStack {
	return &OpStack{
		data: []Operator{},
	}
}


func (stack *OpStack) Push(operador Operator) {
	stack.data = append(stack.data, operador)
}

func (stack *OpStack) Pop() (Operator, bool) {

	if stack.IsEmpty() {
		return Operator(""), false
	}

	var ultimo Operator = stack.Peek()
	stack.data = stack.Reduce()

	return ultimo, true
}

// Peek
func (stack *OpStack) Peek() Operator {
	return stack.data[len(stack.data)-1]
}

// IsEmpty
func (stack *OpStack) IsEmpty() bool {
	return len(stack.data) == 0
}

// reduce
func (stack *OpStack) Reduce() []Operator {
	return stack.data[:len(stack.data)-1]
}

func (stack *OpStack) Print() {
	for i := len(stack.data) - 1; i >= 0; i-- {
		println(stack.data[i])
	}
}