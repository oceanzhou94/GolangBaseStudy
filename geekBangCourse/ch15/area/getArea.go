package area

import "fmt"

const PI = 3

func RectangleArea(length, width int) {
	fmt.Println("rectangle area is", length*width)
}

func CircleArea(radio int) {
	fmt.Printf("circle area is %d\n", PI*radio*radio)
}
