package pattern

import (
	"slices"

	"github.com/SSripilaipong/muto/syntaxtree/base"
)

func ExtractNonObjectHead(p base.Pattern) base.Pattern {
	for base.IsPatternTypeObject(p) {
		p = unwrapToHead(p)
	}
	return p
}

func ExtractParamChain(p base.Pattern) []ParamPart { // TODO unit test
	var paramChain []ParamPart // inserted in reversed order first, corrected before return

	for base.IsPatternTypeObject(p) {
		// Skip DeterminantConjunction - it delegates to main and doesn't add its own level
		if IsPatternDeterminantConjunction(p) {
			p = UnsafePatternToDeterminantConjunction(p).Main()
			continue
		}
		obj := UnsafePatternToObject(p)
		paramChain = append(paramChain, obj.ParamPart())
		p = obj.Head()
	}

	slices.Reverse(paramChain)
	return paramChain
}

// ExtractHeadConjunctions returns conjunction patterns at each nesting level.
// Each level has []Pattern (list of conjs at that level, empty if none).
// Levels are ordered from innermost to outermost.
// The conjunction at level i extracts from reconstruction level i.
func ExtractHeadConjunctions(p base.Pattern) [][]base.Pattern {
	numLevels := countDeterminantObjectLayers(p)
	if numLevels == 0 {
		return nil
	}
	result := make([][]base.Pattern, numLevels)
	for base.IsPatternTypeObject(p) {
		if IsPatternDeterminantConjunction(p) {
			dc := UnsafePatternToDeterminantConjunction(p)
			dc.placeConjsInResult(result)
			p = dc.Main()
			continue
		}
		p = UnsafePatternToObject(p).Head()
	}
	return result
}

// countDeterminantObjectLayers counts the number of DeterminantObject levels in the pattern,
// skipping DeterminantConjunction wrappers.
func countDeterminantObjectLayers(p base.Pattern) int {
	count := 0
	for base.IsPatternTypeObject(p) {
		if IsPatternDeterminantConjunction(p) {
			p = UnsafePatternToDeterminantConjunction(p).Main()
			continue
		}
		count++
		p = UnsafePatternToObject(p).Head()
	}
	return count
}

// unwrapToHead extracts the head from an Object pattern, handling DeterminantConjunction specially.
func unwrapToHead(p base.Pattern) base.Pattern {
	if IsPatternDeterminantConjunction(p) {
		return UnsafePatternToDeterminantConjunction(p).Main()
	}
	return UnsafePatternToObject(p).Head()
}
