package calculator

import "fmt"

// private
var opCount map[string]int

// init
func init() {
	fmt.Println("[calculator] - init fn")
	opCount = make(map[string]int, 0)
}

// public
func GetOpCount() map[string]int {
	return opCount
}
