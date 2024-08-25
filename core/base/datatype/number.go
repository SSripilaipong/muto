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
		return Number{
			isFloat:    true,
			floatValue: float32(v),
		}
	}

	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(fmt.Errorf("unexpected int parsing failure: %w", err))
	}
	return Number{
		isFloat:  false,
		intValue: int32(v),
	}
}
