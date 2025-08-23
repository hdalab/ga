package ga

import "context"

type EnumOptions struct {
    WithNodes bool
    Parallel  int
    MaxPaths  int
}

type PathEnumerator interface {
    Name() string
    Enumerate(ctx context.Context, g *Graph, s, t int, opt EnumOptions, emit func(Path) bool) (Stats, error)
}

func EnumerateMDNF(ctx context.Context, g *Graph, s, t int, opt EnumOptions, emit func(Path) bool) (Stats, error) {
    return dfsminorEnumerate(ctx, g, s, t, opt, emit)
}
