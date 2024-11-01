package builder

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func buildString(s st.String) func(mapping *parameter.Parameter) optional.Of[base.Node] {
	var value string
	_, err := fmt.Sscanf(s.Value(), "%q", &value)
	if err != nil {
		panic(err)
	}
	return func(mapping *parameter.Parameter) optional.Of[base.Node] {
		return optional.Value[base.Node](base.NewString(value))
	}
}
