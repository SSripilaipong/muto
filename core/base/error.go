package base

const errorClassName = "error"

func NewErrorWithMessage(msg string) Object {
	return NewNamedOneLayerObject(errorClassName, NewString(msg))
}

func IsErrorNode(node Node) bool {
	if !IsObjectNode(node) {
		return false
	}
	obj := UnsafeNodeToObject(node)
	head := obj.Head()
	if !IsClassNode(head) {
		return false
	}
	return UnsafeNodeToClass(head).Name() == errorClassName
}
