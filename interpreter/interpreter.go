package interpreter

import (
	"github.com/egoholic/ra/relation"
	"github.com/egoholic/ra/space"
)

type Interpreter struct {
	space *space.Space
}

func New(s *space.Space) *Interpreter {
	return &Interpreter{s}
}

type RelationGetter func() *relation.Relation
