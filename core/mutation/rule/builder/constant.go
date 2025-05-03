package builder

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/base/datatype"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type constantBuilderFactory struct {
	class ClassCollection
}

func newConstantBuilderFactory(classCollection ClassCollection) constantBuilderFactory {
	return constantBuilderFactory{class: classCollection}
}

func (f constantBuilderFactory) NewBuilder(r stResult.Node) optional.Of[mutator.Builder] {
	switch {
	case stResult.IsNodeTypeBoolean(r):
		return optional.Value[mutator.Builder](newBooleanBuilder(stBase.UnsafeRuleResultToBoolean(r)))
	case stResult.IsNodeTypeString(r):
		return optional.Value[mutator.Builder](newStringBuilder(stBase.UnsafeRuleResultToString(r)))
	case stResult.IsNodeTypeNumber(r):
		return optional.Value[mutator.Builder](newNumberBuilder(stBase.UnsafeRuleResultToNumber(r)))
	case stResult.IsNodeTypeClass(r):
		return optional.Value[mutator.Builder](f.newClassBuilder(stBase.UnsafeRuleResultToClass(r)))
	case stResult.IsNodeTypeTag(r):
		return optional.Value[mutator.Builder](newTagBuilder(stBase.UnsafeRuleResultToTag(r)))
	}
	return optional.Empty[mutator.Builder]()
}

func (f constantBuilderFactory) newClassBuilder(x stBase.Class) constantWrapper[*base.Class] {
	return newConstantWrapper(f.class.GetClass(x.Name()))
}

func newBooleanBuilder(x stBase.Boolean) constantWrapper[base.Boolean] {
	return newConstantWrapper(base.NewBoolean(x.BooleanValue()))
}

func newNumberBuilder(x stBase.Number) constantWrapper[base.Node] {
	return newConstantWrapper(base.NewNumber(datatype.NewNumber(x.Value())))
}

func newStringBuilder(s stBase.String) constantWrapper[base.String] {
	var value string
	_, err := fmt.Sscanf(s.Value(), "%q", &value)
	if err != nil {
		panic(err)
	}
	return newConstantWrapper(base.NewString(value))
}

func newTagBuilder(x stBase.Tag) constantWrapper[base.Node] {
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
