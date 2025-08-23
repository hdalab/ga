package ga

// ReachableFrom returns vertices reachable from s.
func ReachableFrom(g *Graph, s int) []bool {
	r := make([]bool, g.N)
	st := []int{s}
	r[s] = true
	for len(st) > 0 {
		u := st[len(st)-1]
		st = st[:len(st)-1]
		for _, ei := range g.Out[u] {
			v := g.Edges[ei].To
			if !r[v] {
				r[v] = true
				st = append(st, v)
			}
		}
	}
	return r
}

// ReachableTo returns vertices that can reach t.
func ReachableTo(g *Graph, t int) []bool {
	r := make([]bool, g.N)
	st := []int{t}
	r[t] = true
	for len(st) > 0 {
		u := st[len(st)-1]
		st = st[:len(st)-1]
		for _, ei := range g.In[u] {
			v := g.Edges[ei].From
			if !r[v] {
				r[v] = true
				st = append(st, v)
			}
		}
	}
	return r
}
