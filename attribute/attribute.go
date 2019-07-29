package attribute

import (
	"fmt"

	"github.com/egoholic/ra/val"
)

type Attr struct {
	name        string
	typ         []byte
	constraints []Constraint
}

func (a *Attr) Name() string {
	return a.name
}

func (a *Attr) MatchType(typ2 []byte) bool {
	if len(a.typ) != len(typ2) {
		return false
	}
	for i, b1 := range a.typ {
		if b1 != typ2[i] {
			return false
		}
	}
	return true
}

type Constraint interface {
	Check(*val.V) (bool, string)
}

type Condition interface {
	Check(*val.V) bool
}

func NewAttribute(name string, typ []byte, constraints []Constraint) *Attr {
	return &Attr{name, typ, constraints}
}

func (a *Attr) Check(v *val.V) (ok bool, messages []string) {
	var msg string
	ok = false
	if len(a.typ) != len(v.Typ) {
		msg = fmt.Sprintf("wrong type\n\tEXPECTED: %s\n\tGOT: %s", string(a.typ), string(v.Typ))
		messages = append(messages, msg)
		return
	}
	for _, constraint := range a.constraints {
		ok, msg = constraint.Check(v)
		if !ok {
			messages = append(messages, msg)
			return
		}
	}
	ok = true
	return
}
