package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/base/datatype"
)

var addMutator = NewRuleBasedMutatorFromFunctions("+", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.AddNumber(x.Value(), y.Value())))
})))

var subtractMutator = NewRuleBasedMutatorFromFunctions("-", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.SubtractNumber(x.Value(), y.Value())))
})))

var multiplyMutator = NewRuleBasedMutatorFromFunctions("*", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.MultiplyNumber(x.Value(), y.Value())))
})))

var divideMutator = NewRuleBasedMutatorFromFunctions("/", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Map(base.NewNumber)(datatype.DivideNumber(x.Value(), y.Value()))
})))

var modIntegerMutator = NewRuleBasedMutatorFromFunctions("mod", slc.Pure(integerBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Map(base.NewNumber)(datatype.ModInteger(x.Value(), y.Value()))
})))

var divIntegerMutator = NewRuleBasedMutatorFromFunctions("div", slc.Pure(integerBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Map(base.NewNumber)(datatype.DivInteger(x.Value(), y.Value()))
})))

var numberGreaterThanMutator = NewRuleBasedMutatorFromFunctions(">", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.GreaterThanNumber(x.Value(), y.Value())))
})))

var numberGreaterThanOrEqualMutator = NewRuleBasedMutatorFromFunctions(">=", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.GreaterThanOrEqualNumber(x.Value(), y.Value())))
})))

var numberLessThanMutator = NewRuleBasedMutatorFromFunctions("<", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.LessThanNumber(x.Value(), y.Value())))
})))

var numberLessThanOrEqualMutator = NewRuleBasedMutatorFromFunctions("<=", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.LessThanOrEqualNumber(x.Value(), y.Value())))
})))

func numberBinaryOp(f func(x, y base.Number) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return binaryOp(func(x, y base.Node) optional.Of[base.Node] {
		if !base.IsNumberNode(x) || !base.IsNumberNode(y) {
			return optional.Empty[base.Node]()
		}
		return f(base.UnsafeNodeToNumber(x), base.UnsafeNodeToNumber(y))
	})
}

func integerBinaryOp(f func(x, y base.Number) optional.Of[base.Node]) func(t base.Object) optional.Of[base.Node] {
	return numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
		if !x.Value().IsInt() || !y.Value().IsInt() {
			return optional.Empty[base.Node]()
		}
		return f(x, y)
	})
}
