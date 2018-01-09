package weightgraph

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//稠密图 - 邻接矩阵

type DenseGraph struct {
	n, m     int  // n表示图中节点个数， m表示几点之间连接的个数
	directed bool // 是否为有向图
	g        [][]*Edge
}

func NewDGraph(n int, directed bool) *DenseGraph {
	var g [][]*Edge = make([][]*Edge, n)
	for k := range g {
		g[k] = make([]*Edge, n)
	}
	return &DenseGraph{
		n:        n,
		m:        0,
		g:        g,
		directed: directed,
	}
}

//节点数
func (dg *DenseGraph) V() int {
	return dg.n
}

//边的个数
func (dg *DenseGraph) E() int {
	return dg.m
}

//添加边
func (dg *DenseGraph) AddEdge(v, w int, weight Weight) {
	if v < 0 || v > dg.n-1 || w < 0 || w > dg.n-1 {
		return
	}
	if dg.HasEdge(v, w) {
		dg.g[v][w] = nil
		if v != w && !dg.directed {
			dg.g[w][v] = nil
		}
		dg.m--
	}
	dg.g[v][w] = NewEdge(v, w, weight)
	if v != w && !dg.directed {
		dg.g[w][v] = NewEdge(w, v, weight)
	}
	dg.m++
}

//是否有v->w的边
func (dg *DenseGraph) HasEdge(v, w int) bool {
	if v < 0 || v > dg.n-1 || w < 0 || w > dg.n-1 {
		return false
	}
	return dg.g[v][w] == nil
}

//遍历临边
func (dg *DenseGraph) AdjIterator(v int) []int {
	//临边
	var limb []int = make([]int, 0)
	for k, v := range dg.g[v] {
		if v != nil {
			limb = append(limb, k)
		}
	}
	return limb
}

func (dg *DenseGraph) show() {
	var header string
	for k := 0; k < dg.V(); k++ {
		header += "  " + strconv.Itoa(k)
	}
	log.Print(header)
	for i := 0; i < dg.V(); i++ {
		var lineWeight []float64 = make([]float64, 0)
		for j := 0; j < dg.V(); j++ {
			edge := dg.g[i][j]
			var weight float64
			if edge != nil {
				weight = (*edge).Weight().(float64)
			} else {
				weight = 0.0
			}
			lineWeight = append(lineWeight, weight)
		}
		//打印输出
		str := strconv.Itoa(i) + " " + strings.Trim(fmt.Sprint(lineWeight), "[]")
		log.Print(str)
		lineWeight = lineWeight[:0]
	}
}
