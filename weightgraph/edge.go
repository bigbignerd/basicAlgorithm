package weightgraph

import (
// "log"
)

//边的权重
type Weight interface{}

//边
type Edge struct {
	a, b   int //边的两个点
	weight Weight
}

func NewEdge(a, b int, weight Weight) *Edge {
	return &Edge{
		a:      a,
		b:      b,
		weight: weight,
	}
}

func (e *Edge) V() int {
	return e.a
}
func (e *Edge) W() int {
	return e.b
}
func (e *Edge) Weight() Weight {
	return e.weight
}

//已知边的一个点 获取另外一个点
func (e *Edge) Other(x int) int {
	if x == e.a {
		return e.b
	}
	return e.a
}

//比较edge权重
func (e *Edge) Less(a, b *Edge) bool {

	return (*a).weight.(float64) < (*b).weight.(float64)
}
