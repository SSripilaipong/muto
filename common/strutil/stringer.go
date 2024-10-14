package strutil

import "fmt"

func ToString[T fmt.Stringer](x T) string {
	return x.String()
}
