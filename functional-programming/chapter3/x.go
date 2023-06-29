// chapter3
//
// Here we are going to explore the concept of function composition through higher-order functions.
//
// - closures
// - partial applications
// - function currying
//
// What is a higher-order function? it is any function that either takes a function as the input or returns a function as the output.
// What is a closure? Any inner function that uses a variable introduced in the outer function to perform its work.
// Function curry:
// func F(a, b, c): int {} ->
//   - func Fa(a): Fb(b)
//   - func Fb(b): Fc(c)
//   - func Fc(c): int
package chapter3

import (
	"fmt"
	"strings"
)

func a() string {
	return "Hello world"
}

func execute(a func() string) string {
	return a()
}

func lexicalScopes() {
	{
		b := true
		_ = b // b declared but no used
	}

	var b bool // Imagine that this line doesn't exists, I have jotted it down because I need the code to compile.
	if b {     // undefined b
		fmt.Println("b is already declared, so cool")
	}
}

func lexicalScope1() {
	s := "Hello"
	{ // Different lexical scope
		s := "World"
		fmt.Println(s) // It will display the word "World"
	}
	fmt.Println(s) // It will display the word "Hello"
}

func lexicalScope2() {
	s := "Hello"
	{
		s = "World"
		fmt.Println(s) // It will display the word "World"
	}
	fmt.Println(s) // It will also display the word "World"
}

func lexicalScope3() {
	s := "Hello"
	// s := "World" // No new variables on the left side, so... Compilation error
	fmt.Println(s) // It will display Hello
}

type GreetingFn func(string) string

func greetingFactory1() GreetingFn {
	baseGreeting := "Hello"
	return func(input string) string { // This is a closure because we are using a variable that is out of the inner function block
		return strings.Join([]string{baseGreeting, input}, " ")
	}
}

func greetingFactory2(baseGreeting string) GreetingFn {
	return func(input string) string { // This is a partial application
		return strings.Join([]string{baseGreeting, input}, " ")
	}
}

func functionCurryExample() {
	threeSum := func(a, b, c int) int {
		return a + b + c
	}

	threeSumUsingFCurrying := func(a int) func(int) func(int) int {
		return func(b int) func(int) int {
			return func(c int) int {
				return a + b + c
			}
		}
	}

	_ = threeSum(10, 20, 30)
	_ = threeSumUsingFCurrying(10)(20)(30)
}
