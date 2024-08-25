package extractor

import (
	"phi-lang/core/base"
)

type ObjectLike interface {
	ClassName() string
	Children() []base.Node
}
