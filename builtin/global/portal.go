package global

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
)

const portalMutatorName = "portal"

type portalMutator struct {
	portal *portal.Portal
}

func newPortalMutator() *portalMutator {
	return &portalMutator{}
}

func (t *portalMutator) Name() string { return portalMutatorName }

func (t *portalMutator) Mutate(obj base.Object) optional.Of[base.Node] {
	return objectStrictUnaryOp(func(x base.Object) optional.Of[base.Node] {
		head := x.Head()
		switch {
		case base.NodeEqual(head, base.NewTag("call")):
			return t.callCommand(x.ParamChain())
		}
		return optional.Empty[base.Node]()
	})(obj)
}

func (t *portalMutator) callCommand(param base.ParamChain) optional.Of[base.Node] {
	keyNode, valueNodes, ok := t.validateCallParam(param)
	if !ok {
		return optional.Empty[base.Node]()
	}
	return t.call(keyNode, valueNodes)
}

func (t *portalMutator) call(keyNode base.Node, valueNodes []base.Node) optional.Of[base.Node] {
	key := base.UnsafeNodeToString(keyNode).Value()
	if port, portExists := t.portal.Port(key).Return(); portExists {
		return port.Call(valueNodes)
	}
	return optional.Empty[base.Node]()
}

func (t *portalMutator) validateCallParam(param base.ParamChain) (base.Node, []base.Node, bool) {
	if param.Size() != 1 {
		return nil, nil, false
	}
	children := param.DirectParams()
	if len(children) < 2 {
		return nil, nil, false
	}
	keyNode := children[0]
	if !base.IsStringNode(keyNode) {
		return nil, nil, false
	}
	return keyNode, children[1:], true
}

func (t *portalMutator) MountPortal(p *portal.Portal) {
	t.portal = p
}

func (t *portalMutator) VisitClass(mutator.ClassVisitor) {}
