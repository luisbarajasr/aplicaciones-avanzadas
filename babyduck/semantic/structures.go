package semantic


type VarStack struct {
    data []int // direccion virtual de la variable
}

type OpStack struct {
    data []Operator
}

func NewVarStack() *VarStack {
	return &VarStack{
		data: []int{},
	}
}

func (stack *VarStack) Push(val int) {
	stack.data = append(stack.data, val)
}

func (stack *VarStack) Pop() int {

	if stack.IsEmpty() {
		return 0
	}

	var ultimo int = stack.Peek()
	stack.data = stack.Reduce()

	return ultimo
}

func (stack *VarStack) Print() {
	for i := len(stack.data) - 1; i >= 0; i-- {
		println(stack.data[i])
	}
}

// Peek
func (stack *VarStack) Peek() int {
	return stack.data[len(stack.data)-1]
}

func (stack *VarStack) PeekDouble() (int, bool) {
	if len(stack.data) < 2 {
        return 0, false
    }
	return stack.data[len(stack.data)-2], true
}

// IsEmpty
func (stack *VarStack) IsEmpty() bool {
	return len(stack.data) == 0
}

// reduce
func (stack *VarStack) Reduce() []int {
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