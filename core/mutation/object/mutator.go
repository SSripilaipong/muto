package object

import (
	"phi-lang/common/optional"
	"phi-lang/core/base"
	ruleMutationExtractor "phi-lang/core/mutation/rule/extractor"
)

type Mutator struct {
	name          string
	mutationRules []func(t ruleMutationExtractor.ObjectLike) optional.Of[base.Node]
}

func (t Mutator) Mutate(obj ruleMutationExtractor.ObjectLike) optional.Of[base.Node] {
	if t.name != obj.ClassName() {
		return optional.Empty[base.Node]()
	}
	for _, mutate := range t.mutationRules {
		if result := mutate(obj); !result.IsEmpty() {
			return result
		}
	}
	return optional.Empty[base.Node]()
}

func (t Mutator) Name() string {
	return t.name
}

func MutatorName(t Mutator) string {
	return t.Name()
}
