package weightgraph

type Graph interface {
	AdjIterator(v int) []int
	V() int
	E() int
	AddEdge(v, w int, weight Weight)
	HasEdge(v, w int) (bool, *Edge)
	GetEdge(v, w int) *Edge
}
