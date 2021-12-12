package graph

type ItemGraph struct {
	nodes []string
	edges map[string][]string
}

func (g *ItemGraph) addNode(n string) {
	g.nodes = append(g.nodes, n)
}

func (g *ItemGraph) GetNode(name string) string {
	for _, node := range g.nodes {
		if node == name {
			return node
		}
	}

	return ""
}

func (g *ItemGraph) NodeExists(n string) bool {
	for _, node := range g.nodes {
		if node == n {
			return true
		}
	}

	return false
}

func (g *ItemGraph) AddEdge(n1, n2 string) {
	if !g.NodeExists(n1) {
		g.addNode(n1)
	}

	if !g.NodeExists(n2) {
		g.addNode(n2)
	}

	if g.edges == nil {
		g.edges = make(map[string][]string)
	}

	g.edges[n1] = append(g.edges[n1], n2)
	g.edges[n2] = append(g.edges[n2], n1)
}

func (g *ItemGraph) NodeCount() int {
	return len(g.nodes)
}

func (g *ItemGraph) GetDestinations(source string) []string {
	return g.edges[source]
}
