package projection

import (
	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/relation/tuple"
)

type SelectPattern interface {
	Copy(fromTuple *tuple.T, toTuple *tuple.T)
}

type SelectAttr struct {
	attrName attr.AttrName
}

func (p *SelectAttr) Copy(fromTuple, toTuple *tuple.T) {

}

type SelectOp struct {
	SelectPatterns []SelectPattern
}

func New(patterns ...SelectPattern) *SelectOp {
	return &SelectOp{patterns}
}
