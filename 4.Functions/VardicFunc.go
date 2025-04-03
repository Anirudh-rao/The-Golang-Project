//Vardic functions can be called with any number of trailing arguments

package functions

import "fmt"

func sum(num ...int) {
	fmt.Println(num, " ")
	total := 0

	for _, num := range num {
		total += num
	}

	fmt.Println(total)
}

func VardicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	num := []int{1, 2, 3, 4}
	sum(num...)
}
