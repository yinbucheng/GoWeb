package main

//
//type BinaryTree struct {
//	data  int
//	prev  *BinaryTree
//	left  *BinaryTree
//	right *BinaryTree
//}
//
//func AddBinaryTreeNode(node *BinaryTree,data int)  {
//	if node==nil{
//		*node = BinaryTree{
//			data:data,
//		}
//	}
//	temp:=node
//	for ;;{
//		if data>temp.data{
//			if temp.right!=nil{
//				temp=temp.right
//			}else{
//				*temp.right =BinaryTree{
//					data:data,
//					prev:temp,
//				}
//			}
//		}
//		if data==temp.data{
//			panic("data exist in tree")
//		}
//		if data<temp.data{
//			if temp.left!=nil{
//				temp=temp.left
//			}else{
//				*temp.left=BinaryTree{
//					data:data,
//					prev:temp,
//				}
//			}
//		}
//	}
//}
//
//func FirstPrintTree(node *BinaryTree)  {
//	if node==nil{
//		return
//	}
//	FirstPrintTree(node.left)
//	fmt.Print("%b",node.data)
//	FirstPrintTree(node.right)
//}
//
//func MidlePrintTree(node *BinaryTree)  {
//	if node==nil{
//		return
//	}
//	fmt.Print("%b",node.data)
//	MidlePrintTree(node.left)
//	MidlePrintTree(node.right)
//}
//
//func BackPrintTree(node *BinaryTree)  {
//	if node==nil{
//		return
//	}
//	MidlePrintTree(node.left)
//	MidlePrintTree(node.right)
//	fmt.Print("%b",node.data)
//}
