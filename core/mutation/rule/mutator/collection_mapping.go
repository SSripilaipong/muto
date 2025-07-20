package mutator

import "github.com/SSripilaipong/muto/common/optional"

type CollectionMapping interface {
	GetCollection(name string) optional.Of[Collection]
}
