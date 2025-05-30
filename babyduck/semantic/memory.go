package semantic

import (
	// "fmt"
)


/*
Hacer division de indices para tipos de vars: 
	a. Globales
		- int 
		- float
	b. Locales
		- int 
		- float
	c. Temporales
		- bool
		- int
		- flotantes
	d. constantes (sugerencia crear su propia tabla para no confundir con enteros)
		- string
		- int 
		- float
*/

type MemoryType int

const (
	FunctionsAddress MemoryType = iota
	GlobalInt
	GlobalFloat
	LocalInt
	LocalFloat
	TemporalBool
	TemporalInt
	TemporalFloat
	ConstantInt
	ConstantFloat
	ConstantString
)

type MemoryManager struct {
	Memory_address map[MemoryType]int //mapa que lleva el registro de las direcciones disponibles para cada tipo de var. Es la pieza clave para la asignación automática de direcciones virtuales.
	ConstInts  []int
	ConstFloats []float64
	ConstStrings []string
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		Memory_address: map[MemoryType]int{
			FunctionsAddress: 100, // Asignación inicial para funciones
			GlobalInt: 1000,
			GlobalFloat: 2000,
			LocalInt: 3000,
			LocalFloat: 4000,
			TemporalBool: 5000,
			TemporalInt: 6000,
			TemporalFloat: 7000,
			ConstantInt: 8000,
			ConstantFloat: 9000,
			ConstantString: 10000,
		},
		ConstInts: []int{},
		ConstFloats: []float64{},
		ConstStrings: []string{""},
	}
}

/*
---- guardar valores de constantes con su dirección virtual (posible manera) ----

type Constant struct {
    Value   interface{}
    Address int
}

type MemoryManager struct {
    Memory_address map[MemoryType]int
    constants      map[MemoryType][]Constant
}

func NewMemoryManager() *MemoryManager {
    return &MemoryManager{
        Memory_address: map[MemoryType]int{
            ConstantInt:    8000,
            ConstantFloat:  9000,
            ConstantString: 10000,
        },
        constants: map[MemoryType][]Constant{
            ConstantInt:    {},
            ConstantFloat:  {},
            ConstantString: {},
        },
    }
}

func (mm *MemoryManager) AddConstantInt(value int) {
    addr := mm.Memory_address[ConstantInt] + len(mm.constants[ConstantInt])
    mm.constants[ConstantInt] = append(mm.constants[ConstantInt], Constant{
        Value:   value,
        Address: addr,
    })
}

// Similar methods for AddConstantFloat, AddConstantString
*/

func (mm *MemoryManager) Allocate(memType MemoryType) int {
	address := mm.Memory_address[memType]
	mm.Memory_address[memType]++
	return address
}

func (mm *MemoryManager) ConvertVarType_ToMemoryType(typ Type) MemoryType {
	switch typ {
	case Int:
		return LocalInt
	case Float:
		return LocalFloat
	default:
		panic("tipo no soportado para memoria")
	}
}

func (mm *MemoryManager) ResetLocalVariables() {
	mm.Memory_address[LocalInt] = 3000
	mm.Memory_address[LocalFloat] = 4000
	mm.Memory_address[TemporalBool] = 5000
	mm.Memory_address[TemporalInt] = 6000
	mm.Memory_address[TemporalFloat] = 7000
}