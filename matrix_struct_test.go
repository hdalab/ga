package ga

import (
	"reflect"
	"testing"
)

func TestStructuralMatrix(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges []Edge
		want  Structural
	}{
		{
			name:  "simple",
			n:     3,
			edges: []Edge{{ID: "a", From: 0, To: 1}, {ID: "b", From: 1, To: 2}},
			want: Structural{
				{"", "a", ""},
				{"", "", "b"},
				{"", "", ""},
			},
		},
		{
			name:  "cycle",
			n:     3,
			edges: []Edge{{ID: "a", From: 0, To: 1}, {ID: "b", From: 1, To: 2}, {ID: "c", From: 2, To: 0}},
			want: Structural{
				{"", "a", ""},
				{"", "", "b"},
				{"c", "", ""},
			},
		},
	}
	for _, tt := range tests {
		g := New(tt.n)
		for _, e := range tt.edges {
			g.AddEdge(e.ID, e.From, e.To)
		}
		got := StructuralMatrix(g)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: StructuralMatrix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
