package mem

import "github.com/egoholic/ra/relation/tuple"

type MemSource struct {
	tuples   []*tuple.Tuple
	pageSize int
	cursor   int
}

func (s *MemSource) Push(t *tuple.Tuple) {
	s.tuples = append(s.tuples, t)
}

func (s *MemSource) Pull() (tuples []*tuple.Tuple) {
	limit := s.cursor + s.pageSize
	for i := s.cursor; i < limit || i < len(s.tuples); i++ {
		tuples = append(tuples, s.tuples[i])
	}
	return
}

func (s *MemSource) Reset() {
	s.cursor = 0
}
