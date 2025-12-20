package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestBuildHTTPServer(t *testing.T) {
	handlerNode := base.NewRuleBasedClass("app", nil)
	requestClass := base.NewRuleBasedClass("request", nil)
	config := base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.NewTag("handler"), handlerNode),
		base.NewStructureRecord(base.NewTag("addr"), base.NewString("127.0.0.1:8080")),
		base.NewStructureRecord(base.NewTag("read-timeout"), base.NewNumberFromString("2")),
		base.NewStructureRecord(base.NewTag("write-timeout"), base.NewNumberFromString("3")),
	})

	server, err := buildHTTPServer(config, requestClass)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1:8080", server.Addr)
	assert.Equal(t, 2*time.Second, server.ReadTimeout)
	assert.Equal(t, 3*time.Second, server.WriteTimeout)
	assert.Equal(t, handlerNode, server.Handler.(*handler).handler)
}
