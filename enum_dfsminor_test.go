package ga

import (
	"context"
	"testing"
)

// DAG with a dead end: 0->1->2->3 (T), plus 1->4 which doesn't lead to T.
func TestStatsPrunedReachability(t *testing.T) {
	g := New(5)
	g.AddEdge("a", 0, 1)
	g.AddEdge("b", 1, 2)
	g.AddEdge("c", 2, 3)
	g.AddEdge("x", 1, 4)
	g.BuildAdj()

	stats, err := EnumerateMDNF(context.Background(), g, 0, 3, EnumOptions{}, func(Path) bool { return true })
	if err != nil {
		t.Fatalf("EnumerateMDNF: %v", err)
	}
	if stats.NumPaths != 1 {
		t.Fatalf("NumPaths = %d, want 1", stats.NumPaths)
	}
	if stats.Pruned < 1 {
		t.Fatalf("Pruned = %d, want >=1", stats.Pruned)
	}
	if stats.NodesExpanded <= 0 {
		t.Fatalf("NodesExpanded = %d, want >0", stats.NodesExpanded)
	}
}

// Graph with a cycle via back edge 2->1.
func TestStatsPrunedVisited(t *testing.T) {
	g := New(4)
	g.AddEdge("a", 0, 1)
	g.AddEdge("b", 1, 2)
	g.AddEdge("c", 2, 3)
	g.AddEdge("back", 2, 1)
	g.BuildAdj()

	stats, err := EnumerateMDNF(context.Background(), g, 0, 3, EnumOptions{}, func(Path) bool { return true })
	if err != nil {
		t.Fatalf("EnumerateMDNF: %v", err)
	}
	if stats.NumPaths != 1 {
		t.Fatalf("NumPaths = %d, want 1", stats.NumPaths)
	}
	if stats.Pruned < 1 {
		t.Fatalf("Pruned = %d, want >=1", stats.Pruned)
	}
	if stats.NodesExpanded <= 0 {
		t.Fatalf("NodesExpanded = %d, want >0", stats.NodesExpanded)
	}
}

// Example from README (a..i edges).
func TestExampleStats(t *testing.T) {
	g := New(6)
	g.AddEdge("a", 0, 1)
	g.AddEdge("b", 0, 2)
	g.AddEdge("c", 1, 2)
	g.AddEdge("d", 1, 3)
	g.AddEdge("e", 2, 3)
	g.AddEdge("f", 2, 4)
	g.AddEdge("g", 3, 4)
	g.AddEdge("h", 4, 5)
	g.AddEdge("i", 2, 5)
	g.BuildAdj()

	stats, err := EnumerateMDNF(context.Background(), g, 0, 5, EnumOptions{}, func(Path) bool { return true })
	if err != nil {
		t.Fatalf("EnumerateMDNF: %v", err)
	}
	if stats.NumPaths != 7 {
		t.Fatalf("NumPaths = %d, want 7", stats.NumPaths)
	}
	if stats.Pruned != 0 {
		t.Fatalf("Pruned = %d, want 0", stats.Pruned)
	}
	if stats.NodesExpanded <= 0 {
		t.Fatalf("NodesExpanded = %d, want >0", stats.NodesExpanded)
	}
}
