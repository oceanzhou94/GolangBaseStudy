package rectangle

import (
	"fmt"
)

func Area(length, width float64) {
	area := length * width
	fmt.Printf("the area is: %.2f", area)
}
