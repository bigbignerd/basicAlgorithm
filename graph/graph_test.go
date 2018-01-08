package graph

import (
	"github.com/bigbignerd/basicAlgorithm/util"
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
	g := NewSGrahp(n, directed)

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
	filename := filepath.Join(curfolder, "testG1.txt")
	n := 13
	directed := false
	g := NewDGraph(n, directed)
	// t.Logf("%v", g)
	ReadGraph(g, filename)
	t.Log(g.g)
}
