package index07

type SubTokenOffset struct {
	subToken []string
	offset   int
}

func (s *SubTokenOffset) SubToken() []string {
	return s.subToken
}

func (s *SubTokenOffset) SetSubToken(subToken []string) {
	s.subToken = subToken
}

func (s *SubTokenOffset) Offset() int {
	return s.offset
}

func (s *SubTokenOffset) SetOffset(offset int) {
	s.offset = offset
}

func NewSubGramOffset(subToken []string, offset int) SubTokenOffset {
	return SubTokenOffset{
		subToken: subToken,
		offset:   offset,
	}
}
