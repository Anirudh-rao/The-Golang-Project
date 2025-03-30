package conditonals

import "fmt"

func ForLoop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("Range", i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
