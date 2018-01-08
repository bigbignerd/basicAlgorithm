package unionfind

//go实现的并查集与c++版本的执行时间差异有点大
type UnionFind struct {
	id []int
	sz []int //存储各个集合中节点的个数
	// rank  []int //使用rank与size类似，不过rank是记录当前这个节点的层数
	count int
}

func (u *UnionFind) Find(p int) int {
	return u.id[p]
}
func (u *UnionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
func (u *UnionFind) UnionEle(p, q int) {
	pid := u.Find(p)
	qid := u.Find(q)
	if qid == pid {
		return
	}
	for i := 0; i < u.count; i++ {
		if u.id[i] == pid {
			u.id[i] = qid
		}
	}
}
func NewUnion(count int) *UnionFind {
	var id []int = make([]int, count)
	var sz []int = make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
		sz[i] = 1 //各节点集合元素个数初始为1
	}
	return &UnionFind{
		id:    id,
		sz:    sz,
		count: count,
	}
}

//v2版本：上面的版本在10万数量级的数据执行时间已经过长了
//优化方式：每一个节点对应数组中的一个元素k1=>v1, 然后v1=k2 也就是说k2是k1的父节点
//也就是k1 k2是连通的

func (u *UnionFind) Findv2(p int) int {
	//严格处理需要返回error
	if p < 0 || p > u.count-1 {
		return -1
	}
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}
func (u *UnionFind) IsConnectedv2(p, q int) bool {
	return u.Findv2(p) == u.Findv2(q)
}
func (u *UnionFind) UnionElev2(p, q int) {
	proot := u.Findv2(p)
	qroot := u.Findv2(q)
	if proot == qroot {
		return
	}
	u.id[proot] = qroot
}

//v3版本 因为v2版本的union element 过程没有分别对两个集合中节点的个数判断
//单纯的将proot指向了qroot,所以这就可能造成一个集合树的高度非常高，造成查询
//时迭代的次数增多
func (u *UnionFind) UnionElev3(p, q int) {
	proot := u.Findv2(p)
	qroot := u.Findv2(q)
	if proot == qroot {
		return
	}
	if u.sz[proot] < u.sz[qroot] {
		u.id[proot] = qroot
		u.sz[qroot] += u.sz[proot]
	} else {
		u.id[qroot] = proot
		u.sz[proot] += u.sz[qroot]
	}
}

//v4版本 路径压缩，就是在每次find的时候，判断当前是否为根节点，如果不是根节点
//将当前节点指向当前节点的父节点的父节点
func (u *UnionFind) Findv4(p int) int {
	if p < 0 || p > u.count-1 {
		return -1
	}
	for p != u.id[p] {
		//当前节点指向父亲的父亲
		u.id[p] = u.id[u.id[p]]
		p = u.id[p] //相当于跳过一个节点
	}
	return p
}
func (u *UnionFind) UnionElev4(p, q int) {
	proot := u.Findv4(p)
	qroot := u.Findv4(q)
	if proot == qroot {
		return
	}
	if u.sz[proot] < u.sz[qroot] {
		u.id[proot] = qroot
		u.sz[qroot] += u.sz[proot]
	} else {
		u.id[qroot] = proot
		u.sz[proot] += u.sz[qroot]
	}
}
