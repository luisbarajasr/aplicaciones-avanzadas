package clases 
import (
	"fmt"
)
type (
	Stack struct {
		top *stackNode
		length int
	}

	stackNode struct{
		value interface{}
		prev *stackNode
	}
)

// crear nuevo stack
func NewStack() *Stack{
	return &Stack{nil, 0}
}

// regresar el numero de elementos en el stack
func (this *Stack) Len() int{
	return this.length
}

// regresar el valor en el tope del stack
func (this *Stack) Peek() interface{}{
	if this.length == 0 {
		return nil
	}	
	return this.top.value
}

// remover el valor en el tope del stack
func (this *Stack) Pop() interface{}{
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// agregar un valor al tope del stack
func (this *Stack) Push(value interface{}){
	n := &stackNode{value, this.top}
	this.top = n
	this.length++
}

func main(){
	q := NewStack()
	q.Push("first")
	q.Push("second")
	q.Push("third")
	fmt.Printf("Queue length after enqueuing 3 items: %d\n", q.Len())
	fmt.Printf("Front of queue: %v\n", q.Peek())
}