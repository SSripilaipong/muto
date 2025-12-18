package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/core/base"
	corePortal "github.com/SSripilaipong/muto/core/portal"
)

type channelCapturePort struct {
	called bool
	nodes  []base.Node
	result base.Node
}

func (p *channelCapturePort) Call(nodes []base.Node) optional.Of[base.Node] {
	p.called = true
	p.nodes = nodes
	return optional.Value[base.Node](p.result)
}

func TestNewChannelUsesPortal(t *testing.T) {
	mod := NewModule()
	expected := base.NewConventionalList(
		base.NewUnlinkedRuleBasedClass("sender"),
		base.NewUnlinkedRuleBasedClass("receiver"),
	)
	port := &channelCapturePort{result: expected}
	mod.MountPortal(corePortal.New(rods.NewMap(map[string]corePortal.Port{
		"chbroker": port,
	})))

	class := mod.GetClass("new-channel")
	result := base.MutateUntilTerminated(base.NewOneLayerObject(class))

	assert.True(t, port.called)
	assert.True(t, base.NodeEqual(result, expected))
	if assert.Len(t, port.nodes, 1) {
		assert.True(t, base.IsClassNode(port.nodes[0]))
		assert.Equal(t, "$", base.UnsafeNodeToClass(port.nodes[0]).Name())
	}
}
