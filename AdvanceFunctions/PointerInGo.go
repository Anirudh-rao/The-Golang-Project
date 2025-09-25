package advancefunctions

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func PonitersInGo() {
	i := 1
	fmt.Println("intial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroPtr(&i)
	fmt.Println("ZeroPtr:", i)

	fmt.Println("pointer:", &i)
}
