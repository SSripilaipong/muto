package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/base/datatype"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type constantBuilderFactory struct{}

func newConstantBuilderFactory() constantBuilderFactory {
	return constantBuilderFactory{}
}

func (f constantBuilderFactory) NewBuilder(r stResult.Node) optional.Of[mutator.Builder] {
	switch {
	case stResult.IsNodeTypeBoolean(r):
		return optional.Value[mutator.Builder](newBooleanBuilder(syntaxtree.UnsafeRuleResultToBoolean(r)))
	case stResult.IsNodeTypeString(r):
		return optional.Value[mutator.Builder](newStringBuilder(syntaxtree.UnsafeRuleResultToString(r)))
	case stResult.IsNodeTypeRune(r):
		return optional.Value[mutator.Builder](newRuneBuilder(syntaxtree.UnsafeRuleResultToRune(r)))
	case stResult.IsNodeTypeNumber(r):
		return optional.Value[mutator.Builder](newNumberBuilder(syntaxtree.UnsafeRuleResultToNumber(r)))
	case stResult.IsNodeTypeClass(r):
		return optional.Value[mutator.Builder](newClassBuilder(syntaxtree.UnsafeRuleResultToClass(r)))
	case stResult.IsNodeTypeTag(r):
		return optional.Value[mutator.Builder](newTagBuilder(syntaxtree.UnsafeRuleResultToTag(r)))
	}
	return optional.Empty[mutator.Builder]()
}

func newBooleanBuilder(x syntaxtree.Boolean) constantWrapper[base.Boolean] {
	return newConstantWrapper(base.NewBoolean(x.BooleanValue()))
}

func newNumberBuilder(x syntaxtree.Number) constantWrapper[base.Node] {
	return newConstantWrapper(base.NewNumber(datatype.NewNumber(x.Value())))
}

func newStringBuilder(s syntaxtree.String) constantWrapper[base.String] {
	return newConstantWrapper(base.NewString(s.StringValue()))
}

func newRuneBuilder(s syntaxtree.Rune) constantWrapper[base.Rune] {
	return newConstantWrapper(base.NewRune(s.RuneValue()))
}

func newTagBuilder(x syntaxtree.Tag) constantWrapper[base.Node] {
	return newConstantWrapper(base.NewTag(x.Name()))
}

type constantWrapper[T base.Node] struct {
	value T
}

func newConstantWrapper[T base.Node](value T) constantWrapper[T] {
	return constantWrapper[T]{value: value}
}

func (b constantWrapper[T]) Build(_ *parameter.Parameter) optional.Of[base.Node] {
	return optional.Value[base.Node](b.value)
}

func (b constantWrapper[T]) DisplayString() string {
	return b.value.TopLevelString()
}
