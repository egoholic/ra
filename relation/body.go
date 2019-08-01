package relation

import "github.com/egoholic/ra/tuple"

type DataPuller interface {
	Pull() ([]*tuple.Tuple, error)
}

type Body struct {
	source DataPuller
}
