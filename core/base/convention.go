package base

var GetTag = UnsafeNodeToTag(NewTag("get"))
var SetTag = UnsafeNodeToTag(NewTag("set"))
var ValueTag = UnsafeNodeToTag(NewTag("value"))
var EmptyTag = UnsafeNodeToTag(NewTag("empty"))

var ResultTag = UnsafeNodeToTag(NewTag("result"))
var RemainderTag = UnsafeNodeToTag(NewTag("remainder"))

func NewConventionalList(nodes []Node) Object {
	return NewNamedOneLayerObject("$", nodes)
}

func Null() Class {
	return NewUnlinkedRuleBasedClass("$")
}
