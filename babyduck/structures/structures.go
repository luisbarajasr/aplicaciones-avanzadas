package structures

import (
	"babyduck/semantic"
)

type VarStack struct {
	data []int
}

func (stack *VarStack) Push(val int) {
	stack.data = append(stack.data, val)
}

func (stack *VarStack) Pop() int {

	if stack.IsEmpty() {
		return -1
	}

	var ultimo int = stack.Peek()
	stack.data = stack.Reduce()

	return ultimo
}

// Peek
func (stack *VarStack) Peek() int {
	return stack.data[len(stack.data)-1]
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

type OpStack struct {
	data []semantic.Operator
}

func (stack *OpStack) Push(opertator semantic.Operator) {
	stack.data = append(stack.data, operador)
}

func (stack *OpStack) Pop() int {

	if stack.IsEmpty() {
		return -1
	}

	var ultimo semantic.Operator = stack.Peek()
	stack.data = stack.Reduce()

	return ultimo
}

// Peek
func (stack *OpStack) Peek() semantic.Operator {
	return stack.data[len(stack.data)-1]
}

// IsEmpty
func (stack *OpStack) IsEmpty() bool {
	return len(stack.data) == 0
}

// reduce
func (stack *OpStack) Reduce() []semantic.Operator {
	return stack.data[:len(stack.data)-1]
}