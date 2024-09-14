package base

import "muto/common/optional"

type Mutation interface {
	Active(name string, obj NamedObject) optional.Of[Node]
	Normal(name string, obj NamedObject) optional.Of[Node]
}
