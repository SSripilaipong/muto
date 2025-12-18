package portal

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type Spawner struct{}

func NewGoroutineSpawner() Spawner {
	return Spawner{}
}

func (s Spawner) Call(nodes []base.Node) optional.Of[base.Node] {
	if len(nodes) == 0 {
		return optional.Empty[base.Node]()
	}
	spawnNode := base.NewOneLayerObject(nodes[0], nodes[1:]...)
	go base.MutateUntilTerminated(spawnNode)
	return optional.Value[base.Node](base.Null())
}

var _ portal.Port = Spawner{}
