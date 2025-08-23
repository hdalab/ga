# ga — Graph Algebra (Go module)

Public API for **symbolic path enumeration** on directed graphs.
Build **structural matrices** and enumerate the **MDNF** (all simple s→t paths).

```bash
go get github.com/hdalab/ga@v0.1.1
```

### Quick example
```go
package main

import (
    "context"
    "fmt"
    "github.com/hdalab/ga"
)

func main() {
    spec, _ := ga.ParseGexp([]byte(`
a: 0->1
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
`))
    // alternatively:
    // spec, _ := ga.ParseJson([]byte(`{"n":6,"edges":[{"id":"a","from":0,"to":1},{"id":"b","from":0,"to":2},{"id":"c","from":1,"to":2},{"id":"d","from":1,"to":3},{"id":"e","from":2,"to":3},{"id":"f","from":2,"to":4},{"id":"g","from":3,"to":4},{"id":"h","from":4,"to":5},{"id":"i","from":2,"to":5}],"s":0,"t":5}`))
    var paths []ga.Path
    ga.EnumerateMDNF(context.Background(), &spec.G, spec.S, spec.T, ga.EnumOptions{}, func(p ga.Path) bool {
        paths = append(paths, p); return true
    })
    fmt.Println(ga.MDNF(paths))
}
```

### Stats

`EnumerateMDNF` returns `Stats` with two useful metrics:

- `NodesExpanded` — how many times DFS expanded a vertex;
- `Pruned` — how many edges were skipped either due to revisiting a vertex or a global reachability check.
