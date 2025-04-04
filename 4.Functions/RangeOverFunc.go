package functions

import "fmt"

func RangOverFunction() {
	nums := []int{1, 2, 3}
	sums := 0

	for _, num := range nums {
		sums += num
	}
	fmt.Println("sum:", sums)

	for i, nums := range nums {
		if nums == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
