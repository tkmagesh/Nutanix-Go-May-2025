/* Variadic functions */
/*
Hint : use strconv.Atoi() to convert string to int
*/

package main

import "fmt"

func main() {
	fmt.Println(sum(10))                      //=> 10
	fmt.Println(sum(10, 20))                  //=> 30
	fmt.Println(sum(10, "20", 30, 40))        //=> 100
	fmt.Println(sum(10, "20", 30, 40, "abc")) //=> 100
	fmt.Println(sum())                        //=> 0
}

func sum(list ...int) int { // list => []int (slice of int)
	var total int
	for i := 0; i < len(list); i++ {
		total += list[i]
	}
	return total
}
