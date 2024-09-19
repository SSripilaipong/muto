package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
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
