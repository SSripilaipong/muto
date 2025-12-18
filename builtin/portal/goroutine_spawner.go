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

func (s Spawner) Call(node base.Node) optional.Of[base.Node] {
	if base.IsMutableNode(node) {
		go mutateUntilTerminated(node)
	}
	return optional.Value[base.Node](base.Null())
}

func mutateUntilTerminated(node base.Node) {
	for base.IsMutableNode(node) {
		next, ok := base.UnsafeNodeToMutable(node).Mutate().Return()
		if !ok {
			return
		}
		node = next
	}
}

var _ portal.Port = Spawner{}
