package portal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestLocalChannel(t *testing.T) {
	port := NewLocalChannel()
	result, ok := port.Call([]base.Node{base.Null()}).Return()
	if !assert.True(t, ok) {
		return
	}
	if !assert.True(t, base.IsObjectNode(result)) {
		return
	}
	obj := base.UnsafeNodeToObject(result)
	params := obj.ParamChain().DirectParams()
	if !assert.Len(t, params, 2) {
		return
	}
	if !assert.True(t, base.IsClassNode(params[0])) {
		return
	}
	if !assert.True(t, base.IsClassNode(params[1])) {
		return
	}
	sender := base.UnsafeNodeToClass(params[0])
	receiver := base.UnsafeNodeToClass(params[1])
	assert.Equal(t, "<sender>", sender.Name())
	assert.Equal(t, "<receiver>", receiver.Name())

	got := make(chan base.Node, 1)
	go func() {
		got <- mutateUntilTerminatedResult(base.NewOneLayerObject(receiver))
	}()

	select {
	case <-got:
		t.Fatal("receiver should block before send")
	case <-time.After(50 * time.Millisecond):
	}

	sent := base.NewString("hello")
	senderResult := mutateUntilTerminatedResult(base.NewOneLayerObject(sender, sent))
	assert.True(t, base.NodeEqual(senderResult, base.Null()))

	select {
	case received := <-got:
		assert.True(t, base.NodeEqual(received, sent))
	case <-time.After(200 * time.Millisecond):
		t.Fatal("timed out waiting for receiver")
	}
}

func mutateUntilTerminatedResult(node base.Node) base.Node {
	if !base.IsMutableNode(node) {
		return node
	}
	if result, isMutated := base.UnsafeNodeToMutable(node).Mutate().Return(); isMutated {
		return mutateUntilTerminatedResult(result)
	}
	return node
}
