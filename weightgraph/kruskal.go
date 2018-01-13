package weightgraph

import (
	"github.com/bigbignerd/basicAlgorithm/heap"
	"github.com/bigbignerd/basicAlgorithm/unionfind"
)

//Kruskal算法的思想：对所有的边进行排序，每次取出权值最小的边，如果选择了这条边后
//不会在最小生成树中形成环，那么这条边一定是最小生成树中的一条边
//算法实现：1>最小堆进行排序 2>并查集判断是否会形成环

type Kruskal struct {
	mst []*Edge
}

func NewKruskal(g Graph) *Kruskal {
	kruskal := &Kruskal{
		mst: make([]*Edge, 0),
	}
	//创建一个最小堆
	minheap := heap.NewMinHeap(g.E())
	//获取所有节点临边入堆
	for v := 0; v < g.V(); v++ {
		limb := g.AdjIterator(v)
		for _, w := range limb {
			//无向图 保证每条边只插入堆中一次 v > w 作为条件亦可
			if v < w {
				edge := g.GetEdge(v, w)
				minheap.Insert(edge)
			}
		}
	}
	//创建一个size为图节点个数的并查集
	uf := unionfind.NewUnion(g.V())
	//从堆中取边
	for !minheap.IsEmpty() {
		edge := minheap.ExtractMin().(*Edge)
		//判断添加此边是否构成环
		if uf.IsConnectedv4((*edge).V(), (*edge).W()) {
			continue
		}
		kruskal.mst = append(kruskal.mst, edge)
		uf.UnionElev4((*edge).V(), (*edge).W())
	}
	return kruskal
}

func (k *Kruskal) MstEdges() []*Edge {
	return k.mst
}
