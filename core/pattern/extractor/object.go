package extractor

func NewObject(head NodeExtractor, children NodeListExtractor) NodeExtractor {
	if children == nil {
		return NewLeafObject(head)
	}
	return NewReducibleObject(head, children)
}
