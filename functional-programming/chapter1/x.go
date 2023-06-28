// chapter1
//
// A function in FP is pure, it means that it cannot mutate the state of a system or produce any side effects.
// Hence a pure function is completely predictable.
//
// Go offers several features that enable us to write functional Go code with (relative) ease:
// - Functions as first-class citizens
// - Higher-order functions
// - Immutability guarantees
// - Generics (not needed per se, but make life easier)
// - Recursion
//
// Features that would improve our quality of life when building go solutions:
// - Tail-call optimization
// - Lazy evaluation
// - Purity guarantee
package chapter1

type (
	Predicate[T any] func(T) bool

	Person struct {
		name string
	}
)

func Filter[T any](input []T, test Predicate[T]) []T {
	result := make([]T, 0, len(input))
	for _, v := range input {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

// We are mutating the Person struct
func (p *Person) ChangeName(name string) {
	p.name = name
}

// We are creating a new Person struct so the "parent" one is not altered at all
func (p Person) SetName(name string) Person {
	p.name = name
	return p
}
