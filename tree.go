package main

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

func RedUncleCaseCheck(X *Node) bool {
	var F, G, U *Node

	F = X.parent
	if F.color == false {
		return false
	}
	G = F.parent
	if G == nil {
		return false
	}

	if F == G.parent {
		U = G.right // left case
	} else {
		U = G.left // right case
	}

	if U == nil {
		return false
	}

	return U.color
}

func RedUncleCase(X *Node) {
	if RedUncleCaseCheck(X) {
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
		RedUncleCase(G)
	}
}

func (tree *Tree) FandGinRowCheck(X *Node) bool {
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

func (tree *Tree) FandGinRow(X *Node) {
	if tree.FandGinRowCheck(X) {
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
		G.RepaintBlack()

	}
}

func BlackUncleCaseCheck(X *Node) bool {
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

func BlackUncleCase(X *Node) {

	// add case3 to this case

}

func (tree *Tree) append(value int) {
	if tree.root == nil {
		tree.root = NewNode(value, nil)
		tree.root.color = false
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
			return
		}
	}

	RedUncleCase(cur)

	tree.root.RepaintBlack()

}
