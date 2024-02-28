package main

import "fmt"

func main() {
	var a, b, c int
	var radius float64

	fmt.Println("Enter three numbers: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	average(a, b, c)

	min, max := min_max(a, b, c)

	fmt.Println("Minimum is ", min)
	fmt.Println("Maximum is ", max)

	fmt.Println("Enter radius of the circle: ")
	fmt.Scanln(&radius)

	area := calculate_circle_area(radius)

	fmt.Println("Area of the circle is ", area)

}
