// semantic/vm.go

/* 
1. replicar functionDir y MemoryManager de semantic
2. usar CuadruploList.Cuadruplos, functionDir y MemoryManager para ejecutar (simplemente pasando CuadruploList ya que contiene todo )
3. ejectuar cada cuadruplo en el orden que se encuentran
*/ 
package semantic

import (
	"fmt"
	"strings"
	

)

type Memory struct {
	Ints    map[int]int
	Floats  map[int]float64
	Bools   map[int]bool
}

type VirtualMachine struct {
	Globals    Memory
	Locals     []Memory // stack of local memory frames
	Temps      Memory
	Constants  Memory
	IP         int // instruction pointer
	Stack      []int // call stack
	Cuadruplos []Cuadruplo
}

func NewMemory() Memory {
	return Memory{
		Ints:   make(map[int]int),
		Floats: make(map[int]float64),
		Bools:  make(map[int]bool),
	}
}

func NewVM(cl []Cuadruplo) *VirtualMachine {
	vm := &VirtualMachine{
		Globals:   NewMemory(),
		Locals:    []Memory{},
		Temps:     NewMemory(),
		Constants: NewMemory(),
		IP:        0,
		Stack:     []int{},
		Cuadruplos: cl,
	}

	return vm
}

func (vm *VirtualMachine) getCurrentLocalMemory() Memory {
	fmt.Printf("getCurrentLocalMemory called, Locals length: %d\n", len(vm.Locals))
	
	return vm.Locals[len(vm.Locals)-1]
}

func (vm *VirtualMachine) getValue(address int) interface{} {
	// fmt.Printf("Accessing variable at address %d\n", address)

	switch {
	case address >= 1000 && address < 2000:
		fmt.Printf("Value: %d\n", vm.Globals.Ints[address])
		return vm.Globals.Ints[address]
	case address >= 2000 && address < 3000:
		fmt.Printf("Value: %f\n", vm.Globals.Floats[address])
		return vm.Globals.Floats[address]
	case address >= 3000 && address < 4000:
		// fmt.Printf("Accessing local memory int at address %d\n", address)
		return vm.getCurrentLocalMemory().Ints[address]
	case address >= 4000 && address < 5000:
		// fmt.Printf("Accessing local memory float at address %d\n", address)
		return vm.getCurrentLocalMemory().Floats[address]
	case address >= 5000 && address < 6000:
		return vm.Temps.Bools[address]
	case address >= 6000 && address < 7000:
		// fmt.Printf("Accessing temporary int at address %d\n", address)
		return vm.Temps.Ints[address]
	case address >= 7000 && address < 8000:
		// fmt.Printf("Accessing temporary float at address %d\n", address)
		return vm.Temps.Floats[address]
	case address >= 8000:
		fmt.Printf("Accessing constant at address %d\n", vm.Constants.Ints[address])
		if val, ok := vm.Constants.Ints[address]; ok {
			fmt.Printf("VALOR DE CONSTANTE: %d\n", val)
			return val
		} else if valf, ok := vm.Constants.Floats[address]; ok {
			fmt.Printf("VALOR DE CONSTANTE FLOAT: %f\n", valf)
			return valf
		} else if valb, ok := vm.Constants.Bools[address]; ok {
			return valb
		}
	}
	return nil
}

func (vm *VirtualMachine) setValue(address int, value interface{}) {
	switch v := value.(type) {
	case int:
		if address >= 1000 && address < 2000 {
			vm.Globals.Ints[address] = v
		} else if address >= 3000 && address < 4000 {
			mem := vm.getCurrentLocalMemory()
			mem.Ints[address] = v
		} else if address >= 6000 && address < 7000 {
			vm.Temps.Ints[address] = v
		} else if address >= 8000 {
			vm.Constants.Ints[address] = v
		}
	case float64:
		if address >= 2000 && address < 3000 {
			vm.Globals.Floats[address] = v
		} else if address >= 4000 && address < 6000 {
			mem := vm.getCurrentLocalMemory()
			mem.Floats[address] = v
		} else if address >= 7000 && address < 8000 {
			vm.Temps.Floats[address] = v
		} else if address >= 8000 {
			vm.Constants.Floats[address] = v
		}
	case bool:
		if address >= 5000 && address < 6000 {
			vm.Temps.Bools[address] = v
		}
	default:
		fmt.Printf("Unknown type for value: %v\n", value)
	}
}

func Run(cl *CuadruploList) (string, error) {
	var output strings.Builder
	vm := NewVM(cl.Cuadruplos)
	vm.ReplicateMM(cl.functionDir)

	// fmt.Printf("Starting execution with Cuadruplos: %v\n", vm.Cuadruplos)

	for vm.IP < len(vm.Cuadruplos) {

		fmt.Printf("Executing Cuadruplo at IP %d: %v\n", vm.IP, vm.Cuadruplos[vm.IP])

		quad := vm.Cuadruplos[vm.IP]
		switch quad.Operator {
		// case Plus:
		// 	v1 := vm.getValue(quad.Arg1).(int)
		// 	v2 := vm.getValue(quad.Arg2).(int)
		// 	vm.setValue(quad.Result, v1+v2)
		// case Minus:
		// 	v1 := vm.getValue(quad.Arg1).(int)
		// 	v2 := vm.getValue(quad.Arg2).(int)
		// 	vm.setValue(quad.Result, v1-v2)
		// case Times:
		// 	v1 := vm.getValue(quad.Arg1).(int)
		// 	v2 := vm.getValue(quad.Arg2).(int)
		// 	vm.setValue(quad.Result, v1*v2)
		// case Divide:
		// 	v1 := vm.getValue(quad.Arg1).(int)
		// 	v2 := vm.getValue(quad.Arg2).(int)
		// 	if v2 == 0 {
		// 		panic("Runtime error: division by zero")
		// 	}
		// 	vm.setValue(quad.Result, v1/v2)
		case Plus, Minus, Times, Divide:
			val1 := vm.getValue(quad.Arg1)
			val2 := vm.getValue(quad.Arg2)

			// Handle both int and float operations
			switch v1 := val1.(type) {
			case int:
				if v2, ok := val2.(int); ok {
					var result int
					switch quad.Operator {
					case Plus:
						result = v1 + v2
					case Minus:
						result = v1 - v2
					case Times:
						result = v1 * v2
					case Divide:
						if v2 == 0 {
							return "", fmt.Errorf("division by zero")
						}
						result = v1 / v2
					}
					vm.setValue(quad.Result, result)
				} else {
					return "", fmt.Errorf("type mismatch: both operands must be int")
				}
			case float64:
				if v2, ok := val2.(float64); ok {
					var result float64
					switch quad.Operator {
					case Plus:
						result = v1 + v2
					case Minus:
						result = v1 - v2
					case Times:
						result = v1 * v2
					case Divide:
						if v2 == 0 {
							return "", fmt.Errorf("division by zero")
						}
						result = v1 / v2
					}
					vm.setValue(quad.Result, result)
				} else {
					return "", fmt.Errorf("type mismatch: both operands must be float")
				}
			default:
				return "", fmt.Errorf("unsupported operand type: %T", v1)
			}
        case Assign:
            val := vm.getValue(quad.Arg1)
            vm.setValue(quad.Result, val)
        
		case GOTO:
			vm.IP = quad.Arg2
			continue
		case GOTOF:
			cond := vm.getValue(quad.Arg1).(bool)
			if !cond {
				vm.IP = quad.Result
				continue
			}
		case ERA:
			// Create a new local memory frame for the function
			fmt.Printf("Creating new local memory frame for function at IP %d\n", vm.IP)			
			vm.Locals = append(vm.Locals, NewMemory())
			fmt.Printf("New local memory frame created: %v\n", vm.Locals[len(vm.Locals)-1])

		case GOSUB:
			// Save the current instruction pointer to the call stack
			vm.Stack = append(vm.Stack, vm.IP+1)
			// Jump to the function address
			vm.IP = quad.Result
			continue

        case PARAM:
			fmt.Printf("Setting parameter at address %d\n", quad.Result)
            val := vm.getValue(quad.Result)
            currentLocal := vm.getCurrentLocalMemory()
            
            // Determine where to store based on parameter index
            switch v := val.(type) {
            case int:
                currentLocal.Ints[quad.Result] = v
            case float64:
                currentLocal.Floats[quad.Result] = v
            case bool:
                currentLocal.Bools[quad.Result] = v
            }
			fmt.Printf("SALIENDO DE PARAM: %v\n", val)

		case ENDFUNC:
			fmt.Printf("TERMINANDO FUNCION EN IP %d\n", vm.IP)
			fmt.Printf("Pila de llamadas: %v\n", vm.Stack)

            // Pop return address
            if len(vm.Stack) > 0 {
                vm.IP = vm.Stack[len(vm.Stack)-1]-1
                vm.Stack = vm.Stack[:len(vm.Stack)-1]
            }

            // Pop local memory
            if len(vm.Locals) > 0 {
                vm.Locals = vm.Locals[:len(vm.Locals)-1]
            }
		
		case Print:
            val := vm.getValue(quad.Result)
            output.WriteString(fmt.Sprintf("%v\n", val))
            vm.IP++
		
		case DONE:
			fmt.Println("Execution completed.")
            return output.String(), nil

		// Add more cases like PRINT, READ, ERA, PARAM, GOSUB, RETURN, etc.
		default:
            return "", fmt.Errorf("unknown operator %v", quad.Operator)
		}
		vm.IP++
	}
    return output.String(), nil
}

func (vm *VirtualMachine) ReplicateMM(fd *FunctionDirectory) {

	for _, variable := range fd.GlobalVars.Variables { 
		if variable.Type == Int {
			vm.Globals.Ints[variable.Virtual_address] = 0 // Initialize global int
		} else if variable.Type == Float {
			vm.Globals.Floats[variable.Virtual_address] = 0.0 // Initialize global float
		}
		
	}

    // Transfer constant integers
    baseInt := 8000
    for _, val := range fd.MemoryManager.ConstInts {
        vm.Constants.Ints[baseInt] = val
    }
    
    // Transfer constant floats
    baseFloat := 9000
    for _, val := range fd.MemoryManager.ConstFloats {
        vm.Constants.Floats[baseFloat] = val
    }

	fmt.Printf("Replicated memory manager: %v\n", vm.Globals)
	fmt.Printf("Replicated constants: %v\n", vm.Constants)

}