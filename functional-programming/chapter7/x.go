// chapter7
//
// How to increase the stack size? debug.SetMaxStack(bytes) // -> 262144000 (limit for 32-bit machine) * 2 
//
// Go compiler does not support Tail-Call optimization, even if we write Tail-Call recursive functions.
package chapter7

import "math"

type (
	Node struct {
		Value int
		Left  *Node
		Right *Node
	}
)

func Factorial(input int) int {
	if input == 0 {
		return 1
	}
	return input * Factorial(input-1)
}

func NodeSumIt(root *Node) (sum int) {
	queue := make(chan *Node, 10)
	queue <- root

	for {
		select {
		case n := <-queue:
			sum += n.Value
			if n.Left != nil {
				queue <- n.Left
			}
			if n.Right != nil {
				queue <- n.Right
			}
		default:
			return
		}
	}
}

func NodeSumRecursive(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Value + NodeSumRecursive(root.Left) + NodeSumRecursive(root.Right)
}

func NodeMax(root *Node) int {
	if root == nil {
		return math.MinInt
	}
	return maxOf(root.Value, NodeMax(root.Left), NodeMax(root.Right))
}

func maxOf(input ...int) int {
	max := math.MinInt
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func AnotherNodeMax(root *Node) int {
	var (
		currentMax = math.MinInt
		inner      func(*Node)
	)

	inner = func(root *Node) {
		if root == nil {
			return
		}

		if root.Value > currentMax {
			currentMax = root.Value
		}

		inner(root.Left)
		inner(root.Right)
	}

	inner(root)
	return currentMax
}

func AnotherNodeMax1(root *Node) int {
	return anotherNodeMax1(root, math.MinInt)
}

func anotherNodeMax1(root *Node, currentMax int) int {
	if root == nil {
		return math.MinInt
	}

	if root.Value > currentMax {
		currentMax = root.Value
	}

	maxLeft := anotherNodeMax1(root.Left, currentMax)
	maxRight := anotherNodeMax1(root.Right, currentMax)
	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight
}
