package matchQuery

import "index07"

type SortTokenInvertList struct {
	tokenArr  []string
	indexList index07.Inverted_index
}

func (s *SortTokenInvertList) TokenArr() []string {
	return s.tokenArr
}

func (s *SortTokenInvertList) SetTokenArr(tokenArr []string) {
	s.tokenArr = tokenArr
}

func (s *SortTokenInvertList) IndexList() index07.Inverted_index {
	return s.indexList
}

func (s *SortTokenInvertList) SetIndexList(indexList index07.Inverted_index) {
	s.indexList = indexList
}

func NewSortTokenInvertList(tokenArr []string, indexList index07.Inverted_index) SortTokenInvertList {
	return SortTokenInvertList{
		tokenArr:  tokenArr,
		indexList: indexList,
	}
}
