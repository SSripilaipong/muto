package builtin

import (
	"muto/common/optional"
	"muto/common/slc"
	"muto/core/base"
	"muto/core/base/datatype"
	"muto/core/mutation/normal/object"
)

var addMutator = object.NewMutator("+", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.AddNumber(x.Value(), y.Value())))
})))

var subtractMutator = object.NewMutator("-", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.SubtractNumber(x.Value(), y.Value())))
})))

var multiplyMutator = object.NewMutator("*", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewNumber(datatype.MultiplyNumber(x.Value(), y.Value())))
})))

var divideMutator = object.NewMutator("/", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Map(base.NewNumber)(datatype.DivideNumber(x.Value(), y.Value()))
})))

var numberGreaterThanMutator = object.NewMutator(">", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.GreaterThanNumber(x.Value(), y.Value())))
})))

var numberGreaterThanOrEqualMutator = object.NewMutator(">=", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.GreaterThanOrEqualNumber(x.Value(), y.Value())))
})))

var numberLessThanMutator = object.NewMutator("<", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
	return optional.Value[base.Node](base.NewBoolean(datatype.LessThanNumber(x.Value(), y.Value())))
})))

var numberLessThanOrEqualMutator = object.NewMutator("<=", slc.Pure(numberBinaryOp(func(x, y base.Number) optional.Of[base.Node] {
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
