package graph

import (
	"bufio"
	"errors"
	"io"
	// "log"
	"os"
	"strconv"
	"strings"
)

func ReadGraph(g Graph, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	var V, E int
	V, E, err = readNewLine(br)
	if V != g.V() {
		return errors.New("读取的数据v与初始化的v值不一致")
	}
	for i := 0; i < E; i++ {
		e1, e2, err := readNewLine(br)
		if err != nil {
			break
		}
		g.AddEdge(e1, e2)
	}
	return nil
}
func readNewLine(br *bufio.Reader) (int, int, error) {
	line, _, err := br.ReadLine()
	if err == io.EOF {
		return -1, -1, err
	}
	newLine := strings.TrimSpace(string(line))
	element := strings.Split(newLine, " ")
	e1, _ := strconv.Atoi(element[0])
	e2, _ := strconv.Atoi(element[1])
	return e1, e2, nil
}
