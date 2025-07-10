package extractor

type displayable interface {
	DisplayString() string
}

const nonDisplayableString = "<non-displayable>"

func DisplayString(x any) string {
	if displayVar, isDisplayable := x.(displayable); isDisplayable {
		return displayVar.DisplayString()
	}
	return nonDisplayableString
}
