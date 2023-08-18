/*基础语法之变量*/
package main

import (
	"fmt"
	"math"
)

//全局变量的声明必须使用var关键字、
var gloableVariable string = "i'm gloableVariable"

func main() {
	var (
		a int //有符号整型一共5个类型，一般开发都用int，其位数由操作系统决定
		b int8
		c int16
		d int32
		e int64
		f float32 //64位浮点数精度比32为更高
		g float64
		h bool       //book类型，包含两个值，false、true
		i complex64  //实部和虚部都是32位浮点数组成的复数
		j complex128 //实部和虚部都是64位浮点数组成的复数
		s string     //字符串类型

	)

	//打印类型,格式化字符串%T，表示打印数据类型
	fmt.Printf("a: %T max: %v\n", a, math.MaxInt)
	fmt.Printf("b: %T max: %v\n", b, math.MaxInt8)
	fmt.Printf("c: %T max: %v\n", c, math.MaxInt16)
	fmt.Printf("d: %T max: %v\n", d, math.MaxInt32)
	fmt.Printf("e: %T max: %v\n", e, math.MaxInt64)
	fmt.Printf("f: %T max: %v\n", f, math.MaxFloat32)
	fmt.Printf("g: %T max: %v\n", g, math.MaxFloat64)
	fmt.Printf("h: %T\n", h)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("j: %T\n", j)
	fmt.Printf("s: %T\n", s)

	//速记声明，不需要声明变量的类型，编译器自动根据变量的值识别数据类型
	var1 := 5
	var2 := 18.4
	var3 := true
	var4 := "name is string"
	var5 := complex(13.4, 28.4)
	fmt.Printf("var1: %T\n", var1)
	fmt.Printf("var2: %T\n", var2)
	fmt.Printf("var3: %T\n", var3)
	fmt.Printf("var4: %T\n", var4)
	fmt.Printf("var5: %T\n", var5)

	fmt.Println("gloable variable:", gloableVariable)

}
