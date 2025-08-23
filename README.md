# ga — Graph Algebra (Go module)

Public API for **symbolic path enumeration** on directed graphs.
Build **structural matrices** and enumerate the **MDNF** (all simple s→t paths).

```bash
go get github.com/hdalab/ga@v0.1.0
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
    var paths []ga.Path
    ga.EnumerateMDNF(context.Background(), &spec.G, spec.S, spec.T, ga.EnumOptions{}, func(p ga.Path) bool {
        paths = append(paths, p); return true
    })
    fmt.Println(ga.MDNF(paths))
}
```
