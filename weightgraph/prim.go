package weightgraph

import (
	"github.com/bigbignerd/basicAlgorithm/heap"
	"log"
)

type PrimMst struct {
	g      Graph
	h      *heap.IndexMinHeap
	edgeTo []*Edge
	marked []bool
	mst    []*Edge
}

func NewPrimMst(g Graph) *PrimMst {
	indexminheap := heap.NewIndexMinHeap(g.V())
	pmst := &PrimMst{
		g:      g,
		h:      indexminheap,
		edgeTo: make([]*Edge, g.V()),
		marked: make([]bool, g.V()),
		mst:    make([]*Edge, 0),
	}
	pmst.visit(0)
	//最小索引堆中获取权值最小的边
	for !pmst.h.IsEmpty() {
		v, err := pmst.h.ExtractMinIndex()
		if err != nil {
			log.Fatal("get min index error.")
		}
		if pmst.edgeTo[v] == nil {
			log.Fatal("edge to not contain v.")
		}
		pmst.mst = append(pmst.mst, pmst.edgeTo[v])
		pmst.visit(v)
	}
	return pmst
}

func (pmst *PrimMst) visit(v int) {
	if pmst.marked[v] {
		return
	}
	pmst.marked[v] = true
	//获取相邻节点
	limb := pmst.g.AdjIterator(v)
	for _, w := range limb {
		edge := pmst.g.GetEdge(v, w)
		//如果边的另一个端点w未被访问
		if !pmst.marked[w] {
			if pmst.edgeTo[w] == nil {
				pmst.edgeTo[w] = edge
				pmst.h.Insert(w, edge)
			} else if pmst.edgeTo[w] != nil && edge.Weight().(float64) < pmst.edgeTo[w].Weight().(float64) {
				//访问过这个节点，但是新的这个边 权值更小
				pmst.edgeTo[w] = edge
				//更新最小堆中的w的值
				pmst.h.Change(w, edge)
			}
		}
	}
}

func (pmst *PrimMst) MstEdges() []*Edge {
	return pmst.mst
}
