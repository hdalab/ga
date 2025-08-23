package ga

type Edge struct{ ID string; From, To int }

type Graph struct {
    N     int
    Edges []Edge
    Out   [][]int
    In    [][]int
}

func New(n int) *Graph { return &Graph{N: n} }

func (g *Graph) AddEdge(id string, u, v int) {
    g.Edges = append(g.Edges, Edge{ID: id, From: u, To: v})
}

func (g *Graph) BuildAdj() {
    g.Out = make([][]int, g.N)
    g.In  = make([][]int, g.N)
    for i, e := range g.Edges {
        _ = i
        g.Out[e.From] = append(g.Out[e.From], i)
        g.In[e.To]    = append(g.In[e.To], i)
    }
}

type Path struct {
    Nodes   []int
    EdgeIDs []string
}

type Stats struct {
    NumPaths      int
    ElapsedNS     int64
    NsPerPath     float64
    NodesExpanded int64
    Pruned        int64
}
