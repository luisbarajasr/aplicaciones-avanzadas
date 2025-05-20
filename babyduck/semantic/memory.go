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
	GlobalInt MemoryType = iota
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
	memory_address map[MemoryType]int //mapa que lleva el registro de las direcciones disponibles para cada tipo de var. Es la pieza clave para la asignación automática de direcciones virtuales.
	constInts  []int
	constFloats []float64
	constStrings []string
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		memory_address: map[MemoryType]int{
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
		constInts: []int{0},
		constFloats: []float64{0},
		constStrings: []string{""},
	}
}

func (mm *MemoryManager) Allocate(memType MemoryType) int {
	address := mm.memory_address[memType]
	mm.memory_address[memType]++
	return address
}

// func (mm *MemoryManager) AllocateConst(value interface{}) int {
// 	switch v := value.(type) {
// 	case Int: 

// 	case Float:

// 	case String: 

// 	default:
// 	}
// }