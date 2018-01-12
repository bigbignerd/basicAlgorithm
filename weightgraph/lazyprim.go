package weightgraph

import (
	"github.com/bigbignerd/basicAlgorithm/heap"
	// "log"
)

//Minimum span tree 最小生成树

type LazyPrimMST struct {
	g         Graph
	pq        *heap.MinHeap
	marked    []bool
	mst       []*Edge //最小生成树的边:g.V()-1
	mstWeight Weight
}

func NewLazyPrim(g Graph) *LazyPrimMST {
	minheap := heap.NewMinHeap(g.E())
	lazyprim := &LazyPrimMST{
		g:         g,
		pq:        minheap,
		marked:    make([]bool, g.V()),
		mst:       make([]*Edge, 0),
		mstWeight: 0,
	}

	//lazy prim
	lazyprim.visit(0)
	for !lazyprim.pq.IsEmpty() {
		//最小堆中获取最小 权值边
		edge := lazyprim.pq.ExtractMin().(*Edge)
		// 如果这条边的两端都已经访问过了, 则扔掉这条边
		if lazyprim.marked[(*edge).V()] == lazyprim.marked[(*edge).W()] {
			continue
		}
		// 否则, 这条边则应该存在在最小生成树中
		lazyprim.mst = append(lazyprim.mst, edge)
		// 访问和这条边连接的还没有被访问过的节点
		if !lazyprim.marked[(*edge).V()] {
			lazyprim.visit((*edge).V())
		} else {
			lazyprim.visit((*edge).W())
		}
	}
	//计算最小生成树的权值
	// for _, edge := range lazyprim.mst {
	// 	lazyprim.mstWeight += (*edge).weight
	// }
	return lazyprim
}

//返回最小生成树的所有边
func (lp *LazyPrimMST) MstEdges() []*Edge {
	return lp.mst
}

// 返回最小生成树的权值
func (lp *LazyPrimMST) Result() Weight {
	return lp.mstWeight
}
func (lp *LazyPrimMST) visit(v int) {
	//已经访问过了直接返回
	if lp.marked[v] {
		return
	}
	lp.marked[v] = true
	//获取当前节点v的所有相邻节点
	limb := lp.g.AdjIterator(v)
	for _, value := range limb {
		edge := lp.g.GetEdge(v, value)
		w := (*edge).Other(v)
		if !lp.marked[w] {
			//v->w的临边入堆
			lp.pq.Insert(edge)
		}
	}
}
