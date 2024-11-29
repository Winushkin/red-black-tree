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

// rePaint

func (node *Node) RepaintBlack() {
	node.color = false //to black
}

func (node *Node) RepaintRed() {
	node.color = true // to red
}

func swapColors(a, b *Node) {
	a.color, b.color = b.color, a.color
}

func color(x *Node) bool {
	if x == nil {
		return false
	}
	return x.color
}
