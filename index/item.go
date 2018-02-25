package index

import (
	"github.com/ipfs-search/ipfs-search/index/properties"
)

type Key struct {
	Index string
	Id    string
	Type  string
}

type Item struct {
	*Key
	Properties *properties.Properties
}
