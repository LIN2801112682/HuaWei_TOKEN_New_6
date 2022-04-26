package dictionary

import (
	"fmt"
	"sort"
)

type TrieTreeNode struct {
	data      string
	frequency int
	children  []*TrieTreeNode
	isleaf    bool
}

func (t *TrieTreeNode) Data() string {
	return t.data
}

func (t *TrieTreeNode) SetData(data string) {
	t.data = data
}

func (t *TrieTreeNode) Frequency() int {
	return t.frequency
}

func (t *TrieTreeNode) SetFrequency(frequency int) {
	t.frequency = frequency
}

func (t *TrieTreeNode) Children() []*TrieTreeNode {
	return t.children
}

func (t *TrieTreeNode) SetChildren(children []*TrieTreeNode) {
	t.children = children
}

func (t *TrieTreeNode) Isleaf() bool {
	return t.isleaf
}

func (t *TrieTreeNode) SetIsleaf(isleaf bool) {
	t.isleaf = isleaf
}

func NewTrieTreeNode(data string) *TrieTreeNode {
	return &TrieTreeNode{
		data:      data,
		frequency: 1,
		isleaf:    false,
		children:  make([]*TrieTreeNode, 0),
	}
}

//剪枝
func (node *TrieTreeNode) PruneNode(T int) {
	if !node.isleaf {
		for _, child := range node.children {
			child.PruneNode(T)
		}
	} else {
		if node.frequency <= T {
			node.PruneStrategyLessT()
		} else {
			node.PruneStrategyMoreT(T)
		}
	}
}

//剪枝策略<=T
func (node *TrieTreeNode) PruneStrategyLessT() {
	node.children = make([]*TrieTreeNode, 0)
}

//剪枝策略>T
//剪掉最大子集，若无法剪枝则递归剪子树
func (node *TrieTreeNode) PruneStrategyMoreT(T int) {
	var freqList = make([]FreqList, len(node.children))
	k := 0
	for _, child := range node.children {
		freqList[k].token = child.data
		freqList[k].freq = child.frequency
		k++
	}
	sort.SliceStable(freqList, func(i, j int) bool {
		if freqList[i].freq < freqList[j].freq {
			return true
		}
		return false
	})
	totoalSum := 0
	for i := k - 1; i >= 0; i-- {
		//从大到小遍历数组
		if totoalSum+freqList[i].freq <= T {
			totoalSum = totoalSum + freqList[i].freq
			for _, child := range node.children {
				if child.data == freqList[i].token && child.frequency == freqList[i].freq {
					node.children = append(node.children[:i], node.children[i+1:]...)
				}
			}
		}
	}
	// 不存在最大子集
	for _, child := range node.children {
		child.PruneNode(T)
	}
}

//判断children有无此节点
func getNode(children []*TrieTreeNode, str string) int {
	for i, child := range children {
		if child.data == str {
			return i
		}
	}
	return -1
}

//输出以node为根的子树
func (node *TrieTreeNode) PrintTreeNode(level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf)
	for _, child := range node.children {
		child.PrintTreeNode(level + 1)
	}
}
