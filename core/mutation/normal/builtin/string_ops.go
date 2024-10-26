package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

var concatMutator = object.NewRuleBasedMutator("++", slc.Pure(stringBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewString(x + y))
})))

var stringGreaterThanMutator = object.NewRuleBasedMutator(">", slc.Pure(stringBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x > y))
})))

var stringGreaterThanOrEqualMutator = object.NewRuleBasedMutator(">=", slc.Pure(stringBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x >= y))
})))

var stringLessThanMutator = object.NewRuleBasedMutator("<", slc.Pure(stringBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x < y))
})))

var stringLessThanOrEqualMutator = object.NewRuleBasedMutator("<=", slc.Pure(stringBinaryOp(func(x, y string) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x <= y))
})))

var stringMutator = object.NewRuleBasedMutator("string", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
	s, isStringer := x.(base.MutoStringer)
	if !isStringer {
		return optional.Empty[base.Node]()
	}
	return optional.Value[base.Node](base.NewString(s.MutoString()))
})))

func stringBinaryOp(f func(x, y string) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return binaryOp(func(x, y base.Node) optional.Of[base.Node] {
		if !base.IsStringNode(x) || !base.IsStringNode(y) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToString(x).Value(), base.UnsafeNodeToString(y).Value())
	})
}
