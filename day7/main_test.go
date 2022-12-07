package main

import (
	"fmt"
	"testing"

	"github.com/lukegriffith/aoc22/tree"
)

func TestGetSize(t *testing.T) {
	n := NewDir("/", nil)
	f1 := NewFile("a", 1, n)
	NewFile("b", 1, n)
	NewFile("c", 1, n)

	if f1.Data.GetSize() != 1 {
		t.Fail()
	}

	if n.Data.GetSize() != 3 {
		fmt.Println(n.Data.GetSize())
		t.Fail()
	}

	n1 := NewDir("a", n)
	NewFile("b", 1, n1)
	NewFile("c", 1, n1)
	if n1.Data.GetSize() != 2 {
		fmt.Println(n.Data.GetSize())
		t.Fail()
	}

	if n.Data.GetSize() != 5 {
		fmt.Println(n.Data.GetSize())
		t.Fail()
	}

}

func TestLoadFileSystem(t *testing.T) {
	i := []string{
		"$ cd /",
		"$ ls",
		"1 b.txt",
		"1 c.txt",
		"dir d",
		"$ cd d",
		"$ ls",
		"1 f.txt",
		"dir e",
		"$ cd e",
		"$ ls",
		"1 g.txt",
	}
	n := loadFileSystem(i)
	if n == nil {
		t.Fail()
	}
	sizeMap := make(map[string]int)
	countDelegate := func(n *tree.Node[Leaf]) {
		if n.Data.Type() == DirType {
			sizeMap[n.Data.GetPath()] = n.Data.GetSize()
		}
	}
	n.ExecuteDelegateOnAllChildren(countDelegate)
}
