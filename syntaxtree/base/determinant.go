package base

type Determinant interface {
	Pattern
	DeterminantType() DeterminantType
	DeterminantName() string
}

type DeterminantType string

const (
	DeterminantTypeObject      DeterminantType = "OBJECT"
	DeterminantTypeClass       DeterminantType = "CLASS"
	DeterminantTypeConjunction DeterminantType = "CONJUNCTION"
)

func ToDeterminant[T Determinant](x T) Determinant { return x }

func UnsafePatternToDeterminant(x Pattern) Determinant {
	return x.(Determinant)
}
