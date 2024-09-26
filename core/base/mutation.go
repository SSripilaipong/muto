package base

import "github.com/SSripilaipong/muto/common/optional"

type Mutation interface {
	Active(name string, obj Object) optional.Of[Node]
	Normal(name string, obj Object) optional.Of[Node]
}
