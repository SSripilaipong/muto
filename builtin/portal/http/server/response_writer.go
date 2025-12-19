package server

import (
	"io"
	"net/http"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
)

type responseWriterRule struct {
	base.NoActiveRule
	writer      http.ResponseWriter
	status      int
	wroteHeader bool
}

func newResponseWriterClass(w http.ResponseWriter) *base.RuleBasedClass {
	return base.NewRuleBasedClass("<response-writer>", &responseWriterRule{
		writer: w,
		status: http.StatusOK,
	})
}

func (r *responseWriterRule) Normal(obj base.Object) optional.Of[base.Node] {
	return base.StrictObjectUnaryOp(func(cmd base.Object) optional.Of[base.Node] {
		head := cmd.Head()
		if !base.IsTagNode(head) {
			return optional.Empty[base.Node]()
		}

		return r.process(base.UnsafeNodeToTag(head), cmd.ParamChain())
	})(obj.ParamChain())
}

func (r *responseWriterRule) process(cmd base.Tag, params base.ParamChain) optional.Of[base.Node] {
	switch cmd.Name() {
	case "status":
		return r.setStatus(params)
	case "add-header":
		return r.addHeader(params)
	case "set-header":
		return r.setHeader(params)
	case "write-string":
		return r.writeString(params)
	}
	return optional.Empty[base.Node]()
}

func (r *responseWriterRule) setStatus(params base.ParamChain) optional.Of[base.Node] {
	return base.StrictNumberUnaryOp(func(num base.Number) optional.Of[base.Node] {
		r.status = int(num.Value().ToInt())
		if !r.wroteHeader {
			r.writer.WriteHeader(r.status)
			r.wroteHeader = true
		}
		return optional.Value[base.Node](base.Null())
	})(params)
}

func (r *responseWriterRule) addHeader(params base.ParamChain) optional.Of[base.Node] {
	return base.StrictStringBinaryOp(func(key, value base.String) optional.Of[base.Node] {
		r.writer.Header().Add(key.Value(), value.Value())
		return optional.Value[base.Node](base.Null())
	})(params)
}

func (r *responseWriterRule) setHeader(params base.ParamChain) optional.Of[base.Node] {
	return base.StrictStringBinaryOp(func(key, value base.String) optional.Of[base.Node] {
		r.writer.Header().Set(key.Value(), value.Value())
		return optional.Value[base.Node](base.Null())
	})(params)
}

func (r *responseWriterRule) writeString(params base.ParamChain) optional.Of[base.Node] {
	return base.StrictStringUnaryOp(func(body base.String) optional.Of[base.Node] {
		if !r.wroteHeader {
			r.writer.WriteHeader(r.status)
			r.wroteHeader = true
		}
		if _, err := io.WriteString(r.writer, body.Value()); err != nil {
			return optional.Value[base.Node](base.NewErrorWithMessage(err.Error()))
		}
		return optional.Value[base.Node](base.Null())
	})(params)
}
