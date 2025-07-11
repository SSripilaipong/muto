package builder

type displayable interface {
	DisplayString() string
}

type nakedDisplayable interface {
	NakedDisplayString() string
}

const nonDisplayableString = "<non-displayable>"

func DisplayString(x any) string {
	if t, isDisplayable := x.(displayable); isDisplayable {
		return t.DisplayString()
	}
	return nonDisplayableString
}

func NakedDisplayString(x any) string {
	if t, isDisplayable := x.(nakedDisplayable); isDisplayable {
		return t.NakedDisplayString()
	}
	return DisplayString(x)
}
