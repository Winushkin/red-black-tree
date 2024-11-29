package main

// обход в глубину через рекурсию
// скобочная конструкция

type BinNode struct {
	value       int
	parent      *Node
	left, right *Node
}

func NewBinNode(value int) *BinNode {
	return &BinNode{value: value}
}

type BinaryTree struct {
	root *BinNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{root: nil}
}

func (tree *BinaryTree) insert(node *BinNode, value int) {

	if tree.root == nil {
		tree.root = &BinNode{value: value, parent: nil, left: nil, right: nil}
		return
	}

	if node == nil {
		node = tree.root
	}

	if node.left == nil {

	}

	//if tree.root == nil {
	//	tree.root = &BinNode{value: value, parent: nil, left: nil, right: nil}
	//	return
	//}
	//// исправить BST на BT
	//cur := tree.root
	//for {
	//	if value < cur.value {
	//		if cur.left != nil {
	//			cur = cur.left
	//		} else {
	//			cur.left = &BinNode{value: value, parent: cur, left: nil, right: nil}
	//		}
	//
	//	} else if cur.value > value {
	//		if cur.right != nil {
	//			cur = cur.right
	//		} else {
	//			cur.right = &BinNode{value: value, parent: cur, left: nil, right: nil}
	//		}
	//	}
	//}

}
