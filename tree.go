package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

func (tree *Tree) RedUncleCaseCheck(X *Node) bool {
	var F, G, U *Node

	F = X.parent

	if F == nil {
		return false
	}

	if F.color == false {
		return false
	}

	G = F.parent
	if G == nil {
		return false
	}

	if F == G.left {
		U = G.right // left case
	} else {
		U = G.left // right case
	}

	if U == nil {
		return false
	}

	return U.color
}

func (tree *Tree) PrintTree() {
	if tree.root == nil {
		fmt.Println("Дерево пустое")
		return
	}
	printNode(tree.root, "", true)
}

// Рекурсивная функция вывода узла
func printNode(node *Node, prefix string, isRight bool) {
	if node != nil {
		// Определяем цвет узла
		color := "R" // Красный
		if !node.color {
			color = "B" // Черный
		}

		// Вывод текущего узла
		fmt.Printf("%s%s── %d(%s)\n", prefix, branch(isRight), node.value, color)

		// Обновляем префикс для дочерних узлов
		newPrefix := prefix + branchPrefix(isRight)
		printNode(node.left, newPrefix, false) // Левый ребенок
		printNode(node.right, newPrefix, true) // Правый ребенок
	}
}

// Функция для определения символа ветви (левый или правый ребенок)
func branch(isRight bool) string {
	if isRight {
		return "└"
	}
	return "├"
}

// Функция для определения отступов для дочерних узлов
func branchPrefix(isRight bool) string {
	if isRight {
		return "   "
	}
	return "│  "
}

func (tree *Tree) insert(value int) {
	if tree.root == nil {
		tree.root = NewNode(value, nil)
		tree.root.color = false
		tree.PrintTree()
		return
	}

	var cur = tree.root
	for {
		if value < cur.value {
			if cur.left != nil {
				cur = cur.left
			} else {
				cur.left = NewNode(value, cur)
				cur = cur.left
				break
			}
		} else if cur.value < value {
			if cur.right != nil {
				cur = cur.right
			} else {
				cur.right = NewNode(value, cur)
				cur = cur.right
				break
			}
		} else {
			tree.PrintTree()
			return
		}
	}

	if tree.RedUncleCaseCheck(cur) {
		tree.RedUncleCase(cur)
	}

	if tree.blackUncleLineCaseCheck(cur) {
		tree.blackUncleLineCase(cur)
	}

	if tree.BlackUncleCaseCheck(cur) {
		tree.BlackUncleCase(cur)
	}

	tree.root.RepaintBlack()

	tree.PrintTree()
}

// searching

func (tree *Tree) search(value int) *Node {
	cur := tree.root
	if cur == nil {
		return nil
	}

	for {
		if value == cur.value {
			return cur
		} else if value < cur.value {
			if cur.left != nil {
				cur = cur.left
			} else {
				break
			}
		} else if cur.value < value {
			if cur.right != nil {
				cur = cur.right
			} else {
				break
			}
		}
	}
	return nil
}

// Removing elements

func (tree *Tree) childFreeRemove(X *Node) {
	if X == tree.root {
		tree.root = nil
		return
	}

	F := X.parent
	if F.left == X {
		if color(X) {
			F.left = nil
		} else {
			B := F.right
			if !color(B) {
				B.RepaintRed()
			}
			X.parent = nil
			F.left = nil
		}
	} else {
		if color(X) {
			F.right = nil
		} else {
			B := F.left
			if !color(B) {
				B.RepaintRed()
			}
			X.parent = nil
			F.right = nil
		}
	}

	if tree.BlackBroNefCaseCheck(X, F) {
		tree.BlackBroNefCase(X, F)
	}
	if tree.BlackBroRedRNefCaseCheck(X, F) {
		tree.BlackBroRedRNefCase(X, F)
	}
	if tree.BlackBroRNEFRedLNefCaseCheck(X, F) {
		tree.BlackBroRNEFRedLNefCase(X, F)
	}
	if tree.RedBroCaseCheck(X, F) {
		tree.RedBroCase(X, F)
	}
}

func (tree *Tree) OneChildRemove(X *Node) {
	var Child *Node

	if X.right == nil { // left child
		Child = X.left
	} else { // right child
		Child = X.right
	}
	F := X.parent

	//removing
	if X == tree.root {
		tree.root = Child
		Child.parent = X.parent
	} else {
		Child.parent = X.parent
		if F.left == X {
			F.left = Child
		} else {
			F.right = Child
		}
	}

	if color(Child) && color(Child.parent) {
		Child.RepaintBlack()
	}

	if tree.BlackBroNefCaseCheck(Child, F) {
		tree.BlackBroNefCase(Child, F)
	}
	if tree.BlackBroRedRNefCaseCheck(Child, F) {
		tree.BlackBroRedRNefCase(Child, F)
	}
	if tree.BlackBroRNEFRedLNefCaseCheck(Child, F) {
		tree.BlackBroRNEFRedLNefCase(Child, F)
	}
	if tree.RedBroCaseCheck(Child, F) {
		tree.RedBroCase(Child, F)
	}
}

func (tree *Tree) twoChildrenRemove(X *Node) {

}

func (tree *Tree) remove(value int) {
	X := tree.search(value)
	if X == nil {
		fmt.Println("No such node")
		return
	}

	if X.left == nil && X.right == nil {
		tree.childFreeRemove(X)
	} else if (X.left == nil && X.right != nil) || (X.left != nil && X.right == nil) { //XOR
		tree.OneChildRemove(X)
	} else {

		//удаление ноды с 2-мя детьми

	}

	if tree.BlackBroNefCaseCheck(X, F) {
		tree.BlackBroNefCase(X, F)
	}
	if tree.BlackBroRedRNefCaseCheck(X, F) {
		tree.BlackBroRedRNefCase(X, F)
	}
	if tree.BlackBroRNEFRedLNefCaseCheck(X, F) {
		tree.BlackBroRNEFRedLNefCase(X, F)
	}
	if tree.RedBroCaseCheck(X, F) {
		tree.RedBroCase(X, F)
	}

}
