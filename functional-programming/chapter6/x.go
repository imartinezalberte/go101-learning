package chapter6

import "golang.org/x/exp/constraints"

type (
	Predicate[T any] func(T) bool
	Mapper[T, K any] func(T) K
	ReducerFn[T any] func(T, T) T
)

var (
	_ = Filter([]int{1, 2, 3, 4}, func(input int) bool {
		return input%2 == 0
	})

	_ = RudimentaryFilter([]int{1, 10, 20, 5}, 10)
)

// TODO: test this function
func Filter[T any](input []T, test Predicate[T]) []T {
	result := make([]T, 0, len(input))
	for _, v := range input {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

// TODO: test this function
func All[T any](input []T, test Predicate[T]) bool {
	for _, v := range input {
		if !test(v) {
			return false
		}
	}
	return true
}

// TODO: test this function
func Any[T any](input []T, test Predicate[T]) bool {
	for _, v := range input {
		if test(v) {
			return true
		}
	}
	return false
}

func DropWhile[T any](input []T, test Predicate[T]) []T {
	for i := 0; i < len(input); i++ {
		if !test(input[i]) {
			return input[i:]
		}
	}
	return []T{} // empty
}

func TakeWhile[T any](input []T, test Predicate[T]) []T {
	for i := 0; i < len(input); i++ {
		if !test(input[i]) {
			return input[:i]
		}
	}
	return input
}

func Map[T, K any](input []T, mapFn Mapper[T, K]) []K {
	result := make([]K, len(input))
	for i, v := range input {
		result[i] = mapFn(v)
	}
	return result
}

// TODO: test this function
func FlatMap[T any](input []T, mapFn Mapper[T,[]T]) []T {
	result := make([]T, 0, len(input))
	for _, v := range input {
		result = append(result, mapFn(v)...)
	}
	return result
}

func Reduce[T any](input []T, fn ReducerFn[T]) T {
	var first T
	if len(input) == 0 {
		return first
	}

	first = input[0]
	for _, second := range input[1:] {
		first = fn(first, second)
	}

	return first
}

// Same as Reduce
func ReduceLazy[T any](input []T, fn ReducerFn[T]) T {
	return ReduceWithStart(input, *new(T), fn)
}

func ReverseReducer[T any](input []T, fn ReducerFn[T]) T {
	return ReverseReducerWithStart(input, *new(T), fn)	
}

func ReduceWithStart[T any](input []T, start T, fn ReducerFn[T]) T {
	if len(input) == 0 {
		return start
	}
	for _, next := range input {
		start = fn(start, next)
	}
	return start
}

func ReverseReducerWithStart[T any](input []T, start T, fn ReducerFn[T]) T {
	if len(input) == 0 {
		return start
	}
	for i := len(input)-1; i >=0; i-- {
		start = fn(start, input[i])
	}
	return start
}

// Conditions are always changing, so we would end up with a ton of filter functions, repeating code and pulling our hair out.
func RudimentaryFilter[T constraints.Integer | constraints.Float](input []T, threshold T) []T {
	result := make([]T, 0, len(input))
	for _, v := range input {
		if v > threshold {
			result = append(result, v)
		}
	}
	return result
}
