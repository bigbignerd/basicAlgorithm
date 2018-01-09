package graph

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDesnseGraph(t *testing.T) {
	n := 10
	m := 20
	directed := false
	g := NewDGraph(n, directed)

	for i := 0; i < m; i++ {
		v := util.RandNumber(0, n-1)
		w := util.RandNumber(0, n-1)
		g.AddEdge(v, w)
	}
	for i := 0; i < n; i++ {
		limb := g.AdjIterator(i)
		t.Logf("%d 的临边：%v", i, limb)
	}
}
func TestSparseGraph(t *testing.T) {
	n := 10
	m := 20
	directed := false
	g := NewSGraph(n, directed)

	for i := 0; i < m; i++ {
		v := util.RandNumber(0, n-1)
		w := util.RandNumber(0, n-1)
		g.AddEdge(v, w)
	}
	for i := 0; i < n; i++ {
		limb := g.AdjIterator(i)
		t.Logf("%d 的临边：%v", i, limb)
	}
}
func TestReadGraph(t *testing.T) {
	curfolder, err := os.Getwd()
	if err != nil {
		t.Log("get file path error.")
	}
	filename1 := filepath.Join(curfolder, "testG1.txt")
	filename2 := filepath.Join(curfolder, "testG2.txt")
	n := 13
	n2 := 6
	directed := false
	g1 := NewDGraph(n, directed)
	g2 := NewSGraph(n2, directed)
	// t.Logf("%v", g)
	ReadGraph(g1, filename1)
	ReadGraph(g2, filename2)
	t.Log(g1.g)
	t.Log(g2.g)
}
func readgraph() (Graph, Graph) {
	curfolder, err := os.Getwd()
	if err != nil {
		log.Fatal("get file path error.")
	}
	filename1 := filepath.Join(curfolder, "testG1.txt")
	filename2 := filepath.Join(curfolder, "testG2.txt")
	n := 13
	n2 := 7
	directed := false
	g1 := NewDGraph(n, directed)
	g2 := NewSGraph(n2, directed)
	// t.Logf("%v", g)
	ReadGraph(g1, filename1)
	ReadGraph(g2, filename2)
	return g1, g2
}

//测试深度优先dfs遍历
func TestIteratorGraph(t *testing.T) {
	g1, g2 := readgraph()
	c1 := NewComponent(g1)
	c2 := NewComponent(g2)
	t.Logf("c1联通量个数:%d", c1.Count())
	t.Logf("c2联通量个数:%d", c2.Count())
	isconnected1, _ := c1.IsConnected(4, 3)
	isconnected2, _ := c2.IsConnected(0, 6)
	t.Logf("c1 4与3是否联通：%v", isconnected1)
	t.Logf("c2 3与5是否联通：%v", isconnected2)

}

func TestShowPath(t *testing.T) {
	_, g2 := readgraph()
	s := 0
	p1 := NewPath(g2, s)
	//最短路径
	p2 := NewShortestPath(g2, s)
	// t.Log(p1.HasPath(10))
	//path 9->10
	p1.ShowPath(6)
	p2.ShowPath(6)
	t.Log(p2.Length(6))
}
