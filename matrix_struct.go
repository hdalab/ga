package ga

type Structural [][]string

func StructuralMatrix(g *Graph) Structural {
    ms := make(Structural, g.N)
    for i := range ms { ms[i] = make([]string, g.N) }
    for _, e := range g.Edges { ms[e.From][e.To] = e.ID }
    return ms
}
