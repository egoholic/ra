package attribute

import (
	"errors"
	"reflect"

	"github.com/egoholic/ra/val"
)

type Checker func(interface{}) bool
type Attr struct {
	checker Checker
}

func New(c Checker) *Attr {
	return &Attr{c}
}

func (a *Attr) MakeVal(raw interface{}) (*val.V, error) {
	if !a.check(raw) {
		return nil, errors.New("wrong value")
	}
	return val.New(raw, []byte(reflect.TypeOf(raw).Name())), nil
}

func (a *Attr) check(raw interface{}) bool {
	return a.checker(raw)
}
