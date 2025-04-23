package result

import stResult "github.com/SSripilaipong/muto/syntaxtree/result"

type objectNode struct {
	head      stResult.Node
	paramPart stResult.FixedParamPart
}

func (n objectNode) ParamPart() stResult.FixedParamPart {
	return n.paramPart
}

func (n objectNode) Head() stResult.Node {
	return n.head
}
