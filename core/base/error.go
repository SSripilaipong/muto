package base

func NewErrorWithMessage(msg string) Object {
	return NewNamedOneLayerObject("error", []Node{NewString(msg)})
}
