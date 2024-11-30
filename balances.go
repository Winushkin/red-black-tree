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

// BlackBroNephewsCaseCheck Case 1.(1,2)
func (tree *Tree) BlackBroNefCaseCheck(X, F *Node) bool {
	var B, N1, N2 *Node

	if F == nil {
		if X != nil {
			F = X.parent
			if F == nil {
				return false
			}
		}
	}

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

func (tree *Tree) BlackBroNefCase(X, F *Node) {

	if F == nil {
		F = X.parent
	}

	var B *Node

	if F.left == X {
		B = F.right
	} else {
		B = F.left
	}

	B.RepaintRed()
	if !F.color {
		F.RepaintBlack()
		if tree.BlackBroNefCaseCheck(F, nil) {
			tree.BlackBroNefCase(F, nil)
		}

	}

	F.RepaintBlack()
}

// BlackBroRedRNefCaseCheck Case 2
func (tree *Tree) BlackBroRedRNefCaseCheck(X, F *Node) bool {
	if F == nil {
		if X != nil {
			F = X.parent
			if F == nil {
				return false
			}
		}
	}

	var B, N2 *Node

	if F.left == X {
		B = F.right
		if B == nil {
			return false
		}
		N2 = B.right

	} else {
		B = F.left
		if B == nil {
			return false
		}
		N2 = B.left
	}

	if !color(B) && color(N2) {
		return true
	}
	return false
}

func (tree *Tree) BlackBroRedRNefCase(X, F *Node) {

	if F == nil {
		F = X.parent
	}

	var B, N1, N2 *Node
	var leftCase = false

	if F.left == X {
		leftCase = true
	}

	if leftCase {
		B = F.right
		N1, N2 = B.left, B.right

		B.left = F
		B.parent = F.parent
		F.right = N1
		N1.parent = F

		N2.RepaintBlack()
		swapColors(F, B)

	} else {
		B = F.left
		N1, N2 = B.right, B.left

		B.right = F
		F.parent = B
		F.left = N1
		N1.parent = F

		N2.RepaintBlack()
		swapColors(F, B)
	}

	F.parent = B
	if B.parent == nil {
		tree.root = B
	}
	if B == B.parent.left {
		B.parent.left = B
	} else {
		B.parent.right = B
	}

	//if F == tree.root {
	//	tree.root = B
	//	B = nil
	//}
}

// BlackBroRNEFRedLNefCaseCheck case 3
func (tree *Tree) BlackBroRNEFRedLNefCaseCheck(X, F *Node) bool {

	if F == nil {
		if X != nil {
			F = X.parent
			if F == nil {
				return false
			}
		}
	}

	var B, N1, N2 *Node
	if F.left == X {
		B = F.right
		if B == nil {
			return false
		}
		N1, N2 = B.left, B.right
	} else {
		B = F.left
		if B == nil {
			return false
		}
		N2, N1 = B.left, B.right
	}

	if !color(B) && !color(N2) && color(N1) {
		return true
	}
	return false
}

func (tree *Tree) BlackBroRNEFRedLNefCase(X, F *Node) {
	if F == nil {
		F = X.parent
	}

	var B, N1 *Node
	var leftCase = false

	if F.left == X {
		leftCase = true
	}

	if leftCase {
		B = F.right
		N1 = B.left
		F.right = N1
		N1.parent = F
		B.left = N1.right
		if N1.right != nil {
			N1.right.parent = B
		}
		N1.right = B
		B.parent = N1

	} else {
		B = F.left
		N1 = B.right
		F.left = N1
		N1.parent = F
		B.right = N1.left
		if N1.left != nil {
			N1.left.parent = B
		}
		N1.left = B
		B.parent = N1
	}

	N1.RepaintBlack()
	B.RepaintRed()

	if tree.BlackBroRedRNefCaseCheck(X, F) {
		tree.BlackBroRedRNefCase(X, F)
	}
}

// RedBroCaseCheck Case 4
func (tree *Tree) RedBroCaseCheck(X, F *Node) bool {

	if F == nil {
		if X != nil {
			F = X.parent
			if F == nil {
				return false
			}
		}
	}

	var B, N1, N2 *Node

	if X == F.left {
		B = F.right
		if B == nil {
			return false
		}
		N1, N2 = B.left, B.right

	} else {
		B = F.left
		if B == nil {
			return false
		}
		N1, N2 = B.right, B.left
	}

	if !color(F) && color(B) && !color(N1) && !color(N2) {
		return true
	}

	return false
}

func (tree *Tree) RedBroCase(X, F *Node) {
	if F == nil {
		F = X.parent
	}
	var B, N1, N2 *Node
	var leftCase = false

	if X == F.left {
		leftCase = true
	}

	if leftCase {
		B = F.right
		N1, N2 = B.left, B.right
		B.left = F
		B.parent = F.parent
		F.right = N1
		N1.parent = F

	} else {
		B = F.left
		N1, N2 = B.right, B.left
		B.right = F
		B.parent = F.parent
		F.left = N1
		N1.parent = F
	}

	F.parent = B
	if B.parent == nil {
		tree.root = B
	}
	if B == B.parent.left {
		B.parent.left = B
	} else {
		B.parent.right = B
	}

	F.RepaintRed()
	B.RepaintBlack()

	if tree.BlackBroNefCaseCheck(N2, F) {
		tree.BlackBroNefCase(N2, F)
	}

	//if tree.BlackBroRedRNefCase()

	//if F == tree.root {
	//	tree.root = B
	//	B = nil
	//}
}
