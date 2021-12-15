package main

import "fmt"

// Source: https://go.dev/doc/tutorial/generics
// Use: go1.18beta1 or higher to execute this code!

type Number interface {
	int64 | float64
}

func ExecuteGenerics() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first": 34,
		"second": 12,
	}
	
	// Initialize a map for the float values
	floats := map[string]float64{
		"first": 35.98,
		"second": 26.99,
	}
	
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers[string, int64](ints),
		SumNumbers[string, float64](floats),
	)
	 
	fmt.Printf("Generic Sums with Inferred Types and Interface Constrants: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

/**
	Declare a SumIntsOrFloats function with two type parameters (inside the square brackets), K and V, and one argument that uses the type parameters,
	m of type map[K]V. The function returns a value of type V.
	Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these, the comparable constraint is predeclared in Go.
	It allows any type whose values may be used as an operand of the comparison operators == and !=. Go requires that map keys be comparable.
	So declaring K as comparable is necessary so you can use K as the key in the map variable. It also ensures that calling code uses an allowable type for map keys.
	Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using | specifies a union of the two types, 
	meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the calling code.
	Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters. 
	Note that we know map[K]V is a valid map type because K is a comparable type. If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.
*/
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}