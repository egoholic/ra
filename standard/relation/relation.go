package relation

import (
	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/standard/relation/tuple"
)

type Header struct {
	Attributes map[string]*attr.Attr
	Metadata   map[string]interface{}
	PrimaryKey []string
}
type Body struct {
	tuples []*tuple.Tuple
}
type Relation struct {
	Header
	Body
}

func New(pk []string, attrs... *attr.Attr) *Relation{
	for _, attrName := range pk {
    for _, attr := range attrs {
      if attr.Name() == attrName {
        next()
			}
		}
	}
}

func

func (b *Body Insert(tuple *Tuple) (ok bool, messages []string) {

	b.tuples = append(r.Body, tuple)
}
