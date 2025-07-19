package portal

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rods"
)

type Portal struct {
	ports rods.Map[string, Port]
}

func New(ports rods.Map[string, Port]) *Portal {
	return &Portal{ports: ports}
}

func (p *Portal) Port(key string) optional.Of[Port] {
	return p.ports.Get(key)
}
