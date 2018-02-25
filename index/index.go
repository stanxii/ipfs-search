package index

import (
	"context"
	"github.com/ipfs-search/ipfs-search/index/properties"
	"gopkg.in/olivere/elastic.v5"
)

type Index struct {
	*Indexer
	name string
}

func (index *Index) Add(ctx context.Context, item *Item) error {
	_, err := index.Client.Index().
		Index(index.name).
		Type(item.Type).
		Id(item.Hash).
		BodyJson(item.Properties).
		Do(ctx)

	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (index *Index) Update(ctx context.Context, item *Item) error {
	_, err := index.Client.Update().
		Index(index.name).
		Type(item.Type).
		Id(item.Hash).
		Doc(item.Properties).
		Do(ctx)

	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (index *Index) Get(ctx context.Context, key *Key, fields []string) (*Item, error) {
	q := index.Client.Get().Index(index.name)

	// Conditionally, filter by type
	if key.Type != "" {
		q = q.Type(key.Type)
	}

	// Conditionally, filter fields
	if fields != nil {
		fsc := elastic.NewFetchSourceContext(true)
		fsc.Include(fields...)

		q = q.FetchSourceContext(fsc)
	}

	r, err := q.Id(key.Hash).Do(ctx)

	if err != nil {
		return nil, err
	}

	properties, err := properties.FromJSON(r.Source)

	if err != nil {
		return nil, err
	}

	item := &Item{
		Key:        key,
		Properties: properties,
	}

	return item, nil
}
