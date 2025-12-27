package server

import (
	"net/http"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/portal"
)

type ModuleMountable interface {
	MountModule(module.Module)
}

type Server struct {
	requestClass base.Class
	okClass      base.Class
	errorClass   base.Class
}

func New() *Server {
	return &Server{}
}

func (s *Server) Call(nodes []base.Node) optional.Of[base.Node] {
	if s.requestClass == nil {
		return optional.Value[base.Node](base.NewErrorWithMessage("http-server: request class not available"))
	}
	return base.StrictStructureUnaryNodesOp(func(config base.Structure) optional.Of[base.Node] {
		server, err := buildHTTPServer(config, s.requestClass)
		if err != nil {
			return optional.Value[base.Node](base.NewOneLayerObject(s.errorClass, base.NewString(err.Error())))
		}

		go func() { _ = server.ListenAndServe() }()

		return optional.Value[base.Node](base.NewOneLayerObject(s.okClass, newController(server)))
	})(nodes)
}

func (s *Server) MountModule(mod module.Module) {
	s.requestClass = mod.GetClass("request")
	s.okClass = mod.GetClass("ok")
	s.errorClass = mod.GetClass("error")
}

func buildHTTPServer(config base.Structure, requestClass base.Class) (*http.Server, error) {
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
		Handler:      newHandler(handlerNode, requestClass),
	}, nil
}

var _ portal.Port = (*Server)(nil)
