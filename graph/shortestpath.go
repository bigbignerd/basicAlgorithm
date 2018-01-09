package graph

import (
	"container/list"
	"fmt"
	"log"
	"strings"
)

type ShortestPath struct {
	g       Graph
	s       int //起点
	visited []bool
	from    []int
	order   []int //路径长度
}

func NewShortestPath(g Graph, s int) *ShortestPath {
	sp := &ShortestPath{
		g:       g,
		s:       s,
		visited: make([]bool, g.V()),
		from:    make([]int, g.V()),
		order:   make([]int, g.V()),
	}
	//初始化from order值为-1
	for i := 0; i < g.V(); i++ {
		sp.from[i] = -1
		sp.order[i] = -1
	}

	sp.bfs()
	return sp
}

//广度优先遍历
func (sp *ShortestPath) bfs() {

	//模拟一个队列
	l := list.New()
	l.PushBack(sp.s)        //起点入队
	sp.visited[sp.s] = true //记录当前元素已入队
	sp.order[sp.s] = 0      //记录开始节点到当前节点距离
	for l.Len() > 0 {
		e := l.Front()
		var element int = e.Value.(int)
		l.Remove(e)
		limb := sp.g.AdjIterator(element) //获取当前出队节点的临边
		for _, v := range limb {
			if !sp.visited[v] {
				l.PushBack(v) //临边入队
				sp.from[v] = element
				sp.visited[v] = true
				sp.order[v] = sp.order[element] + 1
			}
		}
	}
}

//是否有从v->w的路径
func (sp *ShortestPath) HasPath(w int) bool {
	return sp.visited[w]
}

//获取从v->w的路径
func (sp *ShortestPath) GetPath(w int, path *[]int) {
	var stack []int = make([]int, 0)
	node := w
	for node != -1 {
		stack = append(stack, node)
		node = sp.from[node]
	}
	for i := len(stack) - 1; i >= 0; i-- {
		*path = append(*path, stack[i])
	}
}

//打印路径v->w的路径
func (sp *ShortestPath) ShowPath(w int) {
	path := make([]int, 0)
	sp.GetPath(w, &path)
	pathstring := strings.Trim(strings.Replace(fmt.Sprint(path), " ", "->", -1), "[]")
	log.Printf("The path of %d to %d is : %s", sp.s, w, pathstring)
}

//返回从v->w的路径长度
func (sp *ShortestPath) Length(w int) int {
	return sp.order[w]
}
