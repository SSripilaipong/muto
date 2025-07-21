package global

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

var isNumberMutator = NewRuleBasedMutatorFromFunctions("number?", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsNumberNode(x)))
})))

var isStringMutator = NewRuleBasedMutatorFromFunctions("string?", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsStringNode(x)))
})))

var isBooleanMutator = NewRuleBasedMutatorFromFunctions("boolean?", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(base.IsBooleanNode(x)))
})))
