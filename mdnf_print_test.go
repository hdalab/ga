package ga

import "testing"

func TestMDNF(t *testing.T) {
	tests := []struct {
		name  string
		paths []Path
		want  string
	}{
		{name: "empty", want: ""},
		{name: "single", paths: []Path{{EdgeIDs: []string{"a", "b"}}}, want: "a b"},
		{name: "multi", paths: []Path{{EdgeIDs: []string{"a"}}, {EdgeIDs: []string{"b", "c"}}}, want: "a | b c"},
	}
	for _, tt := range tests {
		if got := MDNF(tt.paths); got != tt.want {
			t.Errorf("%s: MDNF() = %q, want %q", tt.name, got, tt.want)
		}
	}
}
