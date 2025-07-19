package rods

import (
	"maps"

	"github.com/SSripilaipong/muto/common/optional"
)

type Map[K comparable, V any] struct {
	data map[K]V
}

func NewMap[K comparable, V any](data map[K]V) Map[K, V] {
	return Map[K, V]{data: data}
}

func (m Map[K, V]) Get(key K) optional.Of[V] {
	v, ok := m.data[key]
	return optional.New(v, ok)
}

func (m Map[K, V]) ToMap() map[K]V {
	return maps.Clone(m.data)
}
