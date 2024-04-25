// Package gomodule contains stuff that is not very interesting
package gomodule

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers. To learn more about addition see https://www.mathsisfun.com/numbers/addition.html
func Add[N Number](num1 N, num2 N) N {
	return num1 + num2
}
