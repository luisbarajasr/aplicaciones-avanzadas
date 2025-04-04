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

### Dictionary 
A key-value data structure implemented using pre-existing 'map' variable. Provides the following methods:
- `NewDictionary()`: Creates a new dictionary
- `Add(key, value)`: Adds a key-value pair
- `Get(key)`: Returns the value associated to a key and a boolean if exits.
- `Remove(key)`: Removes a key-value pair given a key.
- `Update(key, value)`: update the value associated to a key, returns a boolean if exists.
- `Contains(key)`: returns a if a key exists.
- `Size(key, value)`: returns dictionary full size.
- `GetAll()`: returns whole dictionary.


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
-----------Testing Queue-----------
Testing Enqueue:
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
-----------Testing Stack-----------
Stack length after pushing 3 items: 3
Top of stack: third
Popped from stack: third
New top of stack: second
Stack length after popping: 2
Popped from stack: second
New top of stack: first
-----------Testing Dictionary-----------
Dictionary size: 3
Get 'name': John
Get 'age': 30
'country' key does not exist
Updated 'age': 31
Dictionary size after removing 'city': 2
Contains 'city': false
Contains 'name': true
All items in dictionary: map[age:31 name:John]
Clearing dictionary...
Dictionary size after clearing: 0
```