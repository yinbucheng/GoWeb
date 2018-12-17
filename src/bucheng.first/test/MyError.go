package main

import "fmt"

/**
go中自定义异常
*/
type MyError struct {
	error
}

func (this MyError) Error() string {
	return "自定义异常"
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				fmt.Println("panic-recover()等到是Error类型")
			}
			if _, ok := r.(MyError); ok {
				fmt.Println("painc-recover()等到是MyError类型")
			}
		}
	}()
	panic(MyError{})
}
