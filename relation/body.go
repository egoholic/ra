package relation

import "github.com/egoholic/ra/relation/tuple"

type Source interface {
	Push(*tuple.Tuple)
	Pull() []*tuple.Tuple
	Reset()
}

type Body struct {
	source Source
	loaded []*tuple.Tuple
	added  []*tuple.Tuple
}

type BodyIterator struct {
	body Body
}

func NewBody(source Source) *Body {
	return &Body{source, nil, nil}
}

func (b *Body) Add(t *tuple.Tuple) {
	b.added = append(b.added, t)
}

func (b *Body) MakeIterator() *BodyIterator {

}