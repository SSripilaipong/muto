package module

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rods"

	"github.com/SSripilaipong/muto/core/base"
	mutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildUserDefinedModule(t *testing.T) {
	t.Run("should compute imported class", func(t *testing.T) {
		stTree := st.NewModule([]st.File{st.NewFile([]stBase.Statement{
			st.NewImport([]string{"my-module"}), // :import my-module
			st.NewRule(
				stPattern.NewDeterminantObject(st.NewLocalClass("main"), stPattern.FixedParamPart{}),
				stResult.NewNakedObject(st.NewImportedClass("my-module", "test"), stResult.FixedParamPart{}),
			), // main = my-module.test
		})})

		ruleBuilder := mutation.NewRuleBuilder()
		test := ruleBuilder.Build(st.NewRule(
			stPattern.NewDeterminantObject(stPattern.NewDeterminantObject(st.NewLocalClass("test"), stPattern.FixedParamPart{}), stPattern.FixedParamPart{}),
			stResult.NewNakedObject(st.NewString(`"ok"`), stResult.FixedParamPart{}),
		)) // (test) = "ok"

		module := BuildUserDefinedModuleFromBase(NewBase(mutator.NewCollectionFromMutators(nil, nil), ruleBuilder), stTree).
			Attach(NewDependency(nil, NewImportMapping(rods.NewMap(map[string]Module{
				"my-module": NewBase(mutator.NewCollectionFromMutators([]mutator.NamedUnit{test}, nil), ruleBuilder),
			}))))
		main := base.NewOneLayerObject(module.GetClass("main"))
		assert.Equal(t, base.NewString("ok"), main.Mutate().Value().(base.Object).Mutate().Value())
	})
}
