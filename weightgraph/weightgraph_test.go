package weightgraph

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestReadGraph(t *testing.T) {
	curfolder, err := os.Getwd()
	if err != nil {
		log.Fatal("get file path error.")
	}
	filename1 := filepath.Join(curfolder, "testG1.txt")
	n := 8
	directed := false
	// g1 := NewDGraph(n, directed)
	g2 := NewSGraph(n, directed)
	ReadGraph(g2, filename1)
	g2.show()
}
func readGraph() Graph {
	curfolder, err := os.Getwd()
	if err != nil {
		log.Fatal("get file path error.")
	}
	filename1 := filepath.Join(curfolder, "testG1.txt")
	n := 8
	directed := false
	// g1 := NewDGraph(n, directed)
	g2 := NewSGraph(n, directed)
	ReadGraph(g2, filename1)
	return g2
}
func TestLazyPrimMST(t *testing.T) {
	g := readGraph()
	mst := NewLazyPrim(g)
	for _, v := range mst.MstEdges() {
		t.Logf("from:%d to:%d weight:%f", (*v).V(), (*v).W(), (*v).Weight())
	}
}

func TestPrimMST(t *testing.T) {
	g := readGraph()
	mst := NewPrimMst(g)
	for _, v := range mst.MstEdges() {
		t.Logf("from:%d to:%d weight:%f", (*v).V(), (*v).W(), (*v).Weight())
	}
}
