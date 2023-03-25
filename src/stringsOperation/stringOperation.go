package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	str1 = "my name is nan!"
	fmt.Printf("%s\n", str1)

	str2 := "test git"
	fmt.Printf("%d\n", strings.LastIndex(str2, "t"))
}
