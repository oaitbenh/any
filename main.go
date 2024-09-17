package main

import "fmt"

var (
	Paths   = map[int][]int{}
	Visited = map[int]bool{}
)

func Djikstra(graph map[int][]int, start int) map[int][]int {
	Visited[start] = true
	for _, to := range graph[start] {
		Path := append(Paths[start], to)
		if len(Paths[to]) == 0 || len(Paths[to]) >= len(Path) {
			Paths[to] = Path
		}
		if !Visited[to] {
			Djikstra(graph, to)
		}
	}
	return Paths
}

func main() {
	graph := map[int][]int{
		0: {2},
		2: {3, 4, 1},
		3: {1},
		4: {2, 1},
	}
	start := 0
	end := 1
	PathOf := Djikstra(graph, start)
	for i, v := range PathOf[end] {
		if v == start {
			PathOf[end] = PathOf[end][i:]
		}
	}
	fmt.Println(PathOf)
}
