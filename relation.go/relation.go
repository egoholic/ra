package relation

type header struct {
	mapping map[string]string
}

type body struct {
	tuples []interface{}
}

// relation
type R struct {
	header
	body
}
