package outerPackage

import (
	"geekBangCourse/ch15/area"
	"testing"
)

func TestOuterPackage(t *testing.T) {
	area.RectangleArea(12, 5)
	area.CircleArea(5)
}
