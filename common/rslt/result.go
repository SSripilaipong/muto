package rslt

import "fmt"

type Of[T any] struct {
	value T
	err   error
	isErr bool
}

func New[T any](value T, err error) Of[T] {
	if err != nil {
		return Error[T](err)
	}
	return Value(value)
}

func (r Of[T]) Return() (T, error) {
	return r.value, r.err
}

func (r Of[T]) Value() T {
	return r.value
}

func (r Of[T]) IsOk() bool {
	return !r.IsErr()
}

func (r Of[T]) IsErr() bool {
	return r.isErr
}

func (r Of[T]) Error() error {
	return r.err
}

func (r Of[T]) String() string {
	if r.isErr {
		return fmt.Sprintf("Error(%s)", fmt.Sprint(r.err))
	}
	return fmt.Sprintf("Value(%s)", fmt.Sprint(r.value))
}

func Value[T any](x T) Of[T] {
	return Of[T]{value: x}
}

func Error[T any](err error) Of[T] {
	return Of[T]{err: err, isErr: true}
}

func ErrorOf[T any](x Of[T]) error {
	return x.Error()
}
