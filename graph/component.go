package graph

import (
	"errors"
)

type Component struct {
	ccount  int //联通量个数
	g       Graph
	visited []bool
	id      []int //记录同一个联通分量中的元素
}

func NewComponent(g Graph) *Component {
	c := &Component{
		ccount:  0,
		g:       g,
		visited: make([]bool, g.V()),
		id:      make([]int, g.V()),
	}
	//初始化为-1
	for i := 0; i < len(c.id); i++ {
		c.id[i] = -1
	}
	//从每个节点出发开始进行深度优先遍历
	for i := 0; i < g.V(); i++ {
		if !c.visited[i] {
			c.dfs(i)
			c.ccount++
		}
	}
	return c
}
func (c *Component) dfs(i int) {
	c.visited[i] = true
	c.id[i] = c.ccount
	//获取i的所有临边
	limb := c.g.AdjIterator(i)
	for _, v := range limb {
		if !c.visited[v] {
			c.dfs(v)
		}
	}
}

//图中联通量的个数
func (c *Component) Count() int {
	return c.ccount
}
func (c *Component) IsConnected(v, w int) (bool, error) {
	if v < 0 || v >= c.g.V() || w < 0 || w >= c.g.V() {
		return false, errors.New("v|w越界")
	}
	return c.id[v] == c.id[w], nil
}
