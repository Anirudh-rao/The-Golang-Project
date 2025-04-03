package functions

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}

func Function() {
	res := plus(12, 13)
	fmt.Println("1+2=", res)

	res = plusPlus(13, 12, 14)
	fmt.Println("13+12+14 = ", res)
}
