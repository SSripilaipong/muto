package server

import (
	"net/http"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type Server struct{}

func New() Server {
	return Server{}
}

func (Server) Call(nodes []base.Node) optional.Of[base.Node] {
	return base.StrictStructureUnaryOp(func(config base.Structure) optional.Of[base.Node] {
		server, err := buildHTTPServer(config)
		if err != nil {
			return optional.Value[base.Node](base.NewErrorWithMessage(err.Error()))
		}

		go func() { _ = server.ListenAndServe() }()

		return optional.Value[base.Node](newController(server))
	})(base.NewParamChain([][]base.Node{nodes}))
}

func buildHTTPServer(config base.Structure) (*http.Server, error) {
	handlerNode, err := parseHandler(config)
	if err != nil {
		return nil, err
	}

	addr, err := parseAddr(config)
	if err != nil {
		return nil, err
	}

	readTimeout, err := parseTimeout(config, "read-timeout")
	if err != nil {
		return nil, err
	}
	writeTimeout, err := parseTimeout(config, "write-timeout")
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:         addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      newHandler(handlerNode),
	}, nil
}

var _ portal.Port = Server{}
