package optional

type Of[T any] struct {
	x       T
	isEmpty bool
}

func Value[T any](x T) Of[T] {
	return Of[T]{x: x, isEmpty: false}
}

func ValueOf[T any](x Of[T]) T {
	return x.Value()
}

func Empty[T any]() Of[T] {
	return Of[T]{isEmpty: true}
}

func (o Of[T]) Value() T {
	return o.x
}

func (o Of[T]) IsEmpty() bool {
	return o.isEmpty
}

func (o Of[T]) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Of[T]) Return() (T, bool) {
	return o.x, !o.isEmpty
}

func New[T any](node T, exists bool) Of[T] {
	return Of[T]{x: node, isEmpty: !exists}
}
