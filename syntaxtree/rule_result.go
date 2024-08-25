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

func IsRuleResultTypeString(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeString
}

func IsRuleResultTypeNumber(r RuleResult) bool {
	return r.RuleResultType() == RuleResultTypeNumber
}

type ObjectParam interface{}
