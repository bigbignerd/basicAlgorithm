package unionfind

import (
	"github.com/bigbignerd/basicAlgorithm/util"
	"testing"
	"time"
)

func TestUnionFindv1(t *testing.T) {
	n := 100000
	uf := NewUnion(n)
	//测试执行时间
	start := time.Now()
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.UnionEle(a, b)
	}
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.IsConnected(a, b)
	}
	delta := time.Now().Sub(start)
	t.Logf("v1 execute time:%.6f", float64(delta)/1e9)
}

// v2时间大于v1 不知道原因
func TestUnionFindv2(t *testing.T) {
	n := 100000
	uf := NewUnion(n)
	//测试执行时间
	start := time.Now()
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.UnionElev2(a, b)
	}
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.IsConnectedv2(a, b)
	}
	delta := time.Now().Sub(start)
	t.Logf("v2 execute time:%.6f", float64(delta)/1e9)
}
func TestUnionFindv3(t *testing.T) {
	n := 100000
	uf := NewUnion(n)
	//测试执行时间
	start := time.Now()
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.UnionElev3(a, b)
	}
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.IsConnectedv2(a, b)
	}
	delta := time.Now().Sub(start)
	t.Logf("v3 execute time:%.6f", float64(delta)/1e9)
}
func TestUnionFindv4(t *testing.T) {
	n := 10
	uf := NewUnion(n)
	t.Log(uf.id)
	//测试执行时间
	start := time.Now()
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		t.Logf("a:%d,b:%d", a, b)
		uf.UnionElev4(a, b)
	}
	for i := 0; i < n; i++ {
		a := util.RandNumber(0, n-1)
		b := util.RandNumber(0, n-1)
		uf.IsConnectedv2(a, b)
	}
	delta := time.Now().Sub(start)
	t.Log(uf.id)

	t.Logf("v4 execute time:%.6f", float64(delta)/1e9)
}

func TestRandnumber(t *testing.T) {
	start := time.Now()
	n := 100000
	for i := 0; i < n; i++ {
		util.RandNumber(0, n-1)
	}
	delta := time.Now().Sub(start)
	t.Logf("t:%.6f", float64(delta)/1e9)
}
