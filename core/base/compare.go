package base

import (
	"muto/core/base/datatype"
)

func NodeEqual(x, y Node) bool {
	switch {
	case IsNumberNode(x):
		return numberEqual(UnsafeNodeToNumber(x), y)
	case IsStringNode(x):
		return stringEqual(UnsafeNodeToString(x), y)
	case IsObjectNode(x):
		return objectEqual(UnsafeNodeToObject(x), y)
	case IsBooleanNode(x): // will not be reached for now
		return objectEqual(UnsafeNodeToObject(x), y)
	}
	return false
}

func numberEqual(x Number, y Node) bool {
	return IsNumberNode(y) && datatype.EqualNumber(x.Value(), UnsafeNodeToNumber(y).Value())
}

func stringEqual(x String, y Node) bool {
	return IsStringNode(y) && x.Value() == UnsafeNodeToString(y).Value()
}

func objectEqual(x Object, y Node) bool {
	return (IsAnonymousObjectNode(x) && anonymousObjectEqual(UnsafeObjectToAnonymousObject(x), y)) ||
		(IsNamedObjectNode(x) && namedObjectEqual(UnsafeObjectToNamedObject(x), y))
}

func anonymousObjectEqual(x AnonymousObject, y Node) bool {
	if !IsAnonymousObjectNode(y) {
		return false
	}
	yObj := UnsafeNodeToAnonymousObject(y)
	if !NodeEqual(x.Head(), yObj.Head()) {
		return false
	}
	return objectChildrenEqual(x.Children(), yObj.Children())
}

func namedObjectEqual(x NamedObject, y Node) bool {
	if !IsNamedObjectNode(y) {
		return false
	}
	yObj := UnsafeNodeToNamedObject(y)
	if x.Name() != yObj.Name() {
		return false
	}
	return objectChildrenEqual(x.Children(), yObj.Children())
}

func objectChildrenEqual(xs []Node, ys []Node) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i, x := range xs {
		if !NodeEqual(x, ys[i]) {
			return false
		}
	}
	return true
}
