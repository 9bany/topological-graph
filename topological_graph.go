package topologicalgraph

func NewGraph[k comparable]() *Graph[k] {
	return &Graph[k]{
		nodes: make(map[k]node[k]),
		check: make(map[k]bool),
	}
}

type Graph[key comparable] struct {
	nodes map[key]node[key]
	check map[key]bool
}

func (g *Graph[key]) Add(from, to key) {
	n := g.getNode(from)
	n.addEdge(to)
}

func (g *Graph[key]) Sort(k key) []key {
	s := g.sort(k)
	g.check = make(map[key]bool)
	return s
}

func (g *Graph[key]) sort(k key) (result []key) {
	n := g.getNode(k)
	edges := n.getEdges()

	if len(edges) == 0 {
		if _, ok := g.check[k]; !ok {
			result = append(result, k)
			g.check[k] = true
		}
	}

	for _, e := range edges {
		result = append(result, g.sort(e)...)
	}

	if _, ok := g.check[k]; !ok {
		result = append(result, k)
		g.check[k] = true
	}
	return
}

func (g *Graph[key]) getNode(k key) node[key] {
	n, ok := g.nodes[k]
	if !ok {
		n = make(node[key])
		g.nodes[k] = n
	}
	return n
}

type node[key comparable] map[key]bool

func (n node[key]) addEdge(k key) {
	n[k] = true
}

func (n node[key]) getEdges() (result []key) {
	for k := range n {
		result = append(result, k)
	}
	return
}
