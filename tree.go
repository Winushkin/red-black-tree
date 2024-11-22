package main

import (
	"fmt"
)

type Node struct {
	parent      *Node
	value       int
	color       bool // true - red, false - black
	left, right *Node
}

func NewNode(value int, parent *Node) *Node {
	return &Node{
		parent: parent,
		value:  value,
		color:  true,
		left:   nil,
		right:  nil}
}

// rePaint

func (node *Node) RepaintBlack() {
	node.color = false //to black
}

func (node *Node) RepaintRed() {
	node.color = true // to red
}

// Tree struct

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

func (tree *Tree) RedUncleCase(X *Node) {
	if tree.RedUncleCaseCheck(X) {
		var F, G, U *Node
		F = X.parent
		G = F.parent

		if F == G.left {
			U = G.right // left case
		} else {
			U = G.left // right case
		}

		F.RepaintBlack()
		U.RepaintBlack()
		tree.RedUncleCase(G)
	}
}

func (tree *Tree) blackUncleLineCaseCheck(X *Node) bool {
	var F, G, U *Node
	F = X.parent

	if F.color == false {
		return false
	}

	G = F.parent
	if G == nil {
		return false
	}

	// left case
	if F == G.left {
		U = G.right
		if U == nil || !U.color {
			if X == F.left && F == G.left {
				return true
			}
		}

		// right case
	} else {
		U = G.left
		if U == nil || !U.color {
			if X == F.right && F == G.right {
				return true
			}
		}
	}

	return false
}

func (tree *Tree) blackUncleLineCase(X *Node) {
	if tree.blackUncleLineCaseCheck(X) {
		var F, G *Node
		F = X.parent
		G = F.parent

		if F == G.left { // left case
			G.left = F.right
			F.right = G
			F.parent = G.parent
			G.parent = F

		} else {
			G.right = F.left
			F.left = G
			F.parent = G.parent
			G.parent = F
		}

		if F.parent == nil {
			tree.root = F
		} else {
			if F.parent.left == G {
				F.parent.left = F
			} else {
				F.parent.right = F
			}
		}

		F.RepaintBlack()
		G.RepaintRed()

	}
}

func (tree *Tree) BlackUncleCaseCheck(X *Node) bool {
	var F, G, U *Node
	F = X.parent

	if F.color == false {
		return false
	}

	G = F.parent
	if G == nil {
		return false
	}

	// left case
	if F == G.left {
		U = G.right
		if U == nil || !U.color {
			if X == F.right && F == G.left {
				return true
			}
		}

		// right case
	} else {
		U = G.left
		if U == nil || !U.color {
			if X == F.left && F == G.right {
				return true
			}
		}
	}

	return false
}

func (tree *Tree) BlackUncleCase(X *Node) {
	if tree.BlackUncleCaseCheck(X) {
		var F, G *Node
		F = X.parent
		G = F.parent
		// left case
		if G.left == F {
			G.left = X
			X.left = F
			F.right = nil

		} else { // right case
			G.right = X
			X.right = F
			F.left = nil
		}

		X.parent = G
		F.parent = X

		tree.blackUncleLineCase(F)
	}
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

	tree.RedUncleCase(cur)
	tree.blackUncleLineCase(cur)
	tree.BlackUncleCase(cur)
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

func (tree *Tree) noChildrenRemove(X *Node) {
	if X == tree.root {
		tree.root = nil
	}

	F := X.parent
	if F.left == X {
		F.left = nil
	} else {
		F.right = nil
	}
}

func (tree *Tree) OneChildRemove(X *Node) {
	var Child *Node
	if X.right == nil { // left child
		Child = X.left
	} else { // right child
		Child = X.right
	}

	if X == tree.root {
		tree.root = Child
	}
	Child.parent = X.parent

	if Child.color && Child.parent.color {
		Child.RepaintBlack()
	}

}

func (tree *Tree) remove(value int) {
	X := tree.search(value)
	if X == nil {
		fmt.Println("No such node")
		return
	}

	if X.left == nil && X.right == nil {
		tree.noChildrenRemove(X)
	} else if (X.left == nil && X.right != nil) || (X.left != nil && X.right == nil) { //XOR
		tree.OneChildRemove(X)
	} else {
		//удаление ноды с 2-мя детьми
	}

}
