package graph

type Graph interface {
	AdjIterator(v int) []int
	V() int
	E() int
	AddEdge(v, w int)
	HasEdge(v, w int) bool
}
