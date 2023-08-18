package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

//可变长参数
func Sum(opts ...int) int {
	sum := 0
	for _, v := range opts {
		sum += v
	}
	return sum
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 3, 5, 7))
	t.Log(Sum(1, 2, 3, 4))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer : clear resources")
	}()
	t.Log("Start")
	//panic异常后还是会执行defer
	panic("error")

}
