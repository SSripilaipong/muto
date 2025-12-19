package server

import "github.com/SSripilaipong/muto/core/base"

func errorMessage(node base.Node) string {
	obj := base.UnsafeNodeToObject(node)
	params := obj.ParamChain().DirectParams()
	if len(params) != 1 || !base.IsStringNode(params[0]) {
		return "http-server: handler error"
	}
	return base.UnsafeNodeToString(params[0]).Value()
}
