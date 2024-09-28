package datatype

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
)

type Number struct {
	isFloat    bool
	intValue   int32
	floatValue float64
}

func NewNumber(s string) Number {
	if strings.Contains(s, ".") {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(fmt.Errorf("unexpected float parsing failure: %w", err))
		}
		return newNumberFloat(v)
	}

	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(fmt.Errorf("unexpected int parsing failure: %w", err))
	}
	return newNumberInt(int32(v))
}

func (n Number) IsFloat() bool {
	return n.isFloat
}

func (n Number) IsInt() bool {
	return !n.IsFloat()
}

func (n Number) ToFloat() float64 {
	if n.IsFloat() {
		return n.floatValue
	}
	return float64(n.intValue)
}

func (n Number) IsZero() bool {
	if n.IsFloat() {
		return n.floatValue == 0
	}
	return n.intValue == 0
}

func (n Number) ToInt() int32 {
	if n.IsFloat() {
		return int32(n.floatValue)
	}
	return n.intValue
}

func newNumberFloat(x float64) Number {
	return Number{
		isFloat:    true,
		floatValue: x,
	}
}

func newNumberInt(x int32) Number {
	return Number{
		isFloat:  false,
		intValue: x,
	}
}

var AddNumber = numberBinaryOp(
	func(a, b float64) Number { return newNumberFloat(a + b) },
	func(a, b int32) Number { return newNumberInt(a + b) },
)

var SubtractNumber = numberBinaryOp(
	func(a, b float64) Number { return newNumberFloat(a - b) },
	func(a, b int32) Number { return newNumberInt(a - b) },
)

var MultiplyNumber = numberBinaryOp(
	func(a, b float64) Number { return newNumberFloat(a * b) },
	func(a, b int32) Number { return newNumberInt(a * b) },
)

func DivideNumber(a Number, b Number) optional.Of[Number] {
	if b.IsZero() {
		return optional.Empty[Number]()
	}

	if a.IsInt() && b.IsInt() {
		fResult := a.ToFloat() / b.ToFloat()
		iResult := a.ToInt() / b.ToInt()
		if fResult == float64(iResult) {
			return optional.Value(newNumberInt(iResult))
		}
		return optional.Value(newNumberFloat(fResult))
	}
	return optional.Value(newNumberFloat(a.ToFloat() / b.ToFloat()))
}

func ModInteger(a Number, b Number) optional.Of[Number] {
	if b.IsZero() || !a.IsInt() || !b.IsInt() {
		return optional.Empty[Number]()
	}

	return optional.Value(newNumberInt(a.ToInt() % b.ToInt()))
}

func DivInteger(a Number, b Number) optional.Of[Number] {
	if b.IsZero() || !a.IsInt() || !b.IsInt() {
		return optional.Empty[Number]()
	}

	return optional.Value(newNumberInt(int32(math.Floor(a.ToFloat() / b.ToFloat()))))
}

var EqualNumber = numberBinaryOp(
	func(a, b float64) bool { return a == b },
	func(a, b int32) bool { return a == b },
)

var GreaterThanNumber = numberBinaryOp(
	func(a, b float64) bool { return a > b },
	func(a, b int32) bool { return a > b },
)

var GreaterThanOrEqualNumber = numberBinaryOp(
	func(a, b float64) bool { return a >= b },
	func(a, b int32) bool { return a >= b },
)

var LessThanNumber = numberBinaryOp(
	func(a, b float64) bool { return a < b },
	func(a, b int32) bool { return a < b },
)

var LessThanOrEqualNumber = numberBinaryOp(
	func(a, b float64) bool { return a <= b },
	func(a, b int32) bool { return a <= b },
)

func numberBinaryOp[T any](ff func(a, b float64) T, fi func(a, b int32) T) func(Number, Number) T {
	return func(a Number, b Number) T {
		if a.IsInt() && b.IsInt() {
			return fi(a.ToInt(), b.ToInt())
		}
		return ff(a.ToFloat(), b.ToFloat())
	}
}
