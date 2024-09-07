package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	"muto/syntaxtree"
)

func objectParam(xs []tokenizer.Token) []tuple.Of2[syntaxtree.ObjectParam, []tokenizer.Token] {
	return ps.Or(
		objectParamValue,
		ps.Map(variableToObjectParam, variable),
		ps.Map(namedObjectWithoutChildToObjectParam, objectName),
		ps.Map(namedObjectToObjectParam, ps.Sequence3(openParenthesis, namedObject, closeParenthesis)),
		ps.Map(anonymousObjectToObjectParam, ps.Sequence3(openParenthesis, anonymousObject, closeParenthesis)),
	)(xs)
}

var objectParamValue = ps.Or(
	ps.Map(stringToObjectParam, string_),
	ps.Map(numberToObjectParam, number),
)

var anonymousObjectToObjectParam = tuple.Fn3(func(_ tokenizer.Token, obj anonymousObjectNode, _ tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewRuleResultAnonymousObject(obj.Head(), obj.Params())
})

func namedObjectWithoutChildToObjectParam(objName tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewRuleResultNamedObject(objName.Value(), nil)
}

var namedObjectToObjectParam = tuple.Fn3(func(_ tokenizer.Token, obj namedObjectNode, _ tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewRuleResultNamedObject(obj.Name(), obj.Params())
})

func numberToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewNumber(x.Value())
}

func stringToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	s := x.Value()
	return syntaxtree.NewString(s[1 : len(s)-1])
}

func variableToObjectParam(x tokenizer.Token) syntaxtree.ObjectParam {
	return syntaxtree.NewVariable(x.Value())
}
