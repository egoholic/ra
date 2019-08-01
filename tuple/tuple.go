package tuple

import (
	"github.com/egoholic/ra/val"
)

type Tuple struct {
	Values map[string]*val.V
}

func New(values map[string]*val.V) *Tuple {
	return &Tuple{values}
}
