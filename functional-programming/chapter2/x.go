// chapter2
// First-class functions
//
// What "first-class citizen" means?
// We mean an entity (object, primitive or function) for which all the common language operations are available:
//   - Assignment
//   - Passing it to a function
//   - Returning from a function
//   - Storing it in another data type such as a map.
package chapter2

import "golang.org/x/exp/constraints"

// Type aliases for functions
type predicate[T any] func(T) bool

func filter1[T any](input []T, p predicate[T]) []T {
	arr := make([]T, 0, len(input))
	for _, v := range input {
		if p(v) {
			arr = append(arr, v)
		}
	}
	return arr
}

func filter2[T any] (input []T, p func(T) bool) []T {
	arr := make([]T, 0, len(input))
	for _, v := range input {
		if p(v) {
			arr = append(arr, v)
		}
	}
	return arr
}

func howToDeclareAFunc() {
	// inline
	var even predicate[int] = func(input int) bool {
		return input % 2 == 0
	}

	// anonymous (I mean, this is inline, but imagine that we are passing that function direclty as an argument to another function)
	_ = func(input int) bool {
		return input % 2 == 0
	}

	// Declaring function locally
	_ = even(2)

	// Creating the func
	_ = largetThanPredicateFactory(5)
	// Creating and executing the func with an argument
	_ = largetThanPredicateFactory(5)(6)
}

func even[T constraints.Integer](input T) bool {
	return input % 2 == 0	
}

// Creating functions that return another function
func largetThanPredicateFactory[T constraints.Integer](largerThan T) predicate[T] {
	return func(input T) bool {
		return input > largerThan
	}
}
