package index

import (
	"gopkg.in/olivere/elastic.v5"
)

type Indexer struct {
	*elastic.Client
}

func (i *Indexer) Index(name string) *Index {
	return &Index{
		Indexer: i,
		name:    name,
	}
}
