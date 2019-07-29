package relation

import (
	attr "github.com/egoholic/ra/attribute"
)

type Relation struct {
	Header
	Body
}

func New(pk []string, attrs ...*attr.Attr) *Relation {
}

type Projection struct {
}
