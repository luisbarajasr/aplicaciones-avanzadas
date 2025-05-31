// semantic/vm.go

/* 
1. replicar functionDir y MemoryManager de semantic
2. usar CuadruploList.Cuadruplos, functionDir y MemoryManager para ejecutar (simplemente pasando CuadruploList ya que contiene todo )
3. ejectuar cada cuadruplo en el orden que se encuentran
*/ 
package semantic

import (
	"fmt"

)

type Memory struct {
	Ints    map[int]int
	Floats  map[int]float64
	Bools   map[int]bool
    Strings map[int]string
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
		Strings: make(map[int]string), 
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
	return vm.Locals[len(vm.Locals)-1]
}

func (vm *VirtualMachine) getValue(address int) interface{} {
	// fmt.Printf("Accessing variable at address %d\n", address)

	switch {
	case address >= 1000 && address < 2000:
		return vm.Globals.Ints[address]
	case address >= 2000 && address < 3000:
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
		if val, ok := vm.Constants.Ints[address]; ok {
			return val
		} else if valf, ok := vm.Constants.Floats[address]; ok {
			return valf
		} else if valb, ok := vm.Constants.Bools[address]; ok {
			return valb
		}
	case address >= 10000:
        if val, ok := vm.Constants.Strings[address]; ok {
            return val
        }
	}
	return nil
}

func (vm *VirtualMachine) setValue(address int, value interface{}) {
	switch v := value.(type) {
	case int:
		if address >= 1000 && address < 2000 {
			vm.Globals.Ints[address] = v
			// fmt.Printf("ASIGNING GLOBAL INT: %d at address %d\n", v, address)
		} else if address >= 3000 && address < 4000 {
			mem := vm.getCurrentLocalMemory()
			mem.Ints[address] = v
			// fmt.Printf("ASIGNING LOCAL INT: %d at address %d\n", v, address)
		} else if address >= 6000 && address < 7000 {
			vm.Temps.Ints[address] = v
			// fmt.Printf("ASIGNING TEMP INT: %d at address %d\n", v, address)
		} else if address >= 8000 {
			vm.Constants.Ints[address] = v
			// fmt.Printf("ASIGNING CONSTANT INT: %d at address %d\n", v, address)
		}
	case float64:
		if address >= 2000 && address < 3000 {
			vm.Globals.Floats[address] = v
			// fmt.Printf("ASIGNING GLOBAL FLOAT: %f at address %d\n", v, address)
		} else if address >= 4000 && address < 6000 {
			mem := vm.getCurrentLocalMemory()
			mem.Floats[address] = v
			// fmt.Printf("ASIGNING LOCAL FLOAT: %f at address %d\n", v, address)
		} else if address >= 7000 && address < 8000 {
			vm.Temps.Floats[address] = v
			// fmt.Printf("ASIGNING TEMP FLOAT: %f at address %d\n", v, address)
		} else if address >= 8000 {
			vm.Constants.Floats[address] = v
			// fmt.Printf("ASIGNING CONSTANT FLOAT: %f at address %d\n", v, address)
		}
	case bool:
		if address >= 5000 && address < 6000 {
			vm.Temps.Bools[address] = v
		}
	default:
		fmt.Printf("Unknown type for value: %v\n", value)
	}
}

func Run(cl *CuadruploList) error {
	vm := NewVM(cl.Cuadruplos)
	vm.ReplicateMM(cl.functionDir)

	for vm.IP < len(vm.Cuadruplos) {
		quad := vm.Cuadruplos[vm.IP]
		switch quad.Operator {

		case Plus:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			if  var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1+v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1+v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1)+v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1+float64(v2))
			} else {
				return fmt.Errorf("type mismatch in addition: %v + %v", var1Type, var2Type)
			}

		case Minus:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1-v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1-v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1)-v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1-float64(v2))
			} else {
				return fmt.Errorf("type mismatch in addition: %v + %v", var1Type, var2Type)
			}

		case Times:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1*v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1*v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1)*v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1*float64(v2))
			} else {
				return fmt.Errorf("type mismatch in addition: %v + %v", var1Type, var2Type)
			}

		case Divide:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1/v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1/v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1)/v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1/float64(v2))
			} else {
				return fmt.Errorf("type mismatch in addition: %v + %v", var1Type, var2Type)
			}
			
		case Less:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			// fmt.Printf("COMPARING LESS THAN: %v < %v\n", vm.getValue(quad.Arg1),vm.getValue(quad.Arg2))
			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 < v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1 < v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1) < v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 < float64(v2))
			} else {
				return fmt.Errorf("type mismatch in less than comparison: %v < %v", var1Type, var2Type)
			}
		case Greater:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)

			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 > v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1 > v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1) > v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 > float64(v2))
			} else {
				return fmt.Errorf("type mismatch in greater than comparison: %v > %v", var1Type, var2Type)
			}
		case NotEqual:
			var1Type := cl.addressToType(quad.Arg1)
			var2Type := cl.addressToType(quad.Arg2)
			if var1Type == Int && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 != v2)
			} else if var1Type == Float && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, v1 != v2)
			} else if var1Type == Int && var2Type == Float {
				v1 := vm.getValue(quad.Arg1).(int)
				v2 := vm.getValue(quad.Arg2).(float64)
				vm.setValue(quad.Result, float64(v1) != v2)
			} else if var1Type == Float && var2Type == Int {
				v1 := vm.getValue(quad.Arg1).(float64)
				v2 := vm.getValue(quad.Arg2).(int)
				vm.setValue(quad.Result, v1 != float64(v2))
			} else if var1Type == Bool && var2Type == Bool {
				v1 := vm.getValue(quad.Arg1).(bool)
				v2 := vm.getValue(quad.Arg2).(bool)
				vm.setValue(quad.Result, v1 != v2)
			} else {
				return fmt.Errorf("type mismatch in not equal comparison: %v != %v", var1Type, var2Type)
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
				vm.IP = quad.Arg2
				continue
			}
			
		case ERA:
			// Create a new local memory frame for the function
			vm.Locals = append(vm.Locals, NewMemory())
			// fmt.Printf("\nCreating new local memory frame: %+v\n", vm.Locals)

		case GOSUB:
			// Save the current instruction pointer to the call stack
			vm.Stack = append(vm.Stack, vm.IP+1)
			// Jump to the function address
			vm.IP = quad.Result
			continue

        case PARAM:
			// fmt.Printf("WHOLE PARAM: %+v\n", quad)
            val := vm.getValue(quad.Result)
            currentLocal := vm.getCurrentLocalMemory()
			// fmt.Printf("Setting parameter at index %d with value %v\n", quad.Arg2, val)
            // Determine where to store based on parameter index
            switch v := val.(type) {
            case int:
                currentLocal.Ints[quad.Arg2] = v
				// fmt.Printf("VALUE OF PARAMETER BEING SENT AND ADDRESS: %d, %d\n", v, quad.Arg2)
            case float64:
                currentLocal.Floats[quad.Arg2] = v
            case bool:
                currentLocal.Bools[quad.Arg2] = v
            }
			// fmt.Printf("Current local memory after PROCESS: %+v\n", currentLocal)


		case ENDFUNC:
            // Pop return address
            if len(vm.Stack) > 0 {
                vm.IP = vm.Stack[len(vm.Stack)-1]
                vm.Stack = vm.Stack[:len(vm.Stack)-1]
            }

            // Pop local memory
            if len(vm.Locals) > 0 {
                vm.Locals = vm.Locals[:len(vm.Locals)-1]
            }

			continue 

		case Print:
			outputStr := vm.PrintValue(quad.Arg1)
			fmt.Print(outputStr," ") // Print to console
			
		case DONE:
            return nil

		default:
            return fmt.Errorf("unknown operator %v", quad.Operator)
		}
		vm.IP++
	}
	return nil
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
		baseInt++
    }
    
    // Transfer constant floats
    baseFloat := 9000
    for _, val := range fd.MemoryManager.ConstFloats {
        vm.Constants.Floats[baseFloat] = val
		baseFloat++
    }

	    // Transfer constant strings
    baseString := 10000
    for _, val := range fd.MemoryManager.ConstStrings {
	//    fmt.Printf("Adding constant string: %s at address %d\n", val, baseString)
       vm.Constants.Strings[baseString] = val
	   baseString++
    }

}


// Add to your VirtualMachine struct
func (vm *VirtualMachine) PrintValue(address int) string {
    // Determine memory type and base address
    var value interface{}
    var found bool
    
    switch {
    // Global Integers (1000-1999)
    case address >= 1000 && address < 2000:
        value, found = vm.Globals.Ints[address]
    // Global Floats (2000-2999)
    case address >= 2000 && address < 3000:
        value, found = vm.Globals.Floats[address]
    // Local Integers (3000-3999)
    case address >= 3000 && address < 4000 && len(vm.Locals) > 0:
        value, found = vm.getCurrentLocalMemory().Ints[address]
    // Local Floats (4000-4999)
    case address >= 4000 && address < 5000 && len(vm.Locals) > 0:
        value, found = vm.getCurrentLocalMemory().Floats[address]
    // Temporal Bools (5000-5999)
    case address >= 5000 && address < 6000:
        value, found = vm.Temps.Bools[address]
    // Temporal Integers (6000-6999)
    case address >= 6000 && address < 7000:
        value, found = vm.Temps.Ints[address]
    // Temporal Floats (7000-7999)
    case address >= 7000 && address < 8000:
        value, found = vm.Temps.Floats[address]
    // Constant Integers (8000-8999)
    case address >= 8000 && address < 9000:
        value, found = vm.Constants.Ints[address]
    // Constant Floats (9000-9999)
    case address >= 9000 && address < 10000:
        value, found = vm.Constants.Floats[address]
    // Constant Strings (10000+)
    case address >= 10000:	
		// fmt.Printf("Accessing constant string at address %d\n", vm.Constants.Strings[address])
        value, found = vm.Constants.Strings[address]
    default:
        return "(invalid address)"
    }

    if !found {
        return "(undefined)"
    }

    // Format the value based on its type
    switch v := value.(type) {
    case int:
        return fmt.Sprintf("%d\n", v)
    case float64:
        // Remove trailing .0 if it's a whole number
        if v == float64(int(v)) {
            return fmt.Sprintf("%d\n", int(v))
        }
        return fmt.Sprintf("%.2f\n", v)
    case bool:
        if v {
            return "true"
        }
        return "false"
    case string:
        return v
    default:
		return fmt.Sprintf("%v\n", v)
    }
}