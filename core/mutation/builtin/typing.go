package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/mutation/object"
)

var isNumberMutator = object.NewMutator("number?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsNumberNode(x)))
})))

var isStringMutator = object.NewMutator("string?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsStringNode(x)))
})))

var isBooleanMutator = object.NewMutator("boolean?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsBooleanNode(x)))
})))
