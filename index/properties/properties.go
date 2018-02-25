package properties

import (
	"encoding/json"
	"github.com/ipfs/go-ipfs-api"
	"log"
	"time"
)

type Metadata map[string]interface{}

type Properties struct {
	references References `json:"references"`
	size       uint64     `json:"size"`
	firstSeen  time.Time  `json:"first-seen"`
	lastSeen   time.Time  `json:"last-seen"`

	// Directories
	links []*shell.UnixLsLink `json:"links,omitempty"`

	// Files
	metadata Metadata `json:"metadata,omitempty"`
}

func New() *Properties {
	return &Properties{
		firstSeen: time.Now().UTC(),
	}
}

func FromJSON(j *json.RawMessage) (*Properties, error) {
	// Parse JSON into Item
	p := new(Properties)

	err := json.Unmarshal(*j, p)
	if err != nil {
		log.Printf("can't unmarshal JSON: %s", j)
		return nil, err
	}

	return p, nil
}

func (p *Properties) UpdateReferences(r *Reference) {
	if !p.references.Contains(r) {
		log.Printf("Adding reference '%v' to %v", r, p)
		p.references = append(p.references, *r)
	}
}
