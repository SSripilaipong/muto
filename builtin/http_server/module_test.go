package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/core/base"
	corePortal "github.com/SSripilaipong/muto/core/portal"
)

type httpServerCapturePort struct {
	called bool
	nodes  []base.Node
}

func (p *httpServerCapturePort) Call(nodes []base.Node) optional.Of[base.Node] {
	p.called = true
	p.nodes = nodes
	return optional.Value[base.Node](base.Null())
}

func TestHTTPServerStartUsesPortal(t *testing.T) {
	mod := NewModule()
	mod.LoadGlobal(global.NewModule())
	port := &httpServerCapturePort{}
	mod.MountPortal(corePortal.New(rods.NewMap(map[string]corePortal.Port{
		"http-server": port,
	})))

	class := mod.GetClass("start")
	config := base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.NewTag("handler"), base.NewUnlinkedRuleBasedClass("handler")),
	})

	result := base.MutateUntilTerminated(base.NewOneLayerObject(class, config))
	assert.Equal(t, base.Null(), result)
	assert.True(t, port.called)
	assert.Equal(t, []base.Node{config}, port.nodes)
}
