/*基础语法之变量*/
package main

import "fmt"

func main() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f float32
		g float64
		s string
	)

	//打印类型,格式化字符串%T，表达打印数据类型
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("d: %T\n", d)
	fmt.Printf("e: %T\n", e)
	fmt.Printf("f: %T\n", f)
	fmt.Printf("g: %T\n", g)
	fmt.Printf("s: %T\n", s)

}
