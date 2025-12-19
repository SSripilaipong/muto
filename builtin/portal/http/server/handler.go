package server

import (
	"net/http"

	"github.com/SSripilaipong/muto/core/base"
)

type handler struct {
	handler base.Node
}

func newHandler(node base.Node) http.Handler {
	return &handler{handler: node}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if recover() != nil {
			http.Error(w, "http-server handler panic", http.StatusInternalServerError)
		}
	}()

	reqObj, err := newRequestObject(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prog := base.NewOneLayerObject(h.handler, newHandleCommand(w, reqObj))
	result := base.MutateUntilTerminated(prog)
	if base.IsErrorNode(result) {
		http.Error(w, errorMessage(result), http.StatusInternalServerError)
	}
}

func newHandleCommand(writer http.ResponseWriter, req base.Node) base.Node {
	return base.NewOneLayerObject(base.NewTag("handle"), newResponseWriterClass(writer), req)
}
