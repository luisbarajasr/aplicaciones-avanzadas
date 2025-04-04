# Go Data Structures Implementation

This repository contains implementations of basic data structures in Go, specifically focusing on Queue, Stack and Dictionary (using map) implementations.

## Project Structure

```
actividades/
├── clases/           # Package containing data structure implementations
│   ├── queue.go      # Queue implementation
│   ├── stack.go      # Stack implementation
│   └── dictionary.go # Dictionary implementation
├── main.go           # Main program to test the implementations
└── go.mod            # Go module definition
```

## Data Structures

### Queue
A First-In-First-Out (FIFO) data structure implemented in `queue.go`. It provides the following operations:
- `NewQueue()`: Creates a new queue
- `Enqueue(value)`: Adds an element to the end of the queue
- `Dequeue()`: Removes and returns the element from the front of the queue
- `Peek()`: Returns the element at the front without removing it
- `Len()`: Returns the number of elements in the queue

### Stack
A Last-In-First-Out (LIFO) data structure implemented in `stack.go`. It provides the following operations:
- `NewStack()`: Creates a new stack
- `Push(value)`: Adds an element to the top of the stack
- `Pop()`: Removes and returns the element from the top of the stack
- `Peek()`: Returns the element at the top without removing it
- `Len()`: Returns the number of elements in the stack

## How to Run

1. Make sure you have Go installed on your system (version 1.20 or higher recommended)

2. Navigate to the project directory:
   ```bash
   cd actividades
   ```

3. Run the main program:
   ```bash
   go run main.go
   ```

The program will demonstrate the usage of Queue, Stack, and Dictionary data structures with various test cases.

## Example Output

When you run the program, you'll see output similar to this:

```
Testing Queue:
Queue length after enqueuing 3 items: 3
Front of queue: first

Testing Dequeue:
Dequeued: first
New front of queue: second
Queue length: 2

Testing empty queue:
Queue length after dequeuing all items: 0
Trying to dequeue empty queue: <nil>
Peeking empty queue: <nil>

Testing mixed operations:
Enqueued numbers, length: 2
Dequeued: 1
Final queue length: 2
Final front of queue: 2

Testing Dictionary:
Dictionary length: 3
Value for key 'one': 1
Keys: [one two three]
Values: [1 2 3]
```

## Notes

- Queue and Stack implementations use interfaces (`interface{}`) to allow storing any type of data
- Dictionary implementation uses Go's built-in map type for efficient key-value storage
- The implementations are not thread-safe
- Memory management is handled automatically by Go's garbage collector 