package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

var isNumberMutator = object.NewRuleBasedMutator("number?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsNumberNode(x)))
})))

var isStringMutator = object.NewRuleBasedMutator("string?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsStringNode(x)))
})))

var isBooleanMutator = object.NewRuleBasedMutator("boolean?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsBooleanNode(x)))
})))
