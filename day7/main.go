package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukegriffith/aoc22"
	"github.com/lukegriffith/aoc22/tree"
)

const (
	FileType LeafType = iota
	DirType

	CommandLS = "ls"
	CommandCD = "cd"
)

func loadFileSystem(inputs []string) *tree.Node[Leaf] {
	node := NewDir("/", nil)
	for _, i := range inputs[1:] {
		ii := strings.Split(i, " ")
		pre := ii[0]
		post := ii[1]
		if pre[0] == '$' {
			switch post {
			case CommandCD:
				if ii[2][0] == '/' {
					// cd to root
					node = node.RootNode()
				} else if ii[2] == ".." {
					node = node.Parent
				} else {
					// cd to directory
					dirName := ii[2]
					directory, err := node.FindChildByTag(dirName)
					aoc22.CheckErr(err)
					node = directory
					if err != nil {
						panic(err)
					}
				}
			case CommandLS:
				continue
			}
		} else {

			pre := ii[0]
			if pre == "dir" {
				dirName := ii[1]
				NewDir(dirName, node)
			} else {
				// LS Output
				size, err := strconv.Atoi(ii[0])
				aoc22.CheckErr(err)
				name := ii[1]
				NewFile(name, size, node)
			}

		}
	}
	return node.RootNode()
}

func main() {
	inputs := aoc22.LoadInputs("input.txt")
	node := loadFileSystem(inputs)

	sizeMap := make(map[string]int)

	n := node.RootNode()
	sizeMap[n.Data.GetPath()] = n.Data.GetSize()

	countDelegate := func(n *tree.Node[Leaf]) {
		if n.Data.Type() == DirType {
			sizeMap[n.Data.GetPath()] = n.Data.GetSize()
		}
	}
	n.ExecuteDelegateOnAllChildren(countDelegate)

	total := 0

	for _, v := range sizeMap {
		if v <= 100000 {
			total = total + v
		}
	}

	fmt.Println(total)

	totalDiskSize := 70000000
	requiredFreeDiskSize := 30000000
	rootSize := n.Data.GetSize()
	freeSpace := totalDiskSize - rootSize
	needToFree := requiredFreeDiskSize - freeSpace

	for k, v := range sizeMap {
		if v > needToFree {
			fmt.Println(k, v)
		}
	}

}
