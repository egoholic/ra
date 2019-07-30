package attribute

import (
	"reflect"

	"github.com/egoholic/ra/val"
)

type Checker func(interface{}) bool
type Attr struct {
	checker Checker
}

func NewAttribute(c Checker) *Attr {
	return &Attr{c}
}

func (a *Attr) MakeVal(raw interface{}) *val.V {
	a.check(raw)
	return val.New(raw, []byte(reflect.TypeOf(raw).Name()))
}

func (a *Attr) check(raw interface{}) bool {
	return a.checker(raw)
}
