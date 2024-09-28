package base

func NewErrorWithMessage(msg string) Object {
	return NewNamedObject("error", []Node{NewString(msg)})
}
