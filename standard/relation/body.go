package relation

import "github.com/egoholic/ra/standard/relation/tuple"

type Body struct {
	tuples []*tuple.Tuple
}

func (b *Body) Insert(tpl *tuple.Tuple) (ok bool, messages []string) {

	b.tuples = append(b.tuples, tpl)
}
