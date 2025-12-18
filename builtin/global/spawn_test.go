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
	node   base.Node
}

func (p *capturePort) Call(node base.Node) optional.Of[base.Node] {
	p.called = true
	p.node = node
	return optional.Value[base.Node](base.Null())
}

func TestSpawn(t *testing.T) {
	mod := NewModule()
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

	result := mutateUntilTerminated(obj)
	assert.Equal(t, base.Null(), result)
	assert.True(t, port.called)
	assert.Equal(t, base.NewOneLayerObject(
		base.NewUnlinkedRuleBasedClass("a"),
		base.NewUnlinkedRuleBasedClass("b"),
		base.NewOneLayerObject(
			base.NewUnlinkedRuleBasedClass("c"),
			base.NewUnlinkedRuleBasedClass("d"),
		),
	), port.node)
}
