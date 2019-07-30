package relation

import (
	"errors"

	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/relation/tuple"
)

type Relation struct {
	Header
	Body
}

func New(pk []string, attrs map[string]*attr.Attr) *Relation {
	return &Relation{Header{attrs, pk}, Body{}}
}

func (r *Relation) Populate(tuples []*tuple.Tuple) error {
	if r.Body.tuples != nil {
		return errors.New("can't populate not empty relation")
	}
	r.Body.tuples = tuples
	return nil
}

func (r *Relation) Add(tuples []*tuple.Tuple) {
	for _, t := range tuples {
		r.Body.tuples = append(r.Body.tuples, t)
	}
}
