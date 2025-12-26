package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/core/base"
	corePortal "github.com/SSripilaipong/muto/core/portal"
)

type capturePort struct {
	called bool
	nodes  []base.Node
}

func (p *capturePort) Call(nodes []base.Node) optional.Of[base.Node] {
	p.called = true
	p.nodes = nodes
	return optional.Value[base.Node](base.Null())
}

func TestSpawn(t *testing.T) {
	mod := NewBaseModule()
	port := &capturePort{}
	mod.MountPortal(corePortal.New(rods.NewMap(map[string]corePortal.Port{
		"spawner": port,
	})))

	spawnClass := mod.GetClass("spawn!")
	obj := base.NewOneLayerObject(
		spawnClass,
		base.NewUnlinkedRuleBasedClass("a"),
		base.NewUnlinkedRuleBasedClass("b"),
		base.NewOneLayerObject(
			base.NewUnlinkedRuleBasedClass("c"),
			base.NewUnlinkedRuleBasedClass("d"),
		),
	)

	result := base.MutateUntilTerminated(obj)
	assert.Equal(t, base.Null(), result)
	assert.True(t, port.called)
	assert.Equal(t, []base.Node{
		base.NewUnlinkedRuleBasedClass("a"),
		base.NewUnlinkedRuleBasedClass("b"),
		base.NewOneLayerObject(
			base.NewUnlinkedRuleBasedClass("c"),
			base.NewUnlinkedRuleBasedClass("d"),
		),
	}, port.nodes)
}
