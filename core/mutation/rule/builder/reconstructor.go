package builder

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	ruleExtractor "github.com/SSripilaipong/muto/core/mutation/rule/extractor"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/extractor"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type reconstructorBuilderFactory struct {
	node            nodeBuilderFactory
	classCollection ClassCollection
}

func newReconstructorBuilderFactory(nodeFactory nodeBuilderFactory, classCollection ClassCollection) reconstructorBuilderFactory {
	return reconstructorBuilderFactory{node: nodeFactory, classCollection: classCollection}
}

func (f reconstructorBuilderFactory) NewBuilder(recon stResult.Reconstructor) optional.Of[ruleMutator.Builder] { // TODO unit test
	isValidExtractor := newExtractorWithVariableFactory(recon.Extractor(), extractor.NewVariableFactory()).IsNotEmpty()
	if !isValidExtractor {
		return optional.Empty[ruleMutator.Builder]()
	}

	return optional.Value[ruleMutator.Builder](reconstructorBuilder{
		extractor: recon.Extractor(),
		builder:   NewObjectBuilderFactory(f.classCollection).NewBuilder(recon.Builder()),
	})
}

type reconstructorBuilder struct {
	extractor stPattern.ParamPart
	builder   ruleMutator.Builder
}

func (r reconstructorBuilder) Build(parameter *parameter.Parameter) optional.Of[base.Node] {
	variableFactory := extractor.NewEmbeddedVariableFactory(parameter.VariableMap(), parameter.VariadicVarMap())
	ext, isValidExtractor := newExtractorWithVariableFactory(r.extractor, variableFactory).Return()
	if !isValidExtractor {
		return optional.Empty[base.Node]()
	}

	embeddedBuilder := withVariablesEmbedded(parameter.VariableMappings(), parameter.VariadicVarMappings(), r.builder)
	return optional.Value[base.Node](NewReconstructor(ext, embeddedBuilder))
}

func newExtractorWithVariableFactory(pattern stPattern.ParamPart, variableFactory ruleExtractor.VariableFactory) optional.Of[extractor.NodeListExtractor] {
	extractorFactory := ruleExtractor.NewTopLevelFactory(variableFactory)
	return extractorFactory.TopLevel(pattern)
}

type Reconstructor struct {
	extractor extractor.NodeListExtractor
	builder   ruleMutator.Builder
}

func NewReconstructor(extractor extractor.NodeListExtractor, builder ruleMutator.Builder) Reconstructor {
	return Reconstructor{extractor: extractor, builder: builder}
}

func (Reconstructor) NodeType() base.NodeType { return base.NodeTypeReconstructor }

func (s Reconstructor) MutateAsHead(params base.ParamChain) optional.Of[base.Node] {
	newChildren := base.MutateParamChain(params)
	if newChildren.IsNotEmpty() {
		return optional.Value[base.Node](base.NewCompoundObject(s, newChildren.Value()))
	}

	build := optional.JoinFmap(s.builder.Build)
	appendRemainingParams := optional.JoinFmap(appendRemainingParamToNode(params.SliceFromOrEmpty(1)))

	return appendRemainingParams(build(s.extractor.Extract(params.DirectParams())))
}

func (s Reconstructor) TopLevelString() string {
	return s.String()
}

func (s Reconstructor) String() string {
	return "[WIP reconstructor]"
}

var _ base.Node = Reconstructor{}
