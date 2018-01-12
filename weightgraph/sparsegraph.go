package weightgraph

import (
	"log"
	"strconv"
	"strings"
)

//稀疏图 - 邻接表
type SparseGraph struct {
	n, m     int
	directed bool
	g        [][]*Edge
}

func NewSGraph(n int, directed bool) *SparseGraph {
	var g [][]*Edge = make([][]*Edge, n)
	for k := range g {
		g[k] = make([]*Edge, 0)
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
func (sg *SparseGraph) GetEdge(v, w int) *Edge {
	if has, edge := sg.HasEdge(v, w); has {
		return edge
	}
	return nil
}

//添加边
func (sg *SparseGraph) AddEdge(v, w int, weight Weight) {
	if v < 0 || v > sg.n-1 || w < 0 || w > sg.n-1 {
		return
	}
	// log.Printf("add %d, %d", v, w)
	sg.g[v] = append(sg.g[v], NewEdge(v, w, weight))
	//自环 或者 无向的情况
	if v != w && !sg.directed {
		sg.g[w] = append(sg.g[w], NewEdge(w, v, weight))
	}
	sg.m++
}

//是否有v->w的边
func (sg *SparseGraph) HasEdge(v, w int) (bool, *Edge) {
	if v < 0 || v > sg.n-1 || w < 0 || w > sg.n-1 {
		log.Fatal("v,w index out of range...")
		return false, nil
	}
	for _, edge := range sg.g[v] {
		if (*edge).Other(v) == w {
			return true, edge
		}
	}
	return false, nil
}

//遍历临边
//c++的实现采用了一个迭代器来避免将g直接暴漏给用户
func (sg *SparseGraph) AdjIterator(v int) []int {
	if v < 0 || v >= sg.n {
		log.Fatal("adj interator v越界")
		return nil
	}
	var limb []int = make([]int, 0)
	for _, edge := range sg.g[v] {
		if edge != nil {
			limb = append(limb, (*edge).Other(v))
		}
	}
	return limb
}
func (sg *SparseGraph) show() {
	for i := 0; i < sg.V(); i++ {
		var str []string = make([]string, 0)
		for _, edge := range sg.g[i] {
			node := "(to:" + strconv.Itoa((*edge).Other(i)) + ",wt:" + strconv.FormatFloat((*edge).Weight().(float64), 'f', 6, 64) + ")"
			str = append(str, node)
		}
		printstr := strings.Join(str, " ")
		log.Print("vertex " + strconv.Itoa(i) + ": " + printstr)
		str = str[:0]
	}
}
