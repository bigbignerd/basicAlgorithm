package graph

import (
// "log"
)

//稀疏图 - 邻接表
type SparseGraph struct {
	n, m     int
	directed bool
	g        [][]int
}

func NewSGraph(n int, directed bool) *SparseGraph {
	var g [][]int = make([][]int, n)
	for k := range g {
		g[k] = make([]int, 0)
	}
	return &SparseGraph{
		n:        n,
		m:        0,
		g:        g,
		directed: directed,
	}
}

//节点数
func (sg *SparseGraph) V() int {
	return sg.n
}

//边数
func (sg *SparseGraph) E() int {
	return sg.m
}

//添加边
func (sg *SparseGraph) AddEdge(v, w int) {
	if v < 0 || v > sg.n-1 || w < 0 || w > sg.n-1 {
		return
	}
	// log.Printf("add %d, %d", v, w)
	sg.g[v] = append(sg.g[v], w)
	//自环 或者 无向的情况
	if v != w && !sg.directed {
		sg.g[w] = append(sg.g[w], v)
	}
	sg.m++
}

//是否有v->w的边
func (sg *SparseGraph) HasEdge(v, w int) bool {
	if v < 0 || v > sg.n-1 || w < 0 || w > sg.n-1 {
		return false
	}
	for _, value := range sg.g[v] {
		if value == w {
			return true
		}
	}
	return false
}

//遍历临边
//c++的实现采用了一个迭代器来避免将g直接暴漏给用户
func (sg *SparseGraph) AdjIterator(v int) []int {
	return sg.g[v]
}
