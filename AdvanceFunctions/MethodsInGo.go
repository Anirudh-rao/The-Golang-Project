package advancefunctions

import "fmt"

type rect struct {
	width, height int
}

// Method as a reciever indicated by *rect
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func MethodsInGo() {
	r := rect{width: 10, height: 30}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
