package main

import (
	"bucheng.first/utils"
	"fmt"
)

func main() {
	//test1()
	//testWrite()
	//testRead()
	//writeMap()
	data := utils.ReadMap("testfile//map.txt")
	fmt.Println(data)
}

func writeMap() {
	data := make(map[string]string)
	data["name"] = "yinchong"
	data["age"] = "20"
	data["address"] = "wuhai"
	flag := utils.WriteMap("testfile//map.txt", data)
	fmt.Println(flag)
}

func testRead() {
	path := "testfile//test1.txt"
	content := utils.FileReadText(path)
	fmt.Println(content)
}

func testWrite() {
	path := "testfile//test1.txt"
	flag := utils.FileWriteText(path, "nice good haha")
	fmt.Println(flag)
}

func test1() {
	content := "nice"
	fmt.Println("content:" + content)
	changeContent(&content)
	fmt.Println("content:" + content)
	fmt.Println("Address1:", &content)
	printAddress(content)
	test := new(Test)
	test.PrintName()
	test.changeName2("hahaha")
	test.PrintName()
	test.changeName("hahaha")
	test.PrintName()
}

func changeContent(content *string) {
	*content = "lala"
}

func printAddress(content string) {
	fmt.Println("Address2:", &content)
}

type Test struct {
	name string
	age  int
}

func (test Test) PrintName() {
	fmt.Println("Test.....name:" + test.name)
}

func (test Test) changeName(name string) {
	test.name = name
}

func (test *Test) changeName2(name string) {
	test.name = name
}
