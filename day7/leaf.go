package main

import (
	"fmt"

	"github.com/lukegriffith/aoc22/tree"
)

type Leaf interface {
	Type() LeafType
	GetName() string
	GetSize() int
	GetPath() string
}

type File struct {
	Size   int
	Name   string
	Parent *tree.Node[Leaf]
}

func (f File) Type() LeafType {
	return FileType
}

func (f File) GetName() string {
	return f.Name
}

func (f File) GetSize() int {
	return f.Size
}

func (f File) GetPath() string {
	return fmt.Sprintf("%s/%s", f.Parent.Data.GetPath(), f.Name)
}

func NewFile(name string, size int, parent *tree.Node[Leaf]) tree.Node[Leaf] {
	f := File{size, name, parent}
	n := tree.Node[Leaf]{
		Data:     f,
		Children: nil,
		Parent:   parent,
	}
	parent.AddChild(&n)
	return n
}

type Dir struct {
	Name string
	Node *tree.Node[Leaf]
}

func (d Dir) Type() LeafType {
	return DirType
}

func (d Dir) GetName() string {
	return d.Name
}

func (d Dir) GetSize() int {
	size := 0
	dirNode := *d.Node
	for _, n := range dirNode.Children {
		size = size + n.Data.GetSize()
	}
	return size
}

func (d Dir) GetPath() string {
	if d.Node.Parent == nil {
		return d.Name
	} else {
		return fmt.Sprintf("%s/%s", d.Node.Parent.Data.GetPath(), d.Name)
	}
}

func NewDir(name string, parent *tree.Node[Leaf]) *tree.Node[Leaf] {
	directory := Dir{
		Name: name,
		Node: nil,
	}
	newNode := tree.Node[Leaf]{
		Tag:      name,
		Data:     nil,
		Children: []*tree.Node[Leaf]{},
		Parent:   parent,
	}
	newNode.Data = &directory
	directory.Node = &newNode
	parent.AddChild(&newNode)
	return &newNode
}

type LeafType int
