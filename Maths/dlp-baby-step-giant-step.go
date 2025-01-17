package main

import (
	"fmt"
	"math")


func modex(base, exp, mod int) int {
    // modex calculates (base^exp) % mod efficiently.
    //Reduce the base using .
    //If the exponent is odd, multiply the result by the base.
    //Halve the exponent at each step, recalculating .
    //Repeat until the exponent is 0.

	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp / 2
		base = (base * base) % mod
	}
	return result
}

func babySgiantS(g, h, p int) int {
    // babyStepGiantStep computes the discrete logarithm x such that g^x ≡ h (mod p).
    // Returns x if it exists, or -1 if there is no solution.
	// Create the table of baby steps.
	// Calculate the factor for the giant steps.
	// Perform the giant steps.
	// Return the discrete logarithm.
    
    m := int(math.Ceil(math.Sqrt(float64(p))))

	babySteps := make(map[int]int)
	for j := 0; j < m; j++ {
		babySteps[modex(g, j, p)] = j
	}

	gInv := modex(g, p-2, p) // g Mod inverse (using Fermat).
	factor := modex(gInv, m, p)
	
	//Big steps
	current := h
	for i := 0; i < m; i++ {
		if val, exists := babySteps[current]; exists {
			return i*m + val
		}
		current = (current * factor) % p
	}
	
	return -1
}

func main() {
	g := 3     // Base
	h := 9  // Result
	p := 13  // Mod

	x := babySgiantS(g, h, p)
	if x != -1 {
		fmt.Printf("The discrete logarithm of %d in base %d modulo %d is: %d\n.", h, g, p, x)
	} else {
		fmt.Println("There is no solution for the given discrete logarithm.")
	}
}