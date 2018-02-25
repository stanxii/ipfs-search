package index

import (
	"encoding/json"
	"log"
)

type Properties map[string]interface{}

type Key struct {
	Hash  string
	Index string
	Type  string
}

type Item struct {
	Key
	Properties Properties
}

func ItemFromJSON(source *json.RawMessage) (*Item, error) {
	// Parse JSON into Item
	item := new(Item)

	err := json.Unmarshal(*source, item)
	if err != nil {
		log.Printf("can't unmarshal JSON: %s", source)
		return nil, err
	}

	return item, nil
}
