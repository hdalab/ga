package ga

import (
    "context"
    "time"
)

func dfsminorEnumerate(ctx context.Context, g *Graph, s, t int, opt EnumOptions, emit func(Path) bool) (Stats, error) {
    start := time.Now()
    reachS := ReachableFrom(g, s)
    reachT := ReachableTo(g, t)

    visited := make([]bool, g.N)
    var nodes []int
    if opt.WithNodes { nodes = []int{s} }
    visited[s] = true
    curEdges := []string{}
    var stats Stats
    var stop bool

    var dfs func(u int)
    dfs = func(u int) {
        if stop { return }
        select { case <-ctx.Done(): stop = true; return; default: }
        if u == t {
            stats.NumPaths++
            p := Path{EdgeIDs: append([]string(nil), curEdges...)}
            if opt.WithNodes {
                p.Nodes = append([]int(nil), nodes...)
            }
            if !emit(p) { stop = true }
            return
        }
        for _, ei := range g.Out[u] {
            e := g.Edges[ei]; v := e.To
            if visited[v] { continue }
            if !reachS[u] || !reachT[v] { continue }
            visited[v] = true
            curEdges = append(curEdges, e.ID)
            if opt.WithNodes { nodes = append(nodes, v) }
            dfs(v)
            if opt.WithNodes { nodes = nodes[:len(nodes)-1] }
            curEdges = curEdges[:len(curEdges)-1]
            visited[v] = false
            if stop { return }
            if opt.MaxPaths > 0 && stats.NumPaths >= opt.MaxPaths { stop = true; return }
        }
    }

    dfs(s)
    stats.ElapsedNS = time.Since(start).Nanoseconds()
    if stats.NumPaths > 0 { stats.NsPerPath = float64(stats.ElapsedNS) / float64(stats.NumPaths) }
    return stats, ctx.Err()
}
