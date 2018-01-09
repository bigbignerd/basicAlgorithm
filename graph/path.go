package graph

import (
	"fmt"
	"log"
	"strings"
)

type Path struct {
	g       Graph
	s       int
	visited []bool
	from    []int
}

func NewPath(g Graph, s int) *Path {
	p := &Path{
		g:       g,
		s:       s,
		visited: make([]bool, g.V()),
		from:    make([]int, g.V()),
	}
	for i := 0; i < g.V(); i++ {
		p.from[i] = -1
	}
	//寻找路径
	p.dfs(p.s)
	return p
}

func (p *Path) dfs(v int) {
	p.visited[v] = true
	limb := p.g.AdjIterator(v)
	for _, value := range limb {
		if !p.visited[value] {
			p.from[value] = v
			p.dfs(value)
		}
	}
}

//是否有从v->w的路径
func (p *Path) HasPath(w int) bool {
	return p.visited[w]
}

//获取从v->w的路径
func (p *Path) GetPath(w int, path *[]int) {
	var stack []int = make([]int, 0)
	node := w
	for node != -1 {
		stack = append(stack, node)
		node = p.from[node]
	}
	for i := len(stack) - 1; i >= 0; i-- {
		*path = append(*path, stack[i])
	}
}

//打印路径v->w的路径
func (p *Path) ShowPath(w int) {
	path := make([]int, 0)
	p.GetPath(w, &path)
	pathstring := strings.Trim(strings.Replace(fmt.Sprint(path), " ", "->", -1), "[]")
	log.Printf("The path of %d to %d is : %s", p.s, w, pathstring)
}
