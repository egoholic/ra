package relation

import "github.com/egoholic/ra/relation/tuple"

type Source interface {
	Push(*tuple.Tuple) error
	Pull() ([]*tuple.Tuple, error)
	Reset()
}

type Body struct {
	source Source
	loaded []*tuple.Tuple
}

type BodyIterator struct {
	body *Body
	cur  int
}

func (bi *BodyIterator) Next() (tuple *tuple.Tuple) {
	loaded := bi.body.loaded
	len = len(loaded)
	if bi.cur < len {
		bi.cur = bi.cur + 1
		return bi.loaded[bi.cur]
	}

	bi.body.source.Pull()
	loaded = bi.body.loaded
	len = len(loaded)
	if bi.cur < len {
		bi.cur = bi.cur + 1
		return bi.loaded[bi.cur]
	}
	return nil
}

func NewBody(source Source) *Body {
	return &Body{source, nil, nil}
}

func (b *Body) Load() error {
	tuples := b.source.Pull()
	b.loaded = append(b.loaded, tuples)
}

func (b *Body) Insert(t *tuple.Tuple) (err error) {
	err = b.source.Push(t)
	if err != nil {
		return err
	}
	b.added = append(b.added, t)
}

func (b *Body) MakeIterator() *BodyIterator {
	return &BodyIterator{b}
}
