package functions

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func MultipleFunctions() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//If we want to use subset of returned
	//valus we use ' _ ' .
	_, c := vals()
	fmt.Println(c)
}
