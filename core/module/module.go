package module

import (
	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/portal"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Module interface {
	BuildNode(node stResult.SimplifiedNode) optional.Of[base.Node]
	GetClass(name string) base.Class
	MountPortal(q *portal.Portal)
	MutatorCollection() mutator.Collection
	MapImportedModules(mapping ImportMapping)
	ExtendCollection(collection mutator.Collection)
	ExtendImportedCollection(moduleName string, collection mutator.Collection)
}
