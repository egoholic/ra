package val

type V struct {
	Val interface{}
	Typ []byte
}

func New(v interface{}, t []byte) *V {
	return &V{v, t}
}
