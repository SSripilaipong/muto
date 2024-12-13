package builder

import (
	"fmt"

	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree/base"
)

func newStringBuilder(s st.String) constantBuilder[base.String] {
	var value string
	_, err := fmt.Sscanf(s.Value(), "%q", &value)
	if err != nil {
		panic(err)
	}
	return newConstantBuilder(base.NewString(value))
}
