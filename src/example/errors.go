package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// 自定义错误，实现接口error
// type error interface {Error() string} 标准库的接口
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s:%s", e.arg, "defined", e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "can't work with 42"}
	}

	return arg + 3, nil
}

func main() {
	for _, v := range []int{7, 42} {
		if r, e := f1(v); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)
		}
	}

	for _, v := range []int{7, 42} {
		if r, e := f2(v); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
