package chapter1

import (
	"math"

	"golang.org/x/exp/constraints"
)

type (
	numberRange[T constraints.Integer] struct {
		arr []T
	}

	NumberStream[T constraints.Integer] interface {
		Filter(Predicate[T]) NumberStream[T]
		Abs() NumberStream[T]
		Sum() T
		ToArr() []T
	}

	Stream[T any] interface {
		Filter(Predicate[T]) Stream[T]
	}
)

func NumberRange[T constraints.Integer](from, to T) NumberStream[T] {
	return &numberRange[T]{arr: arrFromRange(from, to)}
}

func (r numberRange[T]) Filter(test Predicate[T]) NumberStream[T] {
	result := make([]T, 0, len(r.arr))

	for i := 0; i < len(r.arr); i++ {
		if test(r.arr[i]) {
			result = append(result, r.arr[i])
		}
	}
	
	return numberRange[T]{result}
}

func (r numberRange[T]) Abs() NumberStream[T] {
	for i := 0; i < len(r.arr); i++ {
		r.arr[i] = T(math.Abs(float64(r.arr[i])))
	}
	return r
}

func (r numberRange[T]) Sum() T {
	var sum T

	for i := 0; i < len(r.arr); i++ {
		sum += r.arr[i]
	}

	return sum
}

func (r numberRange[T]) ToArr() []T {
	return r.arr
}

func arrFromRange[T constraints.Integer](from, to T) []T {
	if from == to {
		return []T{}
	}

	if from > to {
		from, to = to, from
	}

	arr := make([]T, int(to-from) + 1)

	for i := 0; i < len(arr); i++ {
		arr[i] = from
		from++
	}

	return arr
}
