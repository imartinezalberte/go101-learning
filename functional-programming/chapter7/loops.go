package chapter7

type (
	Consumer[T any] func(T)
	Predicate[T any]  func(T) bool

	RangeOption func(Range) Range

	Range struct {
		from, to, steps int
	}
)

// from < to, steps (positive)
// from < to, steps (negative) Never reaches
// from > to, steps (negative)
// from > to, steps (positive) Never reaches
// steps = 0 Never ends
func StdLoop(from, to, steps int, accept Consumer[int]) {
	
}

func stdLoop(from, to, steps int, accept Consumer[int]) {
	if from >= to {
		return
	}

	accept(from)
	// StdLoop(from+steps)
}

func Loop(from, to, steps int, cond Predicate[int], accept Consumer[int]) {

}
