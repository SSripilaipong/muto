package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

var andMutator = NewRuleBasedMutatorFromFunctions("&", slc.Pure(booleanBinaryOp(func(x, y base.Boolean) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x.Value() && y.Value()))
})))

var orMutator = NewRuleBasedMutatorFromFunctions("|", slc.Pure(booleanBinaryOp(func(x, y base.Boolean) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(x.Value() || y.Value()))
})))

var notMutator = NewRuleBasedMutatorFromFunctions("!", slc.Pure(booleanUnaryOp(func(x base.Boolean) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(!x.Value()))
})))

func booleanBinaryOp(f func(x, y base.Boolean) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return binaryOp(func(x, y base.Node) optional.Of[base.Node] {
		if !base.IsBooleanNode(x) || !base.IsBooleanNode(y) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToBoolean(x), base.UnsafeNodeToBoolean(y))
	})
}

func booleanUnaryOp(f func(x base.Boolean) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return unaryOp(func(x base.Node) optional.Of[base.Node] {
		if !base.IsBooleanNode(x) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToBoolean(x))
	})
}
