package clases

// definir el tipo de dato queue, es como una clase
type (
	Queue struct {
		start, end *queueNode // apunta al primer y ultimo nodo
		length int
	}

	queueNode struct{
		value interface{} // recibe cualquier tipo de dato
		next *queueNode // apunta al siguiente nodo
	}
)

//crear nuevo queue
func NewQueue() *Queue{
	return &Queue{nil, nil, 0}
}

//tomar el next item fuera de frente del queue
func (this *Queue) Dequeue() interface{}{
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else{
		this.start = this.start.next
	}
	this.length--
	return n.value
}

// poner un item al final del queue
func (this *Queue) Enqueue(value interface{}){
	n := &queueNode{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// regresar el numero de elementos en el queue
func (this *Queue) Len() int{
	return this.length
}

// regresar el valor en el frente del queue sin quitarlo
func (this *Queue) Peek() interface{}{
	if this.length == 0 {
		return nil
	}
	return this.start.value
}

// func main(){
// 	// Create a new queue
// 	q := NewQueue()

// 	// Test enqueue operations
// 	fmt.Println("Testing Enqueue:")
// 	q.Enqueue("first")
// 	q.Enqueue("second") 
// 	q.Enqueue("third")
// 	fmt.Printf("Queue length after enqueuing 3 items: %d\n", q.Len())
// 	fmt.Printf("Front of queue: %v\n", q.Peek())

// 	// Test dequeue operations
// 	fmt.Println("\nTesting Dequeue:")
// 	fmt.Printf("Dequeued: %v\n", q.Dequeue())
// 	fmt.Printf("New front of queue: %v\n", q.Peek())
// 	fmt.Printf("Queue length: %d\n", q.Len())

// 	// Test empty queue behavior
// 	fmt.Println("\nTesting empty queue:")
// 	q.Dequeue()
// 	q.Dequeue()
// 	fmt.Printf("Queue length after dequeuing all items: %d\n", q.Len())
// 	fmt.Printf("Trying to dequeue empty queue: %v\n", q.Dequeue())
// 	fmt.Printf("Peeking empty queue: %v\n", q.Peek())

// 	// Test mixed operations
// 	fmt.Println("\nTesting mixed operations:")
// 	q.Enqueue(1)
// 	q.Enqueue(2)
// 	fmt.Printf("Enqueued numbers, length: %d\n", q.Len())
// 	fmt.Printf("Dequeued: %v\n", q.Dequeue())
// 	q.Enqueue("mixed")
// 	fmt.Printf("Final queue length: %d\n", q.Len())
// 	fmt.Printf("Final front of queue: %v\n", q.Peek())
// }