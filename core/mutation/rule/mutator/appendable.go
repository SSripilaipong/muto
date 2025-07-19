package mutator

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type Appendable struct {
	name   string
	active Switch
	normal Switch
}

func NewAppendable(name string, normal, active Switch) Appendable {
	return Appendable{name: name, normal: normal, active: active}
}

func (m Appendable) Active(obj base.Object) optional.Of[base.Node] {
	return m.active.Mutate(obj)
}

func (m Appendable) Normal(obj base.Object) optional.Of[base.Node] {
	return m.normal.Mutate(obj)
}

func (m Appendable) AppendNormal(u Unit) Appendable {
	return NewAppendable(m.name, m.normal.Append(u), m.active)
}

func (m Appendable) AppendActive(u Unit) Appendable {
	return NewAppendable(m.name, m.normal, m.active.Append(u))
}

func (m Appendable) ConcatNormal(sw Switch) Appendable {
	return NewAppendable(m.name, m.normal.Concat(sw), m.active)
}

func (m Appendable) ConcatActive(sw Switch) Appendable {
	return NewAppendable(m.name, m.normal, m.active.Concat(sw))
}

func (m Appendable) ActiveSwitch() Switch {
	return m.active
}

func (m Appendable) VisitClass(visitor ClassVisitor) {
	m.normal.VisitClass(visitor)
	m.active.VisitClass(visitor)
}

func (m Appendable) MountPortal(p *portal.Portal) {
	portal.MountPortalToMutator(p, m.normal)
	portal.MountPortalToMutator(p, m.active)
}
