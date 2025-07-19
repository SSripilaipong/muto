package global

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

var stringToRunesMutator = NewRuleBasedMutatorFromFunctions("string-to-runes", slc.Pure(stringStrictUnaryOp(func(x string) optional.Of[base.Node] {
	var runes []base.Node
	for _, r := range []rune(x) {
		runes = append(runes, base.NewRune(r))
	}
	return optional.Value[base.Node](base.NewConventionalList(runes...))
})))

var concatMutator = NewRuleBasedMutatorFromFunctions("++", slc.Pure(stringStrictBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewString(x + y))
})))

var stringGreaterThanMutator = NewRuleBasedMutatorFromFunctions(">", slc.Pure(stringStrictBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x > y))
})))

var stringGreaterThanOrEqualMutator = NewRuleBasedMutatorFromFunctions(">=", slc.Pure(stringStrictBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x >= y))
})))

var stringLessThanMutator = NewRuleBasedMutatorFromFunctions("<", slc.Pure(stringStrictBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x < y))
})))

var stringLessThanOrEqualMutator = NewRuleBasedMutatorFromFunctions("<=", slc.Pure(stringStrictBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x <= y))
})))

var stringMutator = NewRuleBasedMutatorFromFunctions("string", slc.Pure(strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
	s, isStringer := x.(base.MutoStringer)
	if !isStringer {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](base.NewString(s.MutoString()))
})))

func stringStrictUnaryOp(f func(x string) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return strictUnaryOp(func(x base.Node) optional.Of[base.Node] {
		if !base.IsStringNode(x) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToString(x).Value())
	})
}

func stringStrictBinaryOp(f func(x, y string) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return strictBinaryOp(func(x, y base.Node) optional.Of[base.Node] {
		if !base.IsStringNode(x) || !base.IsStringNode(y) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToString(x).Value(), base.UnsafeNodeToString(y).Value())
	})
}
