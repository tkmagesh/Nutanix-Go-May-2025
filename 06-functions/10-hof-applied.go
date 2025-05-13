package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// ver 2.0
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	// ver 3.0
	/*
		logAdd := wrapLogger(add)
		logAdd(100, 200)

		logSubtract := wrapLogger(subtract)
		logSubtract(100, 200)
	*/
	/*
		add := wrapLogger(add)
		subtract := wrapLogger(subtract)
	*/

	/*
		logAdd := wrapLogger(add)
		add := wrapProfiler(logAdd)

		logSubtract := wrapLogger(subtract)
		subtract := wrapProfiler(logSubtract)
	*/

	add := wrapProfiler(wrapLogger(add))
	subtract := wrapProfiler(wrapLogger(subtract))

	add(100, 200)
	subtract(100, 200)
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// ver 2.0
/*
func logAdd(x, y int) {
	log.Println("Operation started...")
	add(x, y)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started...")
	subtract(x, y)
	log.Println("Operation completed!")
}
*/
// After refactoring the above to avoid duplication
func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started...")
	op(x, y)
	log.Println("Operation completed!")
}

// ver 3.0
func wrapLogger(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started...")
		op(x, y)
		log.Println("Operation completed!")
	}
}

// ver 4.0
func wrapProfiler(fn func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		start := time.Now()
		fn(i1, i2)
		elapsed := time.Since(start)
		fmt.Printf("Time Taken : %v\n", time.Duration(elapsed))
	}
}
