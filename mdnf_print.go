package ga

import "strings"

// MDNF formats a set of paths as a minimal disjunctive normal form.
func MDNF(paths []Path) string {
	terms := make([]string, len(paths))
	for i, p := range paths {
		terms[i] = strings.Join(p.EdgeIDs, " ")
	}
	return strings.Join(terms, " | ")
}
