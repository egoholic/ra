package relation

import (
	attr "github.com/egoholic/ra/attribute"
)

type Relation struct {
	Meta
	Header
	Body
}

func New(pk []string, attrs map[string]*attr.Attr) *Relation {
	return &Relation{Meta{}, Header{attrs, pk}, Body{}}
}
