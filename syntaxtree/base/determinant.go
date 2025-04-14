package base

type Determinant interface {
	Pattern
	DeterminantType() DeterminantType
}

type DeterminantType string

const (
	DeterminantTypeObject DeterminantType = "OBJECT"
	DeterminantTypeClass  DeterminantType = "CLASS"
)

func ToDeterminant[T Determinant](x T) Determinant { return x }
