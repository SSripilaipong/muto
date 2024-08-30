package syntaxtree

type RuleResult interface {
	RuleResultType() RuleResultType
}

type RuleResultType string

const (
	RuleResultTypeString   RuleResultType = "STRING"
	RuleResultTypeNumber   RuleResultType = "NUMBER"
	RuleResultTypeObject   RuleResultType = "OBJECT"
	RuleResultTypeVariable RuleResultType = "VARIABLE"
)

type RuleResultObject struct {
	objectName string
	params     []ObjectParam
}

func NewRuleResultObject(objectName string, params []ObjectParam) RuleResultObject {
	return RuleResultObject{objectName: objectName, params: params}
}

func (RuleResultObject) RuleResultType() RuleResultType {
	return RuleResultTypeObject
}

func (obj RuleResultObject) ObjectName() string {
	return obj.objectName
}

func (obj RuleResultObject) Params() []ObjectParam {
	return obj.params
}

func IsRuleResultTypeString(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeString
}

func IsRuleResultTypeNumber(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeNumber
}

func IsRuleResultTypeObject(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeObject
}

func IsRuleResultTypeVariable(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeVariable
}

type ObjectParam interface {
	RuleResultType() RuleResultType
}

func ObjectParamToRuleResult(x ObjectParam) RuleResult {
	return x
}
