package base

func NewErrorWithMessage(msg string) Object {
	return NewNamedOneLayerObject("error", NewString(msg))
}
