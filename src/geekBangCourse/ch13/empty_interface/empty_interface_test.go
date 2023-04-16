package empty_interface

import (
	"fmt"
	"testing"
)

//类型断言
func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Interger", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("Unknow types")
	switch v := p.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(true)
}
