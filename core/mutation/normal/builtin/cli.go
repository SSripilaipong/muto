package builtin

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/normal/object"
)

func cliPrintMutator(print func(string)) object.RuleBasedMutator {
	return object.NewRuleBasedMutator("print!", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
		if base.IsStringNode(x) {
			print(base.UnsafeNodeToString(x).Value())
			return optional.Value[base.Node](base.NewClass("$"))
		}
		return optional.Empty[base.Node]()
	})))
}

func cliInputMutator(read func() rslt.Of[string]) object.RuleBasedMutator {
	return object.NewRuleBasedMutator("input!", slc.Pure(unaryOp(func(x base.Node) optional.Of[base.Node] {
		if base.IsClassNode(x) && base.UnsafeNodeToClass(x).Name() == "$" {
			s, err := read().Return()
			if err != nil {
				return optional.Value[base.Node](base.NewErrorWithMessage(err.Error()))
			}
			return optional.Value[base.Node](base.NewString(s))
		}
		return optional.Empty[base.Node]()
	})))
}
