package mutator

import "github.com/SSripilaipong/go-common/optional"

type CollectionMapping interface {
	GetCollection(name string) optional.Of[Collection]
}
