package space

import (
	"fmt"

	"github.com/egoholic/ra/relation"
)

type Space struct {
	persistentRelations map[string]*relation.Relation
}

func New() *Space {
	return &Space{map[string]*relation.Relation{}}
}

func (s *Space) Find(rn string) (rel *relation.Relation, err error) {
	var ok bool
	rel, ok = s.persistentRelations[rn]
	if !ok {
		return nil, fmt.Errorf("relation `%s` not found", rn)
	}
	return
}

func (s *Space) Add(n string, r *relation.Relation) error {
	_, found := s.persistentRelations[n]
	if found {
		return fmt.Errorf("relation `%s` already exists", n)
	}
	s.persistentRelations[n] = r
	return nil
}
