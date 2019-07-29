package relation

import (
	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/standard/relation/tuple"
)

type Header struct {
	Attributes map[string]*attr.Attr
	PrimaryKey []string
}

func (h *Header) checkPK(pk []string, attrs []*attr.Attr) bool {
	results := map[string]bool{}
	for _, pkAttrName := range pk {
		for _, attr := range attrs {
			if attr.Name() == pkAttrName {
				results[pkAttrName] = true
				continue
			}
		}
		ok, _ := results[pkAttrName]
		if !ok {
			return false
		}
	}
	return true
}

// Checks tuple structure to be sure about that it could be added to the relation.
func (h *Header) checkTuple(tpl *tuple.Tuple) bool {
	if len(tpl.Values) != len(h.Attributes) {
		return false
	}

	for attrName, attr := range h.Attributes {
		val, ok := tpl.Values[attrName]
		if !ok {
			return false
		}
		ok, _ = attr.Check(val)
		if !ok {
			return false
		}
	}
	return true
}
