package strutil

func WithPrefix(p string) func(s string) string {
	return func(s string) string {
		return p + s
	}
}
