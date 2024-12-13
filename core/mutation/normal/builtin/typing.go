package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

var isNumberMutator = NewRuleBasedMutatorFromFunctions("number?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsNumberNode(x)))
})))

var isStringMutator = NewRuleBasedMutatorFromFunctions("string?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsStringNode(x)))
})))

var isBooleanMutator = NewRuleBasedMutatorFromFunctions("boolean?", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsBooleanNode(x)))
})))
