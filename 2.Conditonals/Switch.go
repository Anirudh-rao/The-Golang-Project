package conditonals

import (
	"fmt"
	"time"
)

func SwtichConditonals() {
	i := 2

	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is Weekend")
	default:
		fmt.Println("Today is WeekDay")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before 12")
	default:
		fmt.Println("Its after 12")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
