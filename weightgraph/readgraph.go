package weightgraph

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
	V, E, _, err = readNewLine(br)
	if V != g.V() {
		return errors.New("读取的数据v与初始化的v值不一致")
	}
	for i := 0; i < E; i++ {
		e1, e2, weight, err := readNewLine(br)
		if err != nil {
			break
		}
		g.AddEdge(e1, e2, weight)
	}
	return nil
}
func readNewLine(br *bufio.Reader) (int, int, Weight, error) {
	line, _, err := br.ReadLine()
	if err == io.EOF {
		return -1, -1, -1, err
	}
	newLine := strings.TrimSpace(string(line))
	element := strings.Split(newLine, " ")
	//这里需要对_ err 进行判断和处理 防止索引越界
	e1, _ := strconv.Atoi(element[0])
	e2, _ := strconv.Atoi(element[1])
	weight, _ := strconv.ParseFloat(element[2], 64)
	return e1, e2, weight, nil
}
