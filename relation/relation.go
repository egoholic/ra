package relation

import (
	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/relation/tuple"
)

type Relation struct {
	Meta
	Header
	Body
}

func New(pk []string, attrs map[string]*attr.Attr, source Source) *Relation {
	return &Relation{Meta{}, Header{attrs, pk}, *NewBody(source)}
}

func (r *Relation) Insert(rawAttrs map[string]interface{}) (tpl *tuple.Tuple, err error) {
	tpl, err = r.Header.MakeTuple(rawAttrs)
	if err != nil {
		return nil, err
	}
	r.Body.Add(tpl)
	return
}

func Project(r *Relation, attrNames []string) (rel *Relation, err error) {
  r.Body.
}

func Crossjoin() {

}

func Transform() {

}

func Group() *Relation {

}

func Sort() *Relation {

}
