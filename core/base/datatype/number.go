package datatype

import (
	"fmt"
	"strconv"
	"strings"
)

type Number struct {
	isFloat    bool
	intValue   int32
	floatValue float32
}

func NewNumber(s string) Number {
	if strings.Contains(s, ".") {
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			panic(fmt.Errorf("unexpected float parsing failure: %w", err))
		}
		return newNumberFloat(float32(v))
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

func (n Number) ToFloat() float32 {
	if n.IsFloat() {
		return n.floatValue
	}
	return float32(n.intValue)
}

func (n Number) ToInt() int32 {
	if n.IsFloat() {
		return int32(n.floatValue)
	}
	return n.intValue
}

func newNumberFloat(x float32) Number {
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

func AddNumber(a Number, b Number) Number {
	if a.IsInt() && b.IsInt() {
		return newNumberInt(a.ToInt() + b.ToInt())
	}
	return newNumberFloat(a.ToFloat() + b.ToFloat())
}
