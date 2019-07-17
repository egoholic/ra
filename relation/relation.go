package relation

import "github.com/egoholic/ra/relation/tuple"

type Header struct {
	mapping map[string]string
}

type Body struct {
	tuples []*tuple.T
}

// relation
type R struct {
	Header
	Body
}
