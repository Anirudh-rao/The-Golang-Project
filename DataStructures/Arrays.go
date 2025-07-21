package datastructures

import "fmt"

func ArraysInGo() {
	var a [5]int
	fmt.Println("Empty Array", a)

	a[4] = 100
	fmt.Println("Set", a)
	fmt.Println("Get", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// We can also as the compile to count the Number of element
	b = [...]int{1, 2, 3, 3, 4}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
