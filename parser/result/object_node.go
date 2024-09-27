package result

import stResult "github.com/SSripilaipong/muto/syntaxtree/result"

type objectNode struct {
	head      stResult.Node
	paramPart stResult.ParamPart
}

func (n objectNode) ParamPart() stResult.ParamPart {
	return n.paramPart
}

func (n objectNode) Head() stResult.Node {
	return n.head
}
