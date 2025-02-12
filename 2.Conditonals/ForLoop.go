package conditonals

import "fmt"

func ForLoop() {
	i := 1

	for i >= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//using range function
	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 2 {
			continue
		}
		fmt.Println(n)
	}
}
