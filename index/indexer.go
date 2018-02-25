package index

import (
	"context"
	"github.com/ipfs-search/ipfs-search/index/properties"
	"gopkg.in/olivere/elastic.v5"
)

type Indexer struct {
	*elastic.Client
}

func (i *Indexer) Add(ctx context.Context, item *Item) error {
	_, err := i.Client.Index().
		Index(item.Index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item.Properties).
		Do(ctx)

	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (i *Indexer) Update(ctx context.Context, item *Item) error {
	_, err := i.Client.Update().
		Index(item.Index).
		Type(item.Type).
		Id(item.Id).
		Doc(item.Properties).
		Do(ctx)

	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func (i *Indexer) Get(ctx context.Context, key *Key, fields []string) (*Item, error) {
	q := i.Client.Get().Index(key.Index)

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

	r, err := q.Id(key.Id).Do(ctx)

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
