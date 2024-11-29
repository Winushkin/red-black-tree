package main

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
		G.RepaintRed()
		tree.RedUncleCase(G)
	}
}

func (tree *Tree) blackUncleLineCaseCheck(X *Node) bool {
	var F, G, U *Node
	F = X.parent

	if !color(F) {
		return false
	}

	G = F.parent
	if G == nil {
		return false
	}
	if color(G) {
		return false
	}

	// left case
	if F == G.left {
		U = G.right
		if !color(U) {
			if X == F.left {
				return true
			}
		}

		// right case
	} else {
		U = G.left
		if !color(U) {
			if X == F.right {
				return true
			}
		}
	}

	return false
}

func (tree *Tree) blackUncleLineCase(X *Node) {
	var F, G *Node
	F = X.parent
	G = F.parent

	if F == G.left { // left case
		G.left = F.right
		if G.left != nil {
			G.left.parent = G
		}
		F.right = G
		F.parent = G.parent
		G.parent = F

	} else { // right case
		G.right = F.left
		if G.right != nil {
			G.right.parent = G
		}
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

func (tree *Tree) BlackUncleCaseCheck(X *Node) bool {
	var F, G, U *Node
	F = X.parent

	if !color(F) {
		return false
	}

	G = F.parent
	if G == nil {
		return false
	}
	if color(G) {
		return false
	}

	// left case
	if F == G.left {
		U = G.right
		if !color(U) {
			if X == F.right {
				return true
			}
		}

		// right case
	} else {
		U = G.left
		if !color(U) {
			if X == F.left {
				return true
			}
		}
	}

	return false
}

func (tree *Tree) BlackUncleCase(X *Node) {
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

// removing

// BlackBroNephewsCaseCheck Case 1
func (tree *Tree) BlackBroNephewsCaseCheck(X *Node) bool {
	F := X.parent
	if F == nil {
		// обработать ситуацию
		return false
	}
	var B, N1, N2 *Node

	if F.left == X {
		B = F.right
		N1, N2 = B.left, B.right
	} else {
		B = F.left
		N1, N2 = B.left, B.right
	}

	if !B.color && !N1.color && !N2.color {
		return true
	}

	return false
}

func (tree *Tree) BlackBroNefCase(X *Node) {

	if tree.BlackBroNephewsCaseCheck(X) {
		F := X.parent
		var B *Node

		if F.left == X {
			B = F.right
		} else {
			B = F.left
		}

		B.RepaintRed()
		if !F.color {
			F.RepaintBlack()
			tree.BlackBroNefCase(F)
		}

		F.RepaintBlack()
	}
}

// BlackBroRedRNefCaseCheck Case 2
func (tree *Tree) BlackBroRedRNefCaseCheck(X *Node) bool {
	F := X.parent
	if F == nil {
		return false
	}

	var B, N2 *Node

	if F.left == X {
		B = F.right
		N2 = B.right
	} else {
		B = F.left
		N2 = B.left
	}

	if !B.color && N2.color {
		return true
	}
	return false
}

func (tree *Tree) BlackBroRedRNefCase(X *Node) {
	if tree.BlackBroRedRNefCaseCheck(X) {
		F := X.parent

		var B, N1, N2 *Node
		var leftCase = false

		if F.left == X {
			leftCase = true
		}

		if leftCase {
			B = F.right
			N1, N2 = B.left, B.right
			B.left = F
			F.parent = B
			F.right = N1
			N1.parent = F
			N2.RepaintBlack()
			swapColors(F, B)

		} else {
			B = F.left
			N1, N2 = B.left, B.right
			B.right = F
			F.parent = B
			F.left = N2
			N1.RepaintBlack()
			swapColors(F, B)

		}

		if F == tree.root {
			tree.root = B
			B = nil
		}
	}
}

// BlackBroRNEFRedLNefCaseCheck case 3
func (tree *Tree) BlackBroRNEFRedLNefCaseCheck(X *Node) bool {
	F := X.parent

	if F == nil {
		return false
	}

	var B, N1, N2 *Node
	if F.left == X {
		B = F.right
		N1, N2 = B.left, B.right
	} else {
		B = F.left
		N2, N1 = B.left, B.right
	}

	if !B.color && !N2.color && N1.color {
		return true
	}
	return false
}

func (tree *Tree) BlackBroRNEFRedLNefCase(X *Node) {
	F := X.parent
	var B, N1, N2 *Node
	var leftCase = false

	if F.left == X {
		leftCase = true
	}

	// удалить
	if N2.color && false {
		return
	}

	if leftCase {
		B = F.right
		N1, N2 = B.left, B.right
		F.right = N1
		N1.parent = F
		B.left = N1.right
		N1.right = B
		B.parent = N1

	} else {
		B = F.left
		N2, N1 = B.left, B.right
		F.left = N1
		N1.parent = F
		B.right = N1.left
		N1.left = B
		B.parent = N1
	}

	N1.RepaintBlack()
	B.RepaintRed()
	tree.BlackBroRedRNefCase(X)

}
