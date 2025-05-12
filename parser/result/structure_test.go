package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestStructure(t *testing.T) {
	t.Run("should parse empty structure", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse empty structure with white spaces", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
	 }abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse empty structure with white spaces", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
	 }abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse string key with no comma", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
  "xxx": 123
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewString(`"xxx"`), syntaxtree.NewNumber("123")),
		})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse string key with comma", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
  "xxx": 123,
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewString(`"xxx"`), syntaxtree.NewNumber("123")),
		})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse class key", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
 f: ""
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewClass("f"), syntaxtree.NewString(`""`)),
		})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse object value", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
 "": (g 555)
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewString(`""`), stResult.NewObject(syntaxtree.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("555")}))),
		})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should parse multiple records", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
1: 2, 
 3:4, 
5: 6
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(syntaxtree.NewNumber("1"), syntaxtree.NewNumber("2")),
			stResult.NewStructureRecord(syntaxtree.NewNumber("3"), syntaxtree.NewNumber("4")),
			stResult.NewStructureRecord(syntaxtree.NewNumber("5"), syntaxtree.NewNumber("6")),
		})
		expectedRemainder := psBase.IgnoreLineAndColumn(psBase.StringToCharTokens("abc"))
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(psBase.IgnoreLineAndColumnInResult(result)))
	})

	t.Run("should not parse records with duplicate key", func(t *testing.T) {
		assert.Empty(t, structure(psBase.StringToCharTokens(`{1: 2,1:4, }abc`)))
	})
}
