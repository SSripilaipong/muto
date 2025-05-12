package builder

import (
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

type variableEmbedded struct {
	builder             ruleMutator.Builder
	variableMappings    []parameter.VariableMapping
	variadicVarMappings []parameter.VariadicVarMapping
}

func (v variableEmbedded) Build(parameter *parameter.Parameter) optional.Of[base.Node] {
	overriddenParameter, ok := v.overrideParameter(parameter).Return()
	if !ok {
		return optional.Empty[base.Node]()
	}
	return v.builder.Build(overriddenParameter)
}

func (v variableEmbedded) overrideParameter(p *parameter.Parameter) optional.Of[*parameter.Parameter] {
	q, ok := p.WithVariableMappings(v.variableMappings).Return()
	if !ok {
		return optional.Empty[*parameter.Parameter]()
	}
	return q.WithVariadicVarMappings(v.variadicVarMappings)
}

func withVariablesEmbedded(
	variableMappings []parameter.VariableMapping,
	variadicVarMappings []parameter.VariadicVarMapping,
	builder ruleMutator.Builder,
) variableEmbedded {
	return variableEmbedded{
		builder:             builder,
		variableMappings:    slices.Clone(variableMappings),
		variadicVarMappings: slices.Clone(variadicVarMappings),
	}
}

var _ ruleMutator.Builder = variableEmbedded{}
