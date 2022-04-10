package index07

type IndexTree struct {
	qmin int
	qmax int
	cout int
	root *IndexTreeNode
}

func (i *IndexTree) Qmin() int {
	return i.qmin
}

func (i *IndexTree) SetQmin(qmin int) {
	i.qmin = qmin
}

func (i *IndexTree) Qmax() int {
	return i.qmax
}

func (i *IndexTree) SetQmax(qmax int) {
	i.qmax = qmax
}

func (i *IndexTree) Cout() int {
	return i.cout
}

func (i *IndexTree) SetCout(cout int) {
	i.cout = cout
}

func (i *IndexTree) Root() *IndexTreeNode {
	return i.root
}

func (i *IndexTree) SetRoot(root *IndexTreeNode) {
	i.root = root
}

//初始化trieTree
func NewIndexTree(qmin int, qmax int) *IndexTree {
	return &IndexTree{
		qmin: qmin,
		qmax: qmax,
		cout: 0,
		root: NewIndexTreeNode(""),
	}
}

func (tree *IndexTree) InsertIntoIndexTree(token []string, sid SeriesId, position int) *IndexTreeNode {
	node := tree.root
	var childIndex = -1
	var addr *IndexTreeNode
	for i, str := range token {
		childIndex = GetIndexNode(node.children, token[i])
		if childIndex == -1 {
			currentNode := NewIndexTreeNode(str)
			node.children = append(node.children, currentNode)
			node = currentNode
		} else {
			node = node.children[childIndex]
			node.frequency++
		}
		if i == len(token)-1 {
			node.isleaf = true
			if _, ok := node.invertedIndex[sid]; !ok {
				node.InsertSidAndPosArrToInvertedIndexMap(sid, position)
			} else {
				node.InsertPosArrToInvertedIndexMap(sid, position)
			}
			addr = node
		}
	}
	return addr
}

func (tree *IndexTree) InsertOnlyGramIntoIndexTree(tokenSubs []SubTokenOffset, addr *IndexTreeNode) {
	var childIndex = -1
	for k := 0; k < len(tokenSubs); k++ {
		token := tokenSubs[k].subToken
		offset := tokenSubs[k].offset
		node := tree.root
		for i, str := range token {
			childIndex = GetIndexNode(node.children, token[i])
			if childIndex == -1 {
				currentNode := NewIndexTreeNode(str)
				node.children = append(node.children, currentNode)
				node = currentNode
			} else {
				node = node.children[childIndex]
				node.frequency++
			}
			if i == len(token)-1 {
				node.isleaf = true
				if _, ok := node.addrOffset[addr]; !ok {
					node.addrOffset[addr] = offset
				}
			}
		}
	}
}

func (tree *IndexTree) PrintIndexTree() {
	tree.root.PrintIndexTreeNode(0)
}

//更新root节点的频率
func (tree *IndexTree) UpdateIndexRootFrequency() {
	for _, child := range tree.root.children {
		tree.root.frequency += child.frequency
	}
	tree.root.frequency--
}

var Res []int

func (root *IndexTreeNode) FixInvertedIndexSize() {
	for i := 0; i < len(root.children); i++ {
		if root.children[i].isleaf == true {
			Res = append(Res, len(root.children[i].invertedIndex))
		}
		root.children[i].FixInvertedIndexSize()
	}
}

var Grams [][]string
var temp []string

func (root *IndexTreeNode) SearchGramsFromIndexTree() {
	if root == nil {
		return
	}
	for i := 0; i < len(root.children); i++ {
		temp = append(temp, root.children[i].data)
		if root.children[i].isleaf == true {
			Grams = append(Grams, temp)
		}
		root.children[i].SearchGramsFromIndexTree()
		if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
	}
}
