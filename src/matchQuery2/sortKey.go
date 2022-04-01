package matchQuery2

type SortKey struct {
	sizeOfInvertedList int
	tokenArr           []string
}

func (s *SortKey) SizeOfInvertedList() int {
	return s.sizeOfInvertedList
}

func (s *SortKey) SetSizeOfInvertedList(sizeOfInvertedList int) {
	s.sizeOfInvertedList = sizeOfInvertedList
}

func (s *SortKey) TokenArr() []string {
	return s.tokenArr
}

func (s *SortKey) SetTokenArr(tokenArr []string) {
	s.tokenArr = tokenArr
}

func NewSortKey(pos int, tokenArr []string) SortKey {
	return SortKey{
		sizeOfInvertedList: pos,
		tokenArr:           tokenArr,
	}
}
