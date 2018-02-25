package index

import (
	"github.com/ipfs/go-ipfs-api"
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
