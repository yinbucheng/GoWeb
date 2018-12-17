package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f()
	func() {
		fmt.Println("error....2")
	}()
}

func f() {
	fmt.Println("panic invoke before")
	panic("this is error")
	fmt.Println("panic invoke after")
}
