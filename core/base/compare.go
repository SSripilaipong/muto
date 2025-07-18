package base

import (
	"github.com/SSripilaipong/muto/core/base/datatype"
)

func NodeNotEqual(x, y Node) bool {
	return !NodeEqual(x, y)
}

func NodeEqual(x, y Node) bool {
	switch {
	case IsNumberNode(x):
		return numberEqual(UnsafeNodeToNumber(x), y)
	case IsStringNode(x):
		return stringEqual(UnsafeNodeToString(x), y)
	case IsTagNode(x):
		return tagEqual(UnsafeNodeToTag(x), y)
	case IsClassNode(x):
		return classEqual(UnsafeNodeToClass(x), y)
	case IsObjectNode(x):
		return objectEqual(UnsafeNodeToObject(x), y)
	case IsBooleanNode(x):
		return booleanEqual(UnsafeNodeToBoolean(x), y)
	case IsRuneNode(x):
		return runeEqual(UnsafeNodeToRune(x), y)
	}
	return false
}

func tagEqual(x Tag, y Node) bool {
	return IsTagNode(y) && x.Name() == UnsafeNodeToTag(y).Name()
}

func booleanEqual(x Boolean, y Node) bool {
	return IsBooleanNode(y) && x.Value() == UnsafeNodeToBoolean(y).Value()
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

func classEqual(x *Class, y Node) bool {
	if !IsClassNode(y) {
		return false
	}
	return x.Equals(UnsafeNodeToClass(y))
}

func runeEqual(x Rune, y Node) bool {
	if !IsRuneNode(y) {
		return false
	}
	return x.Equals(UnsafeNodeToRune(y))
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
