package portal

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type LocalChannel struct{}

func NewLocalChannel() LocalChannel {
	return LocalChannel{}
}

func (LocalChannel) Call(nodes []base.Node) optional.Of[base.Node] {
	if len(nodes) != 1 {
		return optional.Empty[base.Node]()
	}
	x := nodes[0]
	if !base.IsClassNode(x) || base.UnsafeNodeToClass(x).Name() != "$" {
		return optional.Empty[base.Node]()
	}

	ch := make(chan base.Node)
	sender := base.NewRuleBasedClass("<sender>", senderRule{ch: ch})
	receiver := base.NewRuleBasedClass("<receiver>", receiverRule{ch: ch})

	return optional.Value[base.Node](base.NewConventionalList(sender, receiver))
}

type senderRule struct {
	base.NoActiveRule
	ch chan base.Node
}

func (r senderRule) Normal(obj base.Object) optional.Of[base.Node] {
	params := obj.ParamChain()
	children := params.DirectParams()
	if len(children) != 1 {
		return optional.Empty[base.Node]()
	}
	value := children[0]
	go func() {
		r.ch <- value
	}()
	remaining := params.SliceFromNodeOrEmpty(0, 1)
	return base.ProcessMutationResultWithParams(optional.Value[base.Node](base.Null()), remaining)
}

type receiverRule struct {
	base.NoActiveRule
	ch chan base.Node
}

func (r receiverRule) Normal(obj base.Object) optional.Of[base.Node] {
	if obj.ParamChain().TotalNodes() != 0 {
		return optional.Empty[base.Node]()
	}
	value := <-r.ch
	return optional.Value[base.Node](value)
}

var _ portal.Port = LocalChannel{}
