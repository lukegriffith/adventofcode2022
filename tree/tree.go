package tree

import "errors"

type Node[T any] struct {
	Tag      string
	Data     T
	Children []*Node[T]
	Parent   *Node[T]
}

func (n *Node[T]) AddChild(child *Node[T]) {
	if n == nil {
		return
	}
	n.Children = append(n.Children, child)
}

func (n *Node[T]) RootNode() *Node[T] {
	node := n
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}

func (n *Node[T]) FindChildByTag(tag string) (*Node[T], error) {
	var node *Node[T]
	for _, n := range n.Children {
		if n.Tag == tag {
			node = n
			break
		}
	}
	if node == nil {
		return nil, errors.New("unable to find node")
	}
	return node, nil
}

func (n *Node[T]) ExecuteDelegateOnAllChildren(delegate func(*Node[T])) {
	for _, n := range n.Children {
		n.ExecuteDelegateOnAllChildren(delegate)
	}
	delegate(n)
}
