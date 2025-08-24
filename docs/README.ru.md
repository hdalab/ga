# ga — Алгебра графов (Go модуль)

Публичный API для **символьного перечисления путей** на направленных графах.
Стройте **структурные матрицы** и перечисляйте **МДНФ** (все простые пути s→t).

```bash
go get github.com/hdalab/ga@v0.1.1
```

### Быстрый пример
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
    // альтернативный вариант:
    // spec, _ := ga.ParseJson([]byte(`{"n":6,"edges":[{"id":"a","from":0,"to":1},{"id":"b","from":0,"to":2},{"id":"c","from":1,"to":2},{"id":"d","from":1,"to":3},{"id":"e","from":2,"to":3},{"id":"f","from":2,"to":4},{"id":"g","from":3,"to":4},{"id":"h","from":4,"to":5},{"id":"i","from":2,"to":5}],"s":0,"t":5}`))
    var paths []ga.Path
    ga.EnumerateMDNF(context.Background(), &spec.G, spec.S, spec.T, ga.EnumOptions{}, func(p ga.Path) bool {
        paths = append(paths, p); return true
    })
    fmt.Println(ga.MDNF(paths))
}
```

Рёбра, добавленные с помощью `AddEdge`, автоматически обновляют списки смежности; `BuildAdj`
может использоваться для их пересборки, если список рёбер был изменён напрямую.

### Статистика

`EnumerateMDNF` возвращает `Stats` с двумя полезными метриками:

- `NodesExpanded` — сколько раз DFS расширял вершину;
- `Pruned` — сколько рёбер было пропущено либо из-за повторного посещения вершины, либо из-за глобальной проверки достижимости.
