package ga

import "context"

// EnumOptions configures path enumeration.
//
// WithNodes includes the list of visited nodes in each emitted Path.
// MaxPaths limits the total number of paths; zero means no limit.
type EnumOptions struct {
	WithNodes bool
	MaxPaths  int
}

// PathEnumerator enumerates all simple s→t paths in a graph.
type PathEnumerator interface {
	Name() string
	Enumerate(ctx context.Context, g *Graph, s, t int, opt EnumOptions, emit func(Path) bool) (Stats, error)
}

// EnumerateMDNF enumerates all simple paths from s to t and returns
// aggregated statistics.
func EnumerateMDNF(ctx context.Context, g *Graph, s, t int, opt EnumOptions, emit func(Path) bool) (Stats, error) {
	return dfsminorEnumerate(ctx, g, s, t, opt, emit)
}
