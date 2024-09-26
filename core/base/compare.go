package base

import (
	"github.com/SSripilaipong/muto/core/base/datatype"
)

func NodeEqual(x, y Node) bool {
	switch {
	case IsNumberNode(x):
		return numberEqual(UnsafeNodeToNumber(x), y)
	case IsStringNode(x):
		return stringEqual(UnsafeNodeToString(x), y)
	case IsNamedClassNode(x):
		return namedClassEqual(UnsafeNodeToNamedClass(x), y)
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
	if !IsObjectNode(y) {
		return false
	}
	return x.Equals(UnsafeNodeToObject(y))
}

func namedClassEqual(x Class, y Node) bool {
	if !IsNamedClassNode(y) {
		return false
	}
	return x.Name() == UnsafeNodeToNamedClass(y).Name()
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
