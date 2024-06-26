package graph

import "github.com/jayaprasanna31/GOLANGDATASTRUCTURE/queue"

// Graph represents a graph data structure
type Graph struct {
	Vertices []*Vertex
}

// Vertex represents a vertex in the graph
type Vertex struct {
	Value any
	Edges []*Edge
}

// Edge represents an edge in the graph
type Edge struct {
	Source *Vertex
	Target *Vertex
	Weight int
}

// InitializeGraph initializes a graph
func InitializeGraph() *Graph {
	return &Graph{}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(value any) *Vertex {
	vertex := &Vertex{Value: value}
	g.Vertices = append(g.Vertices, vertex)
	return vertex
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(source, target *Vertex, weight int) {
	edge := &Edge{Source: source, Target: target, Weight: weight}
	source.Edges = append(source.Edges, edge)
}

// Display displays the graph
func (g *Graph) Display() {
	for _, vertex := range g.Vertices {
		print(vertex.Value, " -> ")
		for _, edge := range vertex.Edges {
			print(edge.Target.Value, " ")
		}
		println()
	}
}

// DepthFirstSearch performs a depth-first search on the graph
func (g *Graph) DepthFirstSearch() {
	visited := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		if !visited[vertex] {
			g.dfs(vertex, visited)
		}
	}
}

func (g *Graph) dfs(vertex *Vertex, visited map[*Vertex]bool) {
	visited[vertex] = true
	print(vertex.Value, " ")
	for _, edge := range vertex.Edges {
		if !visited[edge.Target] {
			g.dfs(edge.Target, visited)
		}
	}
}

// BreadthFirstSearch performs a breadth-first search on the graph
func (g *Graph) BreadthFirstSearch() {
	visited := make(map[*Vertex]bool)
	queue := queue.InitQueue()
	for _, vertex := range g.Vertices {
		if !visited[vertex] {
			queue.Enqueue(vertex)
			visited[vertex] = true
			for !queue.IsEmpty() {
				current := queue.Dequeue().(*Vertex)
				print(current.Value, " ")
				for _, edge := range current.Edges {
					if !visited[edge.Target] {
						queue.Enqueue(edge.Target)
						visited[edge.Target] = true
					}
				}
			}
		}
	}
}

// Dijkstra performs Dijkstra's algorithm on the graph
func (g *Graph) Dijkstra(source *Vertex) {
	distances := make(map[*Vertex]int)
	previous := make(map[*Vertex]*Vertex)
	queue := queue.InitQueue()
	for _, vertex := range g.Vertices {
		if vertex == source {
			distances[vertex] = 0
		} else {
			distances[vertex] = infinity
		}
		previous[vertex] = nil
		queue.Enqueue(vertex)
	}
	for !queue.IsEmpty() {
		current := g.minDistance(queue, distances)
		for _, edge := range current.Edges {
			alt := distances[current] + edge.Weight
			if alt < distances[edge.Target] {
				distances[edge.Target] = alt
				previous[edge.Target] = current
			}
		}
	}
}

func (g *Graph) minDistance(queue *queue.Queue, distances map[*Vertex]int) *Vertex {
	min := infinity
	var minVertex *Vertex
	for _, vertex := range queue.Elements() {
		if distances[vertex.(*Vertex)] < min {
			min = distances[vertex.(*Vertex)]
			minVertex = vertex.(*Vertex)
		}
	}
	queue.DeleteByValue(minVertex)
	return minVertex
}

const infinity = int(^uint(0) >> 1)

type any interface{}

// BellmanFord performs the Bellman-Ford algorithm on the graph
func (g *Graph) BellmanFord(source *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex, bool) {
	distances := make(map[*Vertex]int)
	previous := make(map[*Vertex]*Vertex)

	// Step 1: Initialize distances and previous
	for _, vertex := range g.Vertices {
		if vertex == source {
			distances[vertex] = 0
		} else {
			distances[vertex] = infinity
		}
		previous[vertex] = nil
	}

	// Step 2: Relax edges repeatedly
	for i := 0; i < len(g.Vertices)-1; i++ {
		for _, vertex := range g.Vertices {
			for _, edge := range vertex.Edges {
				alt := distances[vertex] + edge.Weight
				if alt < distances[edge.Target] {
					distances[edge.Target] = alt
					previous[edge.Target] = vertex
				}
			}
		}
	}

	// Step 3: Check for negative cycles
	hasNegativeCycle := false
	for _, vertex := range g.Vertices {
		for _, edge := range vertex.Edges {
			alt := distances[vertex] + edge.Weight
			if alt < distances[edge.Target] {
				hasNegativeCycle = true
				break
			}
		}
	}

	return distances, previous, hasNegativeCycle
}

// FloydWarshall performs the Floyd-Warshall algorithm on the graph
func (g *Graph) FloydWarshall() (map[*Vertex]map[*Vertex]int, map[*Vertex]map[*Vertex]*Vertex) {
	distances := make(map[*Vertex]map[*Vertex]int)
	next := make(map[*Vertex]map[*Vertex]*Vertex)

	// Step 1: Initialize distances and next
	for _, vertex := range g.Vertices {
		distances[vertex] = make(map[*Vertex]int)
		next[vertex] = make(map[*Vertex]*Vertex)
		for _, other := range g.Vertices {
			distances[vertex][other] = infinity
			next[vertex][other] = nil
		}
		distances[vertex][vertex] = 0
		next[vertex][vertex] = vertex
		for _, edge := range vertex.Edges {
			distances[vertex][edge.Target] = edge.Weight
			next[vertex][edge.Target] = edge.Target
		}
	}

	// Step 2: Relax edges repeatedly
	for _, k := range g.Vertices {
		for _, i := range g.Vertices {
			for _, j := range g.Vertices {
				if distances[i][j] > distances[i][k]+distances[k][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	return distances, next
}

// Prim performs Prim's algorithm on the graph
func (g *Graph) Prim() (map[*Vertex]*Vertex, int) {
	parents := make(map[*Vertex]*Vertex)
	distances := make(map[*Vertex]int)
	visited := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		distances[vertex] = infinity
	}
	distances[g.Vertices[0]] = 0
	for i := 0; i < len(g.Vertices); i++ {
		current := g.minDistancePrim(distances, visited)
		visited[current] = true
		for _, edge := range current.Edges {
			if !visited[edge.Target] && edge.Weight < distances[edge.Target] {
				parents[edge.Target] = current
				distances[edge.Target] = edge.Weight
			}
		}
	}
	totalWeight := 0
	for _, distance := range distances {
		totalWeight += distance
	}
	return parents, totalWeight
}

func (g *Graph) minDistancePrim(distances map[*Vertex]int, visited map[*Vertex]bool) *Vertex {
	min := infinity
	var minVertex *Vertex
	for vertex, distance := range distances {
		if !visited[vertex] && distance < min {
			min = distance
			minVertex = vertex
		}
	}
	return minVertex
}

// Kruskal performs Kruskal's algorithm on the graph
func (g *Graph) Kruskal() (map[*Vertex]*Vertex, int) {
	parents := make(map[*Vertex]*Vertex)
	ranks := make(map[*Vertex]int)
	for _, vertex := range g.Vertices {
		parents[vertex] = vertex
		ranks[vertex] = 0
	}
	edges := make([]*Edge, 0)
	for _, vertex := range g.Vertices {
		for _, edge := range vertex.Edges {
			edges = append(edges, edge)
		}
	}
	quickSort(edges, 0, len(edges)-1)
	totalWeight := 0
	for _, edge := range edges {
		sourceRoot := find(edge.Source, parents)
		targetRoot := find(edge.Target, parents)
		if sourceRoot != targetRoot {
			if ranks[sourceRoot] < ranks[targetRoot] {
				parents[sourceRoot] = targetRoot
			} else if ranks[sourceRoot] > ranks[targetRoot] {
				parents[targetRoot] = sourceRoot
			} else {
				parents[targetRoot] = sourceRoot
				ranks[sourceRoot]++
			}
			totalWeight += edge.Weight
		}
	}
	return parents, totalWeight
}

func find(vertex *Vertex, parents map[*Vertex]*Vertex) *Vertex {
	if parents[vertex] != vertex {
		parents[vertex] = find(parents[vertex], parents)
	}
	return parents[vertex]
}

func quickSort(edges []*Edge, low, high int) {
	if low < high {
		pivot := partition(edges, low, high)
		quickSort(edges, low, pivot-1)
		quickSort(edges, pivot+1, high)
	}
}

func partition(edges []*Edge, low, high int) int {
	pivot := edges[high]
	i := low
	for j := low; j < high; j++ {
		if edges[j].Weight < pivot.Weight {
			edges[i], edges[j] = edges[j], edges[i]
			i++
		}
	}
	edges[i], edges[high] = edges[high], edges[i]
	return i
}

// TopologicalSort performs a topological sort on the graph
func (g *Graph) TopologicalSort() []any {
	stack := make([]any, 0)
	visited := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		if !visited[vertex] {
			g.topologicalSort(vertex, visited, &stack)
		}
	}
	return stack
}

func (g *Graph) topologicalSort(vertex *Vertex, visited map[*Vertex]bool, stack *[]any) {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if !visited[edge.Target] {
			g.topologicalSort(edge.Target, visited, stack)
		}
	}
	*stack = append(*stack, vertex.Value)
}

// IsCyclic checks if the graph contains a cycle
func (g *Graph) IsCyclic() bool {
	visited := make(map[*Vertex]bool)
	recStack := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		if g.isCyclic(vertex, visited, recStack) {
			return true
		}
	}
	return false
}

func (g *Graph) isCyclic(vertex *Vertex, visited, recStack map[*Vertex]bool) bool {
	if !visited[vertex] {
		visited[vertex] = true
		recStack[vertex] = true
		for _, edge := range vertex.Edges {
			if !visited[edge.Target] && g.isCyclic(edge.Target, visited, recStack) {
				return true
			} else if recStack[edge.Target] {
				return true
			}
		}
	}
	recStack[vertex] = false
	return false
}

// IsConnected checks if the graph is connected
func (g *Graph) IsConnected() bool {
	visited := make(map[*Vertex]bool)
	g.dfs(g.Vertices[0], visited)
	for _, vertex := range g.Vertices {
		if !visited[vertex] {
			return false
		}
	}
	return true
}

// IsStronglyConnected checks if the graph is strongly connected
func (g *Graph) IsStronglyConnected() bool {
	for _, vertex := range g.Vertices {
		visited := make(map[*Vertex]bool)
		g.dfs(vertex, visited)
		for _, other := range g.Vertices {
			if !visited[other] {
				return false
			}
		}
	}
	return true
}

// IsBipartite checks if the graph is bipartite
func (g *Graph) IsBipartite() bool {
	colors := make(map[*Vertex]int)
	for _, vertex := range g.Vertices {
		if colors[vertex] == 0 && !g.isBipartite(vertex, colors, 1) {
			return false
		}
	}
	return true
}

func (g *Graph) isBipartite(vertex *Vertex, colors map[*Vertex]int, color int) bool {
	colors[vertex] = color
	for _, edge := range vertex.Edges {
		if colors[edge.Target] == 0 {
			if !g.isBipartite(edge.Target, colors, -color) {
				return false
			}
		} else if colors[edge.Target] == colors[vertex] {
			return false
		}
	}
	return true
}

// IsEulerian checks if the graph is Eulerian
func (g *Graph) IsEulerian() bool {
	if !g.IsConnected() {
		return false
	}
	odd := 0
	for _, vertex := range g.Vertices {
		if len(vertex.Edges)%2 != 0 {
			odd++
		}
	}
	return odd == 0 || odd == 2
}

// IsHamiltonian checks if the graph is Hamiltonian
func (g *Graph) IsHamiltonian() bool {
	if len(g.Vertices) < 3 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isHamiltonian(g.Vertices[0], visited, 1)
}

func (g *Graph) isHamiltonian(vertex *Vertex, visited map[*Vertex]bool, count int) bool {
	if count == len(g.Vertices) {
		return true
	}
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if !visited[edge.Target] && g.isHamiltonian(edge.Target, visited, count+1) {
			return true
		}
	}
	visited[vertex] = false
	return false
}

// IsPlanar checks if the graph is planar
func (g *Graph) IsPlanar() bool {
	if len(g.Vertices) <= 4 {
		return true
	}
	if len(g.Vertices) >= 5 && len(g.Vertices) <= 6 {
		for _, vertex := range g.Vertices {
			if len(vertex.Edges) < 3 {
				return true
			}
		}
	}
	return false
}

// IsComplete checks if the graph is complete
func (g *Graph) IsComplete() bool {
	for _, vertex := range g.Vertices {
		if len(vertex.Edges) != len(g.Vertices)-1 {
			return false
		}
	}
	return true
}

// IsTree checks if the graph is a tree
func (g *Graph) IsTree() bool {
	visited := make(map[*Vertex]bool)
	return !g.IsCyclic() && g.IsConnected() && g.isTree(g.Vertices[0], nil, visited)
}

func (g *Graph) isTree(vertex, parent *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if edge.Target != parent {
			if visited[edge.Target] || !g.isTree(edge.Target, vertex, visited) {
				return false
			}
		}
	}
	return true
}

// IsForest checks if the graph is a forest
func (g *Graph) IsForest() bool {
	visited := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		if visited[vertex] {
			return false
		}
		if !g.isTree(vertex, nil, visited) {
			return false
		}
	}
	return true
}

// IsIsomorphic checks if the graph is isomorphic to another graph
func (g *Graph) IsIsomorphic(other *Graph) bool {
	if len(g.Vertices) != len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isIsomorphic(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isIsomorphic(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isIsomorphic(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsSubgraph checks if the graph is a subgraph of another graph
func (g *Graph) IsSubgraph(other *Graph) bool {
	if len(g.Vertices) > len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isSubgraph(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isSubgraph(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isSubgraph(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsSubdivision checks if the graph is a subdivision of another graph
func (g *Graph) IsSubdivision(other *Graph) bool {
	if len(g.Vertices) > len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isSubdivision(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isSubdivision(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isSubdivision(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsMinor checks if the graph is a minor of another graph
func (g *Graph) IsMinor(other *Graph) bool {
	if len(g.Vertices) > len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isMinor(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isMinor(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isMinor(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsIsometric checks if the graph is isometric to another graph
func (g *Graph) IsIsometric(other *Graph) bool {
	if len(g.Vertices) != len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isIsometric(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isIsometric(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isIsometric(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsIsogonal checks if the graph is isogonal to another graph
func (g *Graph) IsIsogonal(other *Graph) bool {
	if len(g.Vertices) != len(other.Vertices) {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isIsogonal(other, g.Vertices[0], other.Vertices[0], visited)
}

func (g *Graph) isIsogonal(other *Graph, vertex1, vertex2 *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex1] = true
	for i, edge1 := range vertex1.Edges {
		if len(vertex1.Edges) != len(vertex2.Edges) {
			return false
		}
		edge2 := vertex2.Edges[i]
		if edge1.Weight != edge2.Weight {
			return false
		}
		if !visited[edge1.Target] && !g.isIsogonal(other, edge1.Target, edge2.Target, visited) {
			return false
		}
	}
	return true
}

// IsEquitable checks if the graph is equitable
func (g *Graph) IsEquitable() bool {
	if len(g.Vertices) < 3 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isEquitable(g.Vertices[0], visited)
}

func (g *Graph) isEquitable(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isEquitable(edge.Target, visited) {
			return false
		}
	}
	return true
}

// IsRegular checks if the graph is regular
func (g *Graph) IsRegular() bool {
	degree := len(g.Vertices[0].Edges)
	for _, vertex := range g.Vertices {
		if len(vertex.Edges) != degree {
			return false
		}
	}
	return true
}

// IsCompleteBipartite checks if the graph is a complete bipartite graph
func (g *Graph) IsCompleteBipartite() bool {
	if len(g.Vertices) < 4 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isCompleteBipartite(g.Vertices[0], 1, visited)
}

func (g *Graph) isCompleteBipartite(vertex *Vertex, color int, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isCompleteBipartite(edge.Target, -color, visited) {
			return false
		}
	}
	return true
}

// IsWheel checks if the graph is a wheel graph
func (g *Graph) IsWheel() bool {
	if len(g.Vertices) < 4 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isWheel(g.Vertices[0], visited)
}

func (g *Graph) isWheel(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isWheel(edge.Target, visited) {
			return false
		}
	}
	return true
}

// IsStar checks if the graph is a star graph
func (g *Graph) IsStar() bool {
	if len(g.Vertices) < 3 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isStar(g.Vertices[0], visited)
}

func (g *Graph) isStar(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isStar(edge.Target, visited) {
			return false
		}
	}
	return true
}

// IsCaterpillar checks if the graph is a caterpillar graph
func (g *Graph) IsCaterpillar() bool {
	if len(g.Vertices) < 3 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isCaterpillar(g.Vertices[0], visited)
}

func (g *Graph) isCaterpillar(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isCaterpillar(edge.Target, visited) {
			return false
		}
	}
	return true
}

// IsCycle checks if the graph is a cycle graph
func (g *Graph) IsCycle() bool {
	if len(g.Vertices) < 3 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isCycle(g.Vertices[0], nil, visited)
}

func (g *Graph) isCycle(vertex, parent *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if edge.Target != parent {
			if visited[edge.Target] || !g.isCycle(edge.Target, vertex, visited) {
				return false
			}
		}
	}
	return true
}

// IsPath checks if the graph is a path graph
func (g *Graph) IsPath() bool {
	if len(g.Vertices) < 2 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isPath(g.Vertices[0], nil, visited)
}

func (g *Graph) isPath(vertex, parent *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if edge.Target != parent {
			if visited[edge.Target] || !g.isPath(edge.Target, vertex, visited) {
				return false
			}
		}
	}
	return true
}

// IsGrid checks if the graph is a grid graph
func (g *Graph) IsGrid() bool {
	if len(g.Vertices) < 4 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isGrid(g.Vertices[0], visited)
}

func (g *Graph) isGrid(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isGrid(edge.Target, visited) {
			return false
		}
	}
	return true
}

// IsLadder checks if the graph is a ladder graph
func (g *Graph) IsLadder() bool {
	if len(g.Vertices) < 5 {
		return false
	}
	visited := make(map[*Vertex]bool)
	return g.isLadder(g.Vertices[0], visited)
}

func (g *Graph) isLadder(vertex *Vertex, visited map[*Vertex]bool) bool {
	visited[vertex] = true
	for _, edge := range vertex.Edges {
		if visited[edge.Target] {
			return false
		}
		if !g.isLadder(edge.Target, visited) {
			return false
		}
	}
	return true
}
