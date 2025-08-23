package ga

import (
	"encoding/json"
	"fmt"
)

// ParseJson parses a graph specification in JSON format:
//
//	{
//	  "n": 6,
//	  "edges": [{"id":"a","from":0,"to":1}, ...],
//	  "s": 0,
//	  "t": 5
//	}
func ParseJson(b []byte) (*Spec, error) {
	var js struct {
		N     int `json:"n"`
		Edges []struct {
			ID   string `json:"id"`
			From int    `json:"from"`
			To   int    `json:"to"`
		} `json:"edges"`
		S *int `json:"s"`
		T *int `json:"t"`
	}
	if err := json.Unmarshal(b, &js); err != nil {
		return nil, err
	}
	if js.S == nil || js.T == nil {
		return nil, fmt.Errorf("s and t must be provided")
	}
	edges := make([]Edge, len(js.Edges))
	maxv := -1
	for i, e := range js.Edges {
		edges[i] = Edge{ID: e.ID, From: e.From, To: e.To}
		if e.From > maxv {
			maxv = e.From
		}
		if e.To > maxv {
			maxv = e.To
		}
	}
	if *js.S > maxv {
		maxv = *js.S
	}
	if *js.T > maxv {
		maxv = *js.T
	}
	n := js.N
	if n <= 0 || n < maxv+1 {
		n = maxv + 1
	}
	g := New(n)
	for _, e := range edges {
		g.AddEdge(e.ID, e.From, e.To)
	}
	return &Spec{G: *g, S: *js.S, T: *js.T}, nil
}
