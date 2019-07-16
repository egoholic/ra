package attribute

const (
	MaxAttrNameLen = 16
	MaxTypNameLen  = 16
)

type attrName [MaxAttrNameLen]byte
type typName [MaxTypNameLen]byte

// attribute
type ExtractTransformer interface {
	ExtractAndTransform(originalTuple, productTuple *T)
}

// attribute name
type AttrName struct {
	name attrName
}

func (a *AttrName) ExtractAndTransform(originalTuple, productTuple *T) {

}

type tpdVal struct {
	label typName
	value interface{}
}

// functional attribute
type AFn struct {
	name [MaxAttrNameLen]int
	fn   func(*T) *tpdVal
}
