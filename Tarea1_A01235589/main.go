package main

import (
	"fmt"
	"actividades/clases"
)

func main(){
	// Create a new queue
	fmt.Println("-----------Testing Queue-----------")
	q := clases.NewQueue()

	// Test enqueue operations
	fmt.Println("Testing Enqueue:")
	q.Enqueue("first")
	q.Enqueue("second") 
	q.Enqueue("third")
	fmt.Printf("Queue length after enqueuing 3 items: %d\n", q.Len())
	fmt.Printf("Front of queue: %v\n", q.Peek())

	// Test dequeue operations
	fmt.Println("\nTesting Dequeue:")
	fmt.Printf("Dequeued: %v\n", q.Dequeue())
	fmt.Printf("New front of queue: %v\n", q.Peek())
	fmt.Printf("Queue length: %d\n", q.Len())

	// Test empty queue behavior
	fmt.Println("\nTesting empty queue:")
	q.Dequeue()
	q.Dequeue()
	fmt.Printf("Queue length after dequeuing all items: %d\n", q.Len())
	fmt.Printf("Trying to dequeue empty queue: %v\n", q.Dequeue())
	fmt.Printf("Peeking empty queue: %v\n", q.Peek())

	// Test mixed operations
	fmt.Println("\nTesting mixed operations:")
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Printf("Enqueued numbers, length: %d\n", q.Len())
	fmt.Printf("Dequeued: %v\n", q.Dequeue())
	q.Enqueue("mixed")
	fmt.Printf("Final queue length: %d\n", q.Len())
	fmt.Printf("Final front of queue: %v\n", q.Peek())

	// -- Test stack operations
	fmt.Println("-----------Testing Stack-----------")
	s := clases.NewStack()
	s.Push("first")
	s.Push("second")
	s.Push("third")
	fmt.Printf("Stack length after pushing 3 items: %d\n", s.Len())
	fmt.Printf("Top of stack: %v\n", s.Peek())
	fmt.Printf("Popped from stack: %v\n", s.Pop())
	fmt.Printf("New top of stack: %v\n", s.Peek())
	fmt.Printf("Stack length after popping: %d\n", s.Len())
	fmt.Printf("Popped from stack: %v\n", s.Pop())
	fmt.Printf("New top of stack: %v\n", s.Peek())

	// ----------- Test dictionary ----------- 
	fmt.Println("-----------Testing Dictionary-----------")
	d := clases.NewDictionary()
	d.Add("name", "John")
	d.Add("age", "30")
	d.Add("city", "New York")
	fmt.Printf("Dictionary size: %d\n", d.Size())
	
	// Get values with proper error handling
	if name, exists := d.Get("name"); exists {
		fmt.Printf("Get 'name': %s\n", name)
	}
	if age, exists := d.Get("age"); exists {
		fmt.Printf("Get 'age': %s\n", age)
	}
	if country, exists := d.Get("country"); exists {
		fmt.Printf("Get 'country': %s\n", country)
	} else {
		fmt.Println("'country' key does not exist")
	}
	
	d.Update("age", "31")
	if age, exists := d.Get("age"); exists {
		fmt.Printf("Updated 'age': %s\n", age)
	}
	d.Remove("city")
	fmt.Printf("Dictionary size after removing 'city': %d\n", d.Size())
	fmt.Printf("Contains 'city': %t\n", d.Contains("city"))
	fmt.Printf("Contains 'name': %t\n", d.Contains("name"))
	fmt.Printf("All items in dictionary: %v\n", d.GetAll())
	fmt.Printf("Clearing dictionary...\n")
	d.Clear()
	fmt.Printf("Dictionary size after clearing: %d\n", d.Size())
	

}