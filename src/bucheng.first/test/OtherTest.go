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
	//testReadMap()

	testTree()
	//var content *string
	//InitContent(content)
	//fmt.Println(content)
}

func InitContent(content *string) {
	*content = "nice"
}

func testTree() {
	var root BinaryTree
	root = BinaryTree{
		data: 55,
	}
	AddBinaryTreeNode(&root, 40)
	AddBinaryTreeNode(&root, 10)
	AddBinaryTreeNode(&root, 2)
	AddBinaryTreeNode(&root, 35)
	AddBinaryTreeNode(&root, 78)
	AddBinaryTreeNode(&root, 20)
	AddBinaryTreeNode(&root, 100)
	FirstPrintTree(&root)
	fmt.Println()
	BackPrintTree(&root)
	fmt.Println()
	MidlePrintTree(&root)
}

type BinaryTree struct {
	data  int
	prev  *BinaryTree
	left  *BinaryTree
	right *BinaryTree
}

func AddBinaryTreeNode(node *BinaryTree, data int) {
	if node == nil {
		node = &BinaryTree{
			data: data,
		}
		return
	}
	temp := node
	for {
		if data > temp.data {
			if temp.right != nil {
				temp = temp.right
			} else {
				temp.right = &BinaryTree{
					data: data,
					prev: temp,
				}
				return
			}
		}
		if data == temp.data {
			panic("data exist in tree")
		}
		if data < temp.data {
			if temp.left != nil {
				temp = temp.left
			} else {
				temp.left = &BinaryTree{
					data: data,
					prev: temp,
				}
				return
			}
		}
	}
}

func FirstPrintTree(node *BinaryTree) {
	if node == nil {
		return
	}
	FirstPrintTree(node.left)
	fmt.Print(" ", node.data)
	FirstPrintTree(node.right)
}

func MidlePrintTree(node *BinaryTree) {
	if node == nil {
		return
	}
	fmt.Print(" ", node.data)
	MidlePrintTree(node.left)
	MidlePrintTree(node.right)
}

func BackPrintTree(node *BinaryTree) {
	if node == nil {
		return
	}
	BackPrintTree(node.right)
	fmt.Print(" ", node.data)
	BackPrintTree(node.left)
}

func testReadMap() {
	data := utils.ReadMap("testfile//map.txt")
	fmt.Println(data)
}

func writeMap() {
	data := make(map[string]string)
	data["name2"] = "yinchong"
	data["age2"] = "20"
	data["address2"] = "wuhai"
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
