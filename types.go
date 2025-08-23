package ga

// Edge is a directed edge identified by an ID.
type Edge struct {
	ID       string
	From, To int
}

// Graph is a directed graph with adjacency lists.
//
// Edges holds all edges; Out and In provide adjacency indices updated by AddEdge.
type Graph struct {
	N     int
	Edges []Edge
	Out   [][]int
	In    [][]int
}

// New creates a graph with n vertices.
func New(n int) *Graph {
	return &Graph{
		N:   n,
		Out: make([][]int, n),
		In:  make([][]int, n),
	}
}

// AddEdge appends a directed edge and updates adjacency lists.
func (g *Graph) AddEdge(id string, u, v int) {
	ei := len(g.Edges)
	g.Edges = append(g.Edges, Edge{ID: id, From: u, To: v})
	g.Out[u] = append(g.Out[u], ei)
	g.In[v] = append(g.In[v], ei)
}

// BuildAdj rebuilds adjacency lists from the edge list.
// This is normally unnecessary because AddEdge keeps adjacency up to date.
func (g *Graph) BuildAdj() {
	g.Out = make([][]int, g.N)
	g.In = make([][]int, g.N)
	for i, e := range g.Edges {
		g.Out[e.From] = append(g.Out[e.From], i)
		g.In[e.To] = append(g.In[e.To], i)
	}
}

// Path represents a simple path.
// Nodes contains the sequence of vertex IDs if requested.
// EdgeIDs lists edge identifiers along the path.
type Path struct {
	Nodes   []int
	EdgeIDs []string
}

// Stats holds statistics collected during enumeration.
type Stats struct {
	NumPaths      int
	ElapsedNS     int64
	NsPerPath     float64
	NodesExpanded int64
	Pruned        int64
}
