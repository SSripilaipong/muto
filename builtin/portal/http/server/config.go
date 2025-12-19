package server

import (
	"errors"
	"time"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

func parseHandler(config base.Structure) (base.Node, error) {
	handlerNode, ok := structGet(config, base.NewTag("handler")).Return()
	if !ok {
		return nil, errors.New("http-server: missing .handler")
	}
	if !base.IsClassNode(handlerNode) && !base.IsObjectNode(handlerNode) {
		return nil, errors.New("http-server: .handler must be class or object")
	}
	return handlerNode, nil
}

func parseAddr(config base.Structure) (string, error) {
	if addrNode, ok := structGet(config, base.NewTag("addr")).Return(); ok {
		addr, ok := nodeToString(addrNode)
		if !ok {
			return "", errors.New("http-server: .addr must be string")
		}
		if addr == "" {
			return "", errors.New("http-server: .addr cannot be empty")
		}
		return addr, nil
	}

	portNode, portOk := structGet(config, base.NewTag("port")).Return()
	if !portOk {
		return "", errors.New("http-server: missing .addr or .port")
	}
	port, ok := nodeToString(portNode)
	if !ok {
		return "", errors.New("http-server: .port must be string or number")
	}
	host := ""
	if hostNode, ok := structGet(config, base.NewTag("host")).Return(); ok {
		host, ok = nodeToString(hostNode)
		if !ok {
			return "", errors.New("http-server: .host must be string")
		}
	}
	if host == "" {
		return ":" + port, nil
	}
	return host + ":" + port, nil
}

func parseTimeout(config base.Structure, key string) (time.Duration, error) {
	timeoutNode, ok := structGet(config, base.NewTag(key)).Return()
	if !ok {
		return 0, nil
	}
	if !base.IsNumberNode(timeoutNode) {
		return 0, errors.New("http-server: ." + key + " must be number")
	}
	n := base.UnsafeNodeToNumber(timeoutNode).Value()
	seconds := n.ToFloat()
	return time.Duration(seconds * float64(time.Second)), nil
}

func nodeToString(x base.Node) (string, bool) {
	switch {
	case base.IsStringNode(x):
		return base.UnsafeNodeToString(x).Value(), true
	case base.IsNumberNode(x):
		return base.UnsafeNodeToNumber(x).Value().SimpleString(), true
	default:
		return "", false
	}
}

func structGet(s base.Structure, key base.Node) optional.Of[base.Node] {
	getObj := base.NewOneLayerObject(base.GetTag, key)
	return s.MutateAsHead(base.NewParamChain([][]base.Node{{getObj}}))
}
