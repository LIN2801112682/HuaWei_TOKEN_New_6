package index07

import (
	"bufio"
	"dictionary"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

//根据一批日志数据通过字典树划分VG，构建索引项集
func GenerateIndexTree(filename string, qmin int, qmax int, root *dictionary.TrieTreeNode) (*IndexTree, *IndexTreeNode) {
	indexTree := NewIndexTree(qmin, qmax)
	data, err := os.Open(filename)
	defer data.Close()
	if err != nil {
		fmt.Print(err)
	}
	buff := bufio.NewReader(data)
	var id int32 = 0
	var sum1 int64 = 0
	var sum2 int64 = 0
	var sum3 int64 = 0
	timeStamp := time.Now().Unix()
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		var vgMap map[int][]string
		vgMap = make(map[int][]string)
		id++
		timeStamp++
		sid := NewSeriesId(id, timeStamp)
		str := string(data)
		start1 := time.Now().UnixMicro()
		VGCons(root, qmin, qmax, str, vgMap)
		//.Println(vgMap)
		var keys = []int{}
		for key := range vgMap {
			keys = append(keys, key)
		}
		//对map中的key进行排序（map遍历是无序的）
		sort.Sort(sort.IntSlice(keys))
		end1 := time.Now().UnixMicro()
		sum1 += (end1 - start1)

		start2 := time.Now().UnixMicro()
		var addr *IndexTreeNode
		for i := 0; i < len(keys); i++ {
			vgKey := keys[i]
			//字符串变字符串数组
			tokenArr := vgMap[vgKey]
			start3 := time.Now().UnixMicro()
			//addr = nil
			addr = indexTree.InsertIntoIndexTree(tokenArr, sid, vgKey)
			end3 := time.Now().UnixMicro()
			sum3 += (end3 - start3)
			if len(tokenArr) > qmin && len(tokenArr) <= qmax { //Generate all index entries between qmin+1 - len(gram)
				TokenSubs = make([]SubTokenOffset, 0)
				GenerateQmin2QmaxTokens(tokenArr, qmin)
				start4 := time.Now().UnixMicro()
				indexTree.InsertOnlyGramIntoIndexTree(TokenSubs, addr)
				end4 := time.Now().UnixMicro()
				sum3 += (end4 - start4)
			}
		}
		end2 := time.Now().UnixMicro()
		sum2 += (end2 - start2)
	}
	indexTree.cout = (int(id))
	indexTree.UpdateIndexRootFrequency()
	fmt.Println("构建索引项集总花费时间（us）：", sum1+sum2)
	fmt.Println("读取日志并划分索引项花费时间（us）：", sum1+sum2-sum3)
	fmt.Println("插入索引树花费时间（us）：", sum3)
	//indexTree.PrintIndexTree()
	return indexTree, indexTree.root
}

var TokenSubs []SubTokenOffset

//func GenerateQmin2QmaxTokens(tokenArr []string, qmin int) {
//	length := len(tokenArr)
//	//get http 1.0 1.1  http 1.0 1.1 1.0 1.1
//	j := len(tokenArr[0]) + 1
//	for i := 1; i <= length-qmin; i++ {
//		tokenSub := tokenArr[i:length]
//		TokenSubs = append(TokenSubs, NewSubGramOffset(tokenSub, j))
//		j += len(tokenArr[i]) + 1
//	}
//}
func GenerateQmin2QmaxTokens(tokenArr []string, qmin int) {
	length := len(tokenArr)
	//get http 1.0 1.1  http 1.0 1.1 1.0 1.1
	for i := 1; i <= length-qmin; i++ {
		tokenSub := tokenArr[i:length]
		TokenSubs = append(TokenSubs, NewSubGramOffset(tokenSub, i))
	}
}

//根据字典D划分日志为VG
func VGCons(root *dictionary.TrieTreeNode, qmin int, qmax int, str string, vgMap map[int][]string) {
	tokenArray := strings.Fields(str)
	len1 := len(tokenArray)
	for p := 0; p < len1-qmin+1; p++ {
		tSub = tSub[0:0]
		FindLongestGramFromDic(root, tokenArray, p)
		t := tSub
		if len(t) == 0 || len(t) < qmin {
			t = tokenArray[p : p+qmin]
		}
		if !IsSubStrOfVG(t, vgMap) {
			for i := 0; i < len(t); i++ {
				vgMap[p] = append(vgMap[p], t[i])
			}
		}
	}
}

func IsSubStrOfVG(t []string, vgMap map[int][]string) bool {
	var flag = false
	var tstr = ""
	var strNew = ""
	for i := 0; i < len(t); i++ {
		tstr += t[i]
	}
	var keys = []int{}
	for key := range vgMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.IntSlice(keys))
	for i := 0; i < len(keys); i++ {
		vgKey := keys[i]
		str := vgMap[vgKey]
		for j := 0; j < len(str); j++ {
			strNew += str[j]
		}
		if strNew == tstr {
			flag = false
		} else if i == (len(keys)-1) && strings.Contains(strNew, tstr) {
			flag = true
		} else if i < (len(keys)-1) && strings.Contains(strNew, tstr) {
			flag = false
		}
		strNew = ""
	}
	return flag
}

var tSub []string

func FindLongestGramFromDic(root *dictionary.TrieTreeNode, str []string, p int) {
	if p < len(str) {
		c := str[p : p+1]
		s := strings.Join(c, "")
		for i := 0; i < len(root.Children()); i++ {
			if root.Children()[i].Data() == s {
				tSub = append(tSub, s)
				FindLongestGramFromDic(root.Children()[i], str, p+1)
			}
			if i == len(root.Children()) {
				return
			}
		}
	}
}
