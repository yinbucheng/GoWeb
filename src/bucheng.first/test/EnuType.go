package main

import "fmt"

/**
go中枚举
*/
type ServiceCode int32

const (
	OK     = 200
	FAIL   = 500
	UNKNOW = 1000
)

func main() {
	fmt.Println(OK)
}
