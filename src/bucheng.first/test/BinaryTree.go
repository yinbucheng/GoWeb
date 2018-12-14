package main

type BinaryTree struct {
	data  int
	prev  *BinaryTree
	left  *BinaryTree
	right *BinaryTree
}

type Root struct {
	root *BinaryTree
}

func (root *Root) Add(data int) {
	if root.root == nil {
		root.root = &BinaryTree{
			data: data,
		}
	} else {
		temp := root.root
		for {
			if data >= temp.data {
				if temp.right != nil {
					temp = temp.right
				} else {
					temp.right = &BinaryTree{
						data: data,
						prev: temp,
					}
					break
				}
			}
			if data < temp.data {
				if temp.left != nil {
					temp = temp.left
				} else {
					temp.left = &BinaryTree{
						data: data,
						prev: temp,
					}
					break
				}
			}
		}
	}
}

func (root *Root) Remove(data int) {
	if root.root == nil {
		return
	}
	temp := root.root
	for {
		if data >= temp.data {
			if temp.right != nil {
				temp = temp.right
			} else {
				break
			}
		}

		if data < temp.data {
			if temp.left != nil {
				temp = temp.left
			} else {
				break
			}
		}

		if data == temp.data {

		}
	}
}
