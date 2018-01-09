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
