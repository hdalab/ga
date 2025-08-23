package ga

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Spec struct {
	G    Graph
	S, T int
}

// ParseGexp parses a tiny DSL:
//
//	V: 0,1,2
//	a: 0->1
//	S: 0
//	T: 2
func ParseGexp(b []byte) (*Spec, error) {
	sc := bufio.NewScanner(bytes.NewReader(b))
	var edges []Edge
	maxv, nFromV := -1, -1
	var s, t int
	hasS, hasT := false, false
	lineno := 0

	for sc.Scan() {
		lineno++
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch {
		case strings.HasPrefix(line, "V:"):
			list := strings.TrimSpace(strings.TrimPrefix(line, "V:"))
			if list != "" {
				for _, tok := range strings.Split(list, ",") {
					tok = strings.TrimSpace(tok)
					if tok == "" {
						continue
					}
					v, err := strconv.Atoi(tok)
					if err != nil {
						return nil, fmt.Errorf("line %d: bad vertex %q", lineno, tok)
					}
					if v > maxv {
						maxv = v
					}
				}
				nFromV = maxv + 1
			}
		case strings.HasPrefix(line, "S:"):
			val := strings.TrimSpace(strings.TrimPrefix(line, "S:"))
			x, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("line %d: bad S %q", lineno, val)
			}
			s, hasS = x, true
			if s > maxv {
				maxv = s
			}
		case strings.HasPrefix(line, "T:"):
			val := strings.TrimSpace(strings.TrimPrefix(line, "T:"))
			x, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("line %d: bad T %q", lineno, val)
			}
			t, hasT = x, true
			if t > maxv {
				maxv = t
			}
		default:
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("line %d: expected 'id: u->v'", lineno)
			}
			id := strings.TrimSpace(parts[0])
			uv := strings.Split(strings.TrimSpace(parts[1]), "->")
			if len(uv) != 2 {
				return nil, fmt.Errorf("line %d: expected 'u->v'", lineno)
			}
			u, err1 := strconv.Atoi(strings.TrimSpace(uv[0]))
			v, err2 := strconv.Atoi(strings.TrimSpace(uv[1]))
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("line %d: bad edge vertices", lineno)
			}
			edges = append(edges, Edge{ID: id, From: u, To: v})
			if u > maxv {
				maxv = u
			}
			if v > maxv {
				maxv = v
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if !hasS || !hasT {
		return nil, fmt.Errorf("S and T must be provided")
	}

	n := nFromV
	if n < 0 {
		n = maxv + 1
	}

	g := New(n)
	if s < 0 || s >= g.N {
		return nil, fmt.Errorf("S=%d out of range [0..%d)", s, g.N)
	}
	if t < 0 || t >= g.N {
		return nil, fmt.Errorf("T=%d out of range [0..%d)", t, g.N)
	}
	for _, e := range edges {
		if e.From < 0 || e.From >= g.N || e.To < 0 || e.To >= g.N {
			return nil, fmt.Errorf("edge %q: %d->%d out of range [0..%d)", e.ID, e.From, e.To, g.N)
		}
		g.AddEdge(e.ID, e.From, e.To)
	}
	return &Spec{G: *g, S: s, T: t}, nil
}
