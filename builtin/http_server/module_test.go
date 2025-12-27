package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/core/base"
	corePortal "github.com/SSripilaipong/muto/core/portal"
)

type httpServerCallCapturePort struct {
	config []base.Node
}

func (p *httpServerCallCapturePort) Call(config []base.Node) optional.Of[base.Node] {
	p.config = config
	return optional.Value[base.Node](base.Null())
}

func TestHTTPServerStartCallsPortal(t *testing.T) {
	port := &httpServerCallCapturePort{}
	mod := NewModule()
	mod.MountPortal(corePortal.New(rods.NewMap(map[string]corePortal.Port{
		"http-server": port,
	})))

	config := base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.NewTag("handler"), base.NewUnlinkedRuleBasedClass("handler")),
	})

	result := base.MutateUntilTerminated(base.NewOneLayerObject(mod.GetClass("start"), config))
	assert.Equal(t, base.Null(), result)
	assert.Equal(t, []base.Node{config}, port.config)
}
