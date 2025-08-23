package ga

import "testing"

func TestParseGexpOutOfRange(t *testing.T) {
	_, err := ParseGexp([]byte(`
V: 0,1
a: 0->2
S: 0
T: 1
`))
	if err == nil {
		t.Fatalf("expected error for edge out of range")
	}
}

func TestParseGexpSTOutOfRange(t *testing.T) {
	_, err := ParseGexp([]byte(`
V: 0
S: 1
T: 0
`))
	if err == nil {
		t.Fatalf("expected error for S out of range")
	}
}

func TestParseGexpOK(t *testing.T) {
	spec, err := ParseGexp([]byte(`
a: 0->1
S: 0
T: 1
`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.G.N != 2 || spec.S != 0 || spec.T != 1 || len(spec.G.Edges) != 1 {
		t.Fatalf("unexpected parse result: %+v", spec)
	}
}
