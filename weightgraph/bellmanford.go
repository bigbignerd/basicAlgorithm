package weightgraph

import (
	"log"
)

//Bellman Ford算法：可以处理有负权边的图，但是图中不能存在负权环，否则无法找到最短路径
//核心思想：一个点到另一个点的最短路径最多经过V个顶点V-1条边
//对所有的点进行V-1次操作 【未理解透彻】
type BellmanFord struct {
	g              Graph
	s              int      //起点
	destTo         []Weight //记录从源点到某一点的权值
	from           []*Edge  //点的前驱节点
	hasNaviteCycle bool     //是否有负权环
}

func NewBellmanFord(g Graph, s int) *BellmanFord {
	bf := &BellmanFord{
		g:              g,
		s:              s,
		destTo:         make([]Weight, g.V()),
		from:           make([]*Edge, g.V()),
		hasNaviteCycle: false,
	}
	//Bellman ford 对每个点进行v-1此松弛操作
	bf.destTo[s] = NewEdge(s, s, float64(0))
	for pass := 1; pass < g.V(); pass++ {
		//松弛操作
		for i := 0; i < g.V(); i++ {
			//获取节点i的临边
			limb := bf.g.AdjIterator(i)
			for _, j := range limb {
				edge := bf.g.GetEdge(i, j)
				//如果没有访问过j或者通过新的edge 使得到达j的权值更短
				if bf.from[j] == nil || bf.destTo[i].(float64)+(*edge).Weight().(float64) < bf.destTo[j].(float64) {
					bf.destTo[j] = bf.destTo[i].(float64) + (*edge).Weight().(float64)
					bf.from[j] = edge
				}
			}
		}
	}
	bf.detectNegativeCycle()
	return bf
}
func (bf *BellmanFord) detectNegativeCycle() {
	for i := 0; i < bf.g.V(); i++ {
		//获取节点i的临边
		limb := bf.g.AdjIterator(i)
		for _, j := range limb {
			edge := bf.g.GetEdge(i, j)
			if bf.from[j] == nil || bf.destTo[i].(float64)+(*edge).Weight().(float64) < bf.destTo[j].(float64) {
				bf.hasNaviteCycle = true
			}
		}
	}
}

//最短路径权值
func (bf *BellmanFord) ShortestPathTo(w int) float64 {
	if w < 0 || w >= bf.g.V() {
		log.Fatal("w越界.")
	}
	if bf.HasPath(w) {
		return bf.destTo[w].(float64)
	}
	return -1
}

//是否有到w的路径
func (bf *BellmanFord) HasPath(w int) bool {
	if w < 0 || w >= bf.g.V() {
		log.Fatal("w index out of range.")
	}
	return bf.marked[w]
}

//最短路径
func (bf *BellmanFord) ShortestPath(w int, path *[]*Edge) {
	if w < 0 || w >= bf.g.V() {
		log.Fatal("w index out of range.")
	}
	if !bf.HasPath(w) {
		log.Fatal("无此最短路径")
	}
	stack := make([]*Edge, 0)
	edge := bf.from[w]
	for (*edge).V() != bf.s {
		stack = append(stack, edge)
		edge = bf.from[(*edge).V()]
	}
	//起点
	stack = append(stack, edge)
	//调整为v->w的顺序
	for i := len(stack) - 1; i >= 0; i-- {
		*path = append(*path, stack[i])
	}
}

//打印路径
func (bf *BellmanFord) Show(w int) {
	if w < 0 || w >= bf.g.V() {
		log.Fatal("w index out of range.")
	}
	if !bf.HasPath(w) {
		log.Fatal("无此最短路径")
	}
	path := make([]*Edge, 0)
	bf.ShortestPath(w, &path)
	for _, edge := range path {
		log.Printf("%d->%d", (*edge).V(), (*edge).W())
	}
}
