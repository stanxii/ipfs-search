package properties

import (
	"encoding/json"
	"github.com/ipfs/go-ipfs-api"
	"log"
	"time"
)

type Metadata map[string]interface{}

type Properties struct {
	References References `json:"references"`
	Size       uint64     `json:"size"`
	FirstSeen  time.Time  `json:"first-seen"`
	LastSeen   time.Time  `json:"last-seen"`

	// Directories
	Links []*shell.UnixLsLink `json:"links,omitempty"`

	// Files
	Metadata Metadata `json:"metadata,omitempty"`
}

func New() *Properties {
	return &Properties{
		FirstSeen: time.Now().UTC(),
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
	if !p.References.Contains(r) {
		log.Printf("Adding reference '%v' to %v", r, p)
		p.References = append(p.References, *r)
	}
}

func (p *Properties) UpdateLastSeen() {
	p.LastSeen = time.Now().UTC()
}
