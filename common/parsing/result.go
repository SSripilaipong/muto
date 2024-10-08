package parsing

import "github.com/SSripilaipong/muto/common/tuple"

func FilterSuccess[R, S any](rs []tuple.Of2[R, []S]) (ss []tuple.Of2[R, []S]) {
	for _, r := range rs {
		if len(r.X2()) == 0 {
			ss = append(ss, r)
		}
	}
	return
}
