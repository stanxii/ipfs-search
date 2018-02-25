package index

import (
	"github.com/ipfs-search/ipfs-search/index/properties"
)

type Key struct {
	Hash  string
	Index string
	Type  string
}

type Item struct {
	*Key
	Properties *properties.Properties
}
