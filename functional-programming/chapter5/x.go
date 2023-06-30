// chapter5
//
// # Learning about immutability
//
// Garbage collector
//
//		Concurrent mark-and-sweep garbage collector.
//
//	 A stack is a Last-In, First-Out (FIFO) data structure.
//
//	 Compiler will try to prove that a variable is local to a single function (process "escape analysis",
//	 where it looks for variables that escape the context of a single function). If a variable is not local to a single function, it stores it on
//		the heap.
//
//	 This command is to tell the compiler to explain where variables escapte the context where they lay in.
//		Use the pragma //go:noinline in a function where Go is probably inlining functions to optimize the code.
//	 go build -gcflags '-m -l'
//
// What's a functor? It is a func that can apply an operation to each element contained in a data structure. It's name in Haskel is fmap
// What's a monad? It is a software design pattern. It is a data type that can combine functions of similar types and wrap the results of a
// non-nomad type into a new monadic type offering additional functions.
//
// - A function to wrap a value of the T type into Monad[T]
// - A function to combine the function of the Monad[T] type
package chapter5

import (
	"fmt"
	"strconv"
)

type (
	person struct {
		name string
		age  uint
	}

	PersonFn func(person) person
)

var _ = Person(PersonName("Ivan"), PersonAge(25))

func Person(setters ...PersonFn) person {
	p := person{}

	for _, set := range setters {
		p = set(p)
	}

	return p
}

func PersonName(name string) PersonFn {
	return func(p person) person {
		p.name = name
		return p
	}
}

func PersonAge(age uint) PersonFn {
	return func(p person) person {
		p.age = age
		return p
	}
}

func Example() {
	type person struct {
		name string
		age  uint
	}

	// We are not able to modify the person struct because we are passing-by-value
	{
		var (
			p        person
			modifier = func(p person, name string) person {
				p.name = name
				return p
			}
		)

		{
			p = modifier(p, "not setted name")

			fmt.Println(p) // Will print {"not setted name" 0}
		}

		fmt.Println(p) // Will print { 0}
	}

	// We are able to modify the person struct because we are using pointers, so we are passing-by-reference
	{
		var (
			p        person
			modifier = func(p *person, name string) person {
				p.name = name
				return *p
			}
		)

		{
			p = modifier(&p, "setted name")

			fmt.Println(p) // Will print {"setted name" 0}
		}

		fmt.Println(p) // Will print {"setted name" 0}
	}
}

func fmap[A, B any](mapFunc func(A) B, input []A) []B {
	result := make([]B, len(input))
	for i, a := range input {
		result[i] = mapFunc(a)
	}
	return result
}

func ExampleFmap() {
	input := []int{1, 2, 3}

	_ = fmap(strconv.Itoa, input)
}
