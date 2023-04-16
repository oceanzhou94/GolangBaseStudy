package error_test

import (
	"errors"
	"fmt"
	"testing"
)

//斐波那契数的递归实现
func GetFibonacci(n int) ([]int, error) {
	//错误处理
	if n < 0 || n > 100 {
		return nil, errors.New("n should be in [2,100]")
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestTitle(t *testing.T) {
	var x = []int{1, 5: 1, 2}
	fmt.Println(len(x))
}
