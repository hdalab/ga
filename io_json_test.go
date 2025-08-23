package ga

import "testing"

func TestParseJson(t *testing.T) {
	gexp := `a: 0->1
b: 0->2
c: 1->2
d: 1->3
e: 2->3
f: 2->4
g: 3->4
h: 4->5
i: 2->5
S: 0
T: 5
`
	spec1, err := ParseGexp([]byte(gexp))
	if err != nil {
		t.Fatalf("ParseGexp: %v", err)
	}
	jsonSpec := `{"n":6,"edges":[{"id":"a","from":0,"to":1},{"id":"b","from":0,"to":2},{"id":"c","from":1,"to":2},{"id":"d","from":1,"to":3},{"id":"e","from":2,"to":3},{"id":"f","from":2,"to":4},{"id":"g","from":3,"to":4},{"id":"h","from":4,"to":5},{"id":"i","from":2,"to":5}],"s":0,"t":5}`
	spec2, err := ParseJson([]byte(jsonSpec))
	if err != nil {
		t.Fatalf("ParseJson: %v", err)
	}
	if spec1.G.N != spec2.G.N || spec1.S != spec2.S || spec1.T != spec2.T {
		t.Fatalf("spec mismatch: %+v vs %+v", spec1, spec2)
	}
	if len(spec1.G.Edges) != len(spec2.G.Edges) {
		t.Fatalf("edge count mismatch: %d vs %d", len(spec1.G.Edges), len(spec2.G.Edges))
	}
	for i, e := range spec1.G.Edges {
		if e != spec2.G.Edges[i] {
			t.Fatalf("edge %d mismatch: %v vs %v", i, e, spec2.G.Edges[i])
		}
	}
}
