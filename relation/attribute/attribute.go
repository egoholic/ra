package attribute

const (
	MaxAttrNameLen = 16
	MaxTypNameLen  = 16
)

type AttrName [MaxAttrNameLen]byte
type TypName [MaxTypNameLen]byte

type TpdVal struct {
	label TypName
	value interface{}
}
