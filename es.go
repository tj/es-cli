package es

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tj/es/nodes/nodestats"
)

// Client for Elasticsearch.
type Client struct {
	Address string
}

// New client from `addr`.
func New(addr string) *Client {
	return &Client{
		Address: addr,
	}
}

// NodeStats returns all node stats.
func (c *Client) NodeStats() (v *nodestats.Stats, err error) {
	res, err := http.Get(fmt.Sprintf("%s/_nodes/stats", c.Address))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, err
	}

	return v, nil
}
