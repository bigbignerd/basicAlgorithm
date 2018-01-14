package weightgraph

import (
	"github.com/bigbignerd/basicAlgorithm/heap"
	"log"
)

//单源最短路径 【无法处理负权边】

type Dijkstra struct {
	g      Graph
	s      int      //起点
	destTo []Weight //记录从源点到某一点的权值
	marked []bool
	from   []*Edge //点的前驱节点
}

func NewDijkstra(g Graph, s int) *Dijkstra {
	dk := &Dijkstra{
		g:      g,
		s:      s,
		destTo: make([]Weight, g.V()),
		marked: make([]bool, g.V()),
		from:   make([]*Edge, g.V()),
	}
	//创建一个最小索引堆
	minheap := heap.NewIndexMinHeap(g.V())
	//源点先入堆
	startNodeEdge := NewEdge(s, s, float64(0))
	dk.from[dk.s] = startNodeEdge
	dk.destTo[dk.s] = startNodeEdge.Weight()
	dk.marked[dk.s] = true
	//源点入堆
	minheap.Insert(0, startNodeEdge)
	//从最小堆中取数据 【松弛操作】
	for !minheap.IsEmpty() {
		//取出最小值索引
		v, err := minheap.ExtractMinIndex()
		if err != nil {
			log.Fatal(err)
		}
		dk.marked[v] = true
		//获取图中w的所有临边
		limb := dk.g.AdjIterator(v)
		for _, w := range limb {
			edge := dk.g.GetEdge(v, w)
			//判断是否已经记录v->w的路径，没有记录或者新的路径权值小于当前记录 则更新destTo
			if !dk.marked[w] {
				if dk.from[w] == nil || dk.destTo[v].(float64)+(*edge).Weight().(float64) < dk.destTo[w].(float64) {
					dk.destTo[w] = dk.destTo[v].(float64) + (*edge).Weight().(float64)
					dk.from[w] = edge
					//判断堆中是否已存在w 存在更新，不存在插入
					if minheap.Contain(w) {
						minheap.Change(w, edge)
					} else {
						minheap.Insert(w, edge)
					}
				}
			}
		}
	}
	return dk
}

//最短路径权值
func (dk *Dijkstra) ShortestPathTo(w int) float64 {
	if w < 0 || w >= dk.g.V() {
		log.Fatal("w越界.")
	}
	if dk.HasPath(w) {
		return dk.destTo[w].(float64)
	}
	return -1
}

//是否有到w的路径
func (dk *Dijkstra) HasPath(w int) bool {
	if w < 0 || w >= dk.g.V() {
		log.Fatal("w index out of range.")
	}
	return dk.marked[w]
}

//最短路径
func (dk *Dijkstra) ShortestPath(w int, path *[]*Edge) {
	if w < 0 || w >= dk.g.V() {
		log.Fatal("w index out of range.")
	}
	if !dk.HasPath(w) {
		log.Fatal("无此最短路径")
	}
	stack := make([]*Edge, 0)
	edge := dk.from[w]
	for (*edge).V() != dk.s {
		stack = append(stack, edge)
		edge = dk.from[(*edge).V()]
	}
	//起点
	stack = append(stack, edge)
	//调整为v->w的顺序
	for i := len(stack) - 1; i >= 0; i-- {
		*path = append(*path, stack[i])
	}
}

//打印路径
func (dk *Dijkstra) Show(w int) {
	if w < 0 || w >= dk.g.V() {
		log.Fatal("w index out of range.")
	}
	if !dk.HasPath(w) {
		log.Fatal("无此最短路径")
	}
	path := make([]*Edge, 0)
	dk.ShortestPath(w, &path)
	for _, edge := range path {
		log.Printf("%d->%d", (*edge).V(), (*edge).W())
	}
}
