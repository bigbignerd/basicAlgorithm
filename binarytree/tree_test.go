package binarytree

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"os"
	"path/filepath"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	txtfolder, err := os.Getwd()
	if err != nil {
		t.Logf("get file path error:%v", err)
	}
	txtpath := filepath.Join(txtfolder, "bible.txt")
	words, err := util.ReadFile(txtpath)
	if err != nil {
		t.Log("parse txt fail.")
	}
	t.Logf("The total words:%d", len(words))
	//初始化二叉搜索树
	var bst *BST = &BST{}
	for _, v := range words {
		var res *int = bst.Search(v)
		if res == nil {
			bst.Insert(v, 1)
		} else {
			(*res)++
		}
	}
	//输出圣经中一些词出现的频率
	var searchWords []string = []string{"god", "and"}
	for _, v := range searchWords {
		if bst.Contain(v) {
			t.Logf("%v 在圣经中出现了 %d 次", v, *(bst.Search(v)))
		} else {
			t.Logf("%v not found", v)
		}
	}
}
func TestBf(t *testing.T) {

}
