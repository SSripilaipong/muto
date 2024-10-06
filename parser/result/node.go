package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Node = ps.Or(
	nonNestedNode,
	ps.Map(castObjectNode, ps.Filter(objectWithChildren, object)),
)

var nonNestedNode = ps.Or(
	ps.Map(castBooleanNode, psBase.Boolean),
	ps.Map(castStringNode, psBase.String),
	ps.Map(castNumberNode, psBase.Number),
	ps.Map(castClassNode, psBase.Class),
	ps.Map(castVariableNode, psBase.FixedVar),
)

func objectWithChildren(obj objectNode) bool {
	param := obj.ParamPart()
	switch {
	case stResult.IsParamPartTypeFixed(param):
		return stResult.UnsafeParamPartToFixedParamPart(param).Size() > 0
	}
	return false
}

func castClassNode(x tokenizer.Token) stResult.Node {
	return st.NewClass(x.Value())
}

func castNumberNode(x tokenizer.Token) stResult.Node {
	return st.NewNumber(x.Value())
}

func castBooleanNode(x tokenizer.Token) stResult.Node {
	return st.NewBoolean(x.Value())
}

func castStringNode(x tokenizer.Token) stResult.Node {
	return st.NewString(x.Value())
}

func castVariableNode(x tokenizer.Token) stResult.Node {
	return st.NewVariable(x.Value())
}

func castObjectNode(obj objectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}
