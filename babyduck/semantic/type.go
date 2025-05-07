package semantic

// Variable represents a variable with ID, type, and value
type Variable struct {
    Id    string
    Type  string // "int" or "float"
    Value string // Placeholder, e.g., "assigned"
}

// Function represents a function with a name and variables
type Function struct {
    Name string
    Vars map[string]Variable
}

// ProgramFunctions stores all functions, including the program (global scope)
var ProgramFunctions map[string]Function = make(map[string]Function)

// GlobalProgramName stores the program name for global scope
var GlobalProgramName string = ""

// CurrentModule stores the current function or program being processed
var CurrentModule string = ""

// Queue for variable declarations
type Queue struct {
    items []string
}

// varsQueue stores variable IDs before type assignment
var varsQueue = Queue{}

// CurrentType stores the type for queued variables
var CurrentType string = ""

func (q *Queue) Enqueue(item string) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
    if len(q.items) == 0 {
        return ""
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}