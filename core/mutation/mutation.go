package mutation

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

type mutation struct {
	active func(name string, obj base.Object) optional.Of[base.Node]
	normal func(name string, obj base.Object) optional.Of[base.Node]
}

func (m mutation) Active(name string, obj base.Object) optional.Of[base.Node] {
	return m.active(name, obj)
}

func (m mutation) Normal(name string, obj base.Object) optional.Of[base.Node] {
	return m.normal(name, obj)
}
