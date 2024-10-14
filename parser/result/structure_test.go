package result

import (
	"testing"

	"github.com/stretchr/testify/assert"

	psBase "github.com/SSripilaipong/muto/parser/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestStructure(t *testing.T) {
	t.Run("should parse empty structure", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse empty structure with white spaces", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
	 }abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse empty structure with white spaces", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
	 }abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse string key with no comma", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
  "xxx": 123
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewString(`"xxx"`), st.NewNumber("123")),
		})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse string key with comma", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
  "xxx": 123,
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewString(`"xxx"`), st.NewNumber("123")),
		})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse class key", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
 f: ""
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewClass("f"), st.NewString(`""`)),
		})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse object value", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
 "": (g 555)
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewString(`""`), stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("555")}))),
		})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})

	t.Run("should parse multiple records", func(t *testing.T) {
		result := structure(psBase.StringToCharTokens(`{	 
1: 2, 
 3:4, 
5: 6
}abc`))
		expectedResult := stResult.NewStructure([]stResult.StructureRecord{
			stResult.NewStructureRecord(st.NewNumber("1"), st.NewNumber("2")),
			stResult.NewStructureRecord(st.NewNumber("3"), st.NewNumber("4")),
			stResult.NewStructureRecord(st.NewNumber("5"), st.NewNumber("6")),
		})
		expectedRemainder := psBase.StringToCharTokens("abc")
		assert.Equal(t, psBase.SingleResult(expectedResult, expectedRemainder), psBase.AsParserResult(result))
	})
}
