package ga

import "strings"

func MDNF(paths []Path) string {
    terms := make([]string, len(paths))
    for i, p := range paths {
        terms[i] = strings.Join(p.EdgeIDs, " ")
    }
    return strings.Join(terms, " | ")
}
