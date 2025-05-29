// semantic/vm.go

/* 
1. replicar functionDir y MemoryManager de semantic
2. usar CuadruploList.Cuadruplos, functionDir y MemoryManager para ejecutar (simplemente pasando CuadruploList ya que contiene todo )
3. ejectuar cada cuadruplo en el orden que se encuentran
*/ 
package vm

import (
	"fmt"
	"babyduck/semantic"
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
	Cuadruplos []semantic.Cuadruplo
}

func NewMemory() Memory {
	return Memory{
		Ints:   make(map[int]int),
		Floats: make(map[int]float64),
		Bools:  make(map[int]bool),
	}
}

func NewVM(cl CuadruploList, consts Memory) *VirtualMachine {
	vm := &VirtualMachine{
		Globals:   NewMemory(),
		Locals:    []Memory{NewMemory()},
		Temps:     NewMemory(),
		Constants: consts,
		IP:        0,
		Stack:     []int{},
		Cuadruplos: cl.Cuadruplos,
	}
}

func (vm *VirtualMachine) getCurrentLocalMemory() Memory {
	return vm.Locals[len(vm.Locals)-1]
}

func (vm *VirtualMachine) getValue(address int) interface{} {
	switch {
	case address >= GlobalIntStart && address < GlobalFloatStart:
		return vm.Globals.Ints[address]
	case address >= GlobalFloatStart && address < LocalIntStart:
		return vm.Globals.Floats[address]
	case address >= LocalIntStart && address < LocalFloatStart:
		return vm.getCurrentLocalMemory().Ints[address]
	case address >= LocalFloatStart && address < TempIntStart:
		return vm.getCurrentLocalMemory().Floats[address]
	case address >= TempIntStart && address < TempFloatStart:
		return vm.Temps.Ints[address]
	case address >= TempFloatStart && address < TempBoolStart:
		return vm.Temps.Floats[address]
	case address >= TempBoolStart && address < ConstantStart:
		return vm.Temps.Bools[address]
	case address >= ConstantStart:
		if val, ok := vm.Constants.Ints[address]; ok {
			return val
		} else if valf, ok := vm.Constants.Floats[address]; ok {
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
		if address >= GlobalIntStart && address < GlobalFloatStart {
			vm.Globals.Ints[address] = v
		} else if address >= LocalIntStart && address < LocalFloatStart {
			mem := vm.getCurrentLocalMemory()
			mem.Ints[address] = v
		} else if address >= TempIntStart && address < TempFloatStart {
			vm.Temps.Ints[address] = v
		} else if address >= ConstantStart {
			vm.Constants.Ints[address] = v
		}
	case float64:
		if address >= GlobalFloatStart && address < LocalIntStart {
			vm.Globals.Floats[address] = v
		} else if address >= LocalFloatStart && address < TempIntStart {
			mem := vm.getCurrentLocalMemory()
			mem.Floats[address] = v
		} else if address >= TempFloatStart && address < TempBoolStart {
			vm.Temps.Floats[address] = v
		} else if address >= ConstantStart {
			vm.Constants.Floats[address] = v
		}
	case bool:
		if address >= TempBoolStart && address < ConstantStart {
			vm.Temps.Bools[address] = v
		} else if address >= ConstantStart {
			vm.Constants.Bools[address] = v
		}
	default:
		fmt.Printf("Unknown type for value: %v\n", value)
	}
}

func (vm *VirtualMachine) Run(cuads []Cuadruplo, consts Memory) {
	
	for vm.IP < len(vm.Cuadruplos) {
		quad := vm.Cuadruplos[vm.IP]
		switch quad.Operator {
		case ADD:
			v1 := vm.getValue(quad.Arg1).(int)
			v2 := vm.getValue(quad.Arg2).(int)
			vm.setValue(quad.Result, v1+v2)
		case SUB:
			v1 := vm.getValue(quad.Arg1).(int)
			v2 := vm.getValue(quad.Arg2).(int)
			vm.setValue(quad.Result, v1-v2)
		case MUL:
			v1 := vm.getValue(quad.Arg1).(int)
			v2 := vm.getValue(quad.Arg2).(int)
			vm.setValue(quad.Result, v1*v2)
		case DIV:
			v1 := vm.getValue(quad.Arg1).(int)
			v2 := vm.getValue(quad.Arg2).(int)
			if v2 == 0 {
				panic("Runtime error: division by zero")
			}
			vm.setValue(quad.Result, v1/v2)
		case ASSIGN:
			val := vm.getValue(quad.Arg1)
			vm.setValue(quad.Result, val)
		case GOTO:
			vm.IP = quad.Result
			continue
		case GOTOF:
			cond := vm.getValue(quad.Arg1).(bool)
			if !cond {
				vm.IP = quad.Result
				continue
			}
		// Add more cases like PRINT, READ, ERA, PARAM, GOSUB, RETURN, etc.
		default:
			fmt.Printf("Unknown operator: %v\n", quad.Operator)
		}
		vm.IP++
	}
}
