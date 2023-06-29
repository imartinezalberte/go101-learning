// chapter4
//
// Pure functional languages: Haskell or Elm.
// Non-Pure functional languages: Clojure, Lisp or Erlang.
//
// What is Purity?
// - The target function does not generate any side effects
// - Returns the same output when providing the same input(idempotence)
//
// Referential transparency
// A func is referentially transparent if you can replace the func call with its output, without changing the result of the program.
// - Idempotence
// - Stateleness
// - Side Effects
package chapter4

import (
	"fmt"
	"math/rand"
)

func NonPureFunc1() {
	fmt.Println("Hello world") // This function generate a side effects. It displays a string on the screen.
}

func NonPureFunc2(input int) int {
	return rand.Intn(input) // This functions does not return the same output for the same input (it is not idempotence)
}

// Pure function
func Sum(a, b int) int {
	return a + b
}
