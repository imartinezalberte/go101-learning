// chapter2
//
// Map dispatcher pattern
package chapter2

import (
	"errors"

	"golang.org/x/exp/constraints"
)

const (
	AdditionOperation       = "+"
	SubstractionOperation   = "-"
	MultiplicationOperation = "x"
	DivisionOperation       = "/"
)

type operation[T constraints.Float | constraints.Integer] func(T, T) T

// Example of usage
var (
	operationNotFound = errors.New("operation not found")

	intAddition       = AdditionOperationFactory[int]()
	intSubtraction    = SubtractionOperationFactory[int]()
	intMultiplication = MultiplicationOperationFactory[int]()
	intDivision       = DivisionOperationFactory[int]()

	operations = map[string]operation[int]{
		AdditionOperation:       intAddition,
		SubstractionOperation:   intSubtraction,
		MultiplicationOperation: intMultiplication,
		DivisionOperation:       intDivision,
	}
)

func AdditionOperationFactory[T constraints.Float | constraints.Integer]() operation[T] {
	return func(a, b T) T {
		return a + b
	}
}

func SubtractionOperationFactory[T constraints.Float | constraints.Integer]() operation[T] {
	return func(a, b T) T {
		return a - b
	}
}

func MultiplicationOperationFactory[T constraints.Float | constraints.Integer]() operation[T] {
	return func(a, b T) T {
		return a * b
	}
}

func DivisionOperationFactory[T constraints.Float | constraints.Integer]() operation[T] {
	return func(a, b T) T {
		if b == 0 {
			return 0
		}
		return a / b
	}
}

func OperationsFactory[T constraints.Float | constraints.Integer]() map[string]operation[T] {
	return map[string]operation[T]{
		AdditionOperation:       AdditionOperationFactory[T](),
		SubstractionOperation:   SubtractionOperationFactory[T](),
		MultiplicationOperation: MultiplicationOperationFactory[T](),
		DivisionOperation:       DivisionOperationFactory[T](),
	}
}

// We could also create a DispatchFactory so we use the same operations map every time we use this function
func Dispatch[T constraints.Float | constraints.Integer](operation string, a, b T) (T, error) {
	op, ok := OperationsFactory[T]()[operation]
	if !ok {
		return 0, operationNotFound
	}

	return op(a, b), nil
}
