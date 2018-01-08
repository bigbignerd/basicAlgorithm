package graph

//稠密图 - 邻接矩阵

type DenseGraph struct {
	n, m     int  // n表示图中节点个数， m表示几点之间连接的个数
	directed bool // 是否为有向图
	g        [][]bool
}

func NewDGraph(n int, directed bool) *DenseGraph {
	var g [][]bool = make([][]bool, n)
	for k := range g {
		g[k] = make([]bool, n)
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
func (dg *DenseGraph) AddEdge(v, w int) {
	if v < 0 || v > dg.n-1 || w < 0 || w > dg.n-1 {
		return
	}
	if dg.HasEdge(v, w) {
		return
	}
	dg.g[v][w] = true
	if !dg.directed {
		dg.g[w][v] = true
	}
	dg.m++
}

//是否有v->w的边
func (dg *DenseGraph) HasEdge(v, w int) bool {
	if v < 0 || v > dg.n-1 || w < 0 || w > dg.n-1 {
		return false
	}
	return dg.g[v][w]
}

//遍历临边
func (dg *DenseGraph) AdjIterator(v int) []int {
	//临边
	var limb []int = make([]int, 0)
	for k, v := range dg.g[v] {
		if v == true {
			limb = append(limb, k)
		}
	}
	return limb
}
