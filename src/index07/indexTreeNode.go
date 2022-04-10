package index07

import (
	"fmt"
)

type IndexTreeNode struct {
	data          string
	frequency     int
	children      []*IndexTreeNode
	isleaf        bool
	invertedIndex Inverted_index
	addrOffset    map[*IndexTreeNode]int
}

func (node *IndexTreeNode) Data() string {
	return node.data
}

func (node *IndexTreeNode) SetData(data string) {
	node.data = data
}

func (node *IndexTreeNode) Frequency() int {
	return node.frequency
}

func (node *IndexTreeNode) SetFrequency(frequency int) {
	node.frequency = frequency
}

func (node *IndexTreeNode) Children() []*IndexTreeNode {
	return node.children
}

func (node *IndexTreeNode) SetChildren(children []*IndexTreeNode) {
	node.children = children
}

func (node *IndexTreeNode) Isleaf() bool {
	return node.isleaf
}

func (node *IndexTreeNode) SetIsleaf(isleaf bool) {
	node.isleaf = isleaf
}

func (node *IndexTreeNode) InvertedIndex() Inverted_index {
	return node.invertedIndex
}

func (node *IndexTreeNode) SetInvertedIndex(invertedIndex Inverted_index) {
	node.invertedIndex = invertedIndex
}

func (node *IndexTreeNode) AddrOffset() map[*IndexTreeNode]int {
	return node.addrOffset
}

func (node *IndexTreeNode) SetAddrOffset(addrOffset map[*IndexTreeNode]int) {
	node.addrOffset = addrOffset
}

func NewIndexTreeNode(data string) *IndexTreeNode {
	return &IndexTreeNode{
		data:          data,
		frequency:     1,
		isleaf:        false,
		children:      make([]*IndexTreeNode, 0),
		invertedIndex: make(map[SeriesId][]int),
		addrOffset:    make(map[*IndexTreeNode]int),
	}
}

//判断children有无此节点
func GetIndexNode(children []*IndexTreeNode, str string) int {
	for i, child := range children {
		if child.data == str {
			return i
		}
	}
	return -1
}

func (node *IndexTreeNode) InsertPosArrToInvertedIndexMap(sid SeriesId, position int) {
	//倒排列表数组中找到sid的invertedIndex，把position加入到invertedIndex中的posArray中去
	node.invertedIndex[sid] = append(node.invertedIndex[sid], position)
}

//插入倒排
func (node *IndexTreeNode) InsertSidAndPosArrToInvertedIndexMap(sid SeriesId, position int) {
	posArray := []int{}
	posArray = append(posArray, position)
	node.invertedIndex[sid] = posArray
}

//输出以node为根的子树
func (node *IndexTreeNode) PrintIndexTreeNode(level int) {
	fmt.Println()
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Print(node.data, " - ", node.frequency, " - ", node.isleaf, " - ", node.invertedIndex, " - ", node.addrOffset)
	for _, child := range node.children {
		child.PrintIndexTreeNode(level + 1)
	}
}
