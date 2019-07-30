package relation

import (
	"errors"
	"fmt"

	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/relation/tuple"
	"github.com/egoholic/ra/val"
)

type Header struct {
	Attributes map[string]*attr.Attr
	PrimaryKey []string
}

func NewHeader(pk []string, attrs map[string]*attr.Attr) (header *Header, err error) {
	ok := checkPK(pk, attrs)
	if !ok {
		return nil, errors.New("wrong index")
	}
	return &Header{attrs, pk}, nil
}

func checkPK(pk []string, attrs map[string]*attr.Attr) bool {
	results := map[string]bool{}
	for _, pkAttrName := range pk {
		for attrName, _ := range attrs {
			if attrName == pkAttrName {
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

func (h *Header) MakeTuple(rawValues map[string]interface{}) (tpl *tuple.Tuple, err error) {
	if len(rawValues) != len(h.Attributes) {
		return nil, fmt.Errorf("wrong number of attributes\n\texpected: %d\n\tgiven: %d", len(h.Attributes), len(rawValues))
	}
	values := map[string]*val.V{}
	for attrName, attr := range h.Attributes {
		rawVal, ok := rawValues[attrName]
		if !ok {
			return nil, fmt.Errorf("attribute `%s` is not found", attrName)
		}
		val, err := attr.MakeVal(rawVal)
		if err != nil {
			return nil, err
		}
		values[attrName] = val
	}
	return tuple.New(values), nil
}
