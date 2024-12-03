package main

import (
	"fmt"
)

// обход в глубину через рекурсию

type BinNode struct {
	value       int
	parent      *BinNode
	left, right *BinNode
}

func NewBinNode(value int) *BinNode {
	return &BinNode{value: value}
}

type BinaryTree struct {
	root *BinNode
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
		node.left = NewBinNode(value)

	}
}

func (tree *BinaryTree) PrintBinTree() {
	if tree.root == nil {
		fmt.Println("Дерево пустое")
		return
	}
	printBinNode(tree.root, "", true)
}

// Рекурсивная функция вывода узла
func printBinNode(node *BinNode, prefix string, isRight bool) {
	if node != nil {

		// Вывод текущего узла
		fmt.Printf("%s%s── %d\n", prefix, branch(isRight), node.value)

		// Обновляем префикс для дочерних узлов
		newPrefix := prefix + branchPrefix(isRight)
		printBinNode(node.left, newPrefix, false) // Левый ребенок
		printBinNode(node.right, newPrefix, true) // Правый ребенок
	}
}

// Обход в высоту (DFS)
//Прямой (Pre-Order, NLR)

func (tree *BinaryTree) PreOrder(node *BinNode) {
	if node != nil {
		fmt.Print(node.value, " ")
		tree.PreOrder(node.left)
		tree.PreOrder(node.right)
	}
}
