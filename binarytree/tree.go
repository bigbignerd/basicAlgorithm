package binarytree

import (
	"container/list"
	"log"
)

//二分搜索树的定义：每个节点的键值都大于它的左孩子，小于它的右孩子，并且以左右节点为根
//的子树仍为二分搜索树
//二分搜索树优势：高效；查找、插入、删除元素的的时间复杂度都可以保证在O(logn)级别

type node struct {
	key   string
	value int
	left  *node
	right *node
}
type BST struct {
	root  *node
	count int
}

func (b *BST) Size() int {
	return b.count
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

//向二叉搜索树种插入新的元素
func (b *BST) Insert(key string, value int) {
	b.root = b.insertNode(b.root, key, value)
}
func (b *BST) insertNode(n *node, key string, value int) *node {
	//判断元素最终插入位置 递归结束条件
	if n == nil {
		return NewNode(key, value)
	}
	//判断待插入元素与当前节点value的大小
	if (*n).key == key {
		(*n).value = value
	} else if (*n).key < key {
		//插入右子树
		(*n).right = b.insertNode((*n).right, key, value)
	} else {
		(*n).left = b.insertNode((*n).left, key, value)
	}
	return n
}

//判断二叉搜索树种是否包含某一个key
func (b *BST) Contain(key string) bool {
	return b.containKey(b.root, key)
}
func (b *BST) containKey(n *node, key string) bool {
	//节点为nil 则无可比较的点
	if n == nil {
		return false
	}
	if key < (*n).key {
		return b.containKey((*n).left, key)
	} else if key > (*n).key {
		return b.containKey((*n).right, key)
	} else {
		// == key 找到了key
		return true
	}
}

//二叉搜索树种搜索key 对应的value
func (b *BST) Search(key string) *int {
	return b.searchKey(b.root, key)
}
func (b *BST) searchKey(n *node, key string) *int {
	if n == nil {
		return nil
	}
	if key < (*n).key {
		return b.searchKey((*n).left, key)
	} else if key > (*n).key {
		return b.searchKey((*n).right, key)
	} else {
		return &(*n).value
	}
}

//深度优先的三种遍历方式：
//二叉树前序遍历
func (b *BST) Preorder() {
	b.preorder(b.root)
}

//二叉树中序遍历
func (b *BST) Inorder() {
	b.indorder(b.root)
}

//二叉树后续遍历
func (b *BST) Postorder() {
	b.postorder(b.root)
}

//广度优先 层序遍历 需要使用队列辅助
func (b *BST) Levelorder() {
	l := list.New()
	//先将根节点入队
	l.PushBack(b.root)
	//出队遍历
	for e := l.Front(); e != nil; {
		l.Remove(e) //出队
		n := e.Value
		log.Printf("%v ", (*n).key)
		if (*n).left != nil {
			l.PushBack((*n).left)
		}
		if (*n).right != nil {
			l.PushBack((*n).right)
		}
	}
}
func (b *BST) preorder(n *node) {
	if n != nil {
		log.Printf("%v ", (*n).key)
		b.preorder((*n).left)
		b.preorder((*n).right)
	}
}
func (b *BST) inorder(n *node) {
	if n != nil {
		b.preorder((*n).left)
		log.Printf("%v ", (*n).key)
		b.preorder((*n).right)
	}
}
func (b *BST) postorder(n *node) {
	if n != nil {
		b.preorder((*n).left)
		b.preorder((*n).right)
		log.Printf("%v ", (*n).key)
	}
}

//获取树中最小元素的key值
func (b *BST) Minimum() string {
	if b.count <= 0 {
		return ""
	}
	n := b.root
	for (*n).left != nil {
		n = (*n).left
	}
	return (*n).key
}

//获取树中最小元素的key值 【递归版本】
func (b *BST) Minimumr() string {
	minNode := b.minimumr(b.root)
	return (*minNode).key
}
func (b *BST) minimumr(n *node) *node {
	if (*n).left == nil {
		return n
	}
	return b.minimumr((*n).left)
}

//获取树中最大元素的key值
func (b *BST) Maximum() string {
	if b.count <= 0 {
		return ""
	}
	n := b.root
	for (*n).right != nil {
		n = (*n).right
	}
	return (*n).right
}

//获取树中最大元素的key值 【递归版本】
func (b *BST) Maximumr() string {
	maxNode := b.maximumr(b.root)
	return (*maxNode).key
}
func (b *BST) maximumr(n *node) *node {
	if (*n).right == nil {
		return n
	}
	return b.maximumr((*n).right)
}

//二叉树中删除最小元素
func (b *BST) RemoveMin() {
	if b.root != nil {
		b.root = b.removeMin(b.root)
	}
}

//二叉树中删除最大的元素
func (b *BST) RemoveMax() {
	if b.root != nil {
		b.root = b.removeMax(b.root)
	}
}

//递归找到最小元素删除
func (b *BST) removeMin(n *node) *node {
	if (*n).left == nil {
		b.count--
		return (*n).right
	}
	//未到最左子树元素 递归寻找
	(*n).left = b.removeMin((*n).left)
	return n
}

//递归找到最大元素删除
func (b *BST) removeMax(n *node) *node {
	if (*n).right == nil {
		b.count--
		return (*n).left
	}
	(*n).right = b.removeMax((*n).right)
	return n
}

//删除任意的节点：
//1.只有左子树->删除->n.left 指向左子树
//2.只有右子树->删除->n.right 指向右子树
//3.左右子树都存在，删除当前节点，当前节点有子树中的最小值替换到当前位置，并且将当前
//删除节点的左子树作为替换到当前位置的节点的left值
func (b *BST) Remove(key string) {
	b.root = b.remove(b.root, key)
}
func (b *BST) remove(n *node, key string) *node {
	if key < (*n).key {
		b.remove((*n).left, key)
	} else if key > (*n).key {
		b.remove((*n).right, key)
	} else { // key == n.key
		//左子树为空
		if (*n).left == nil {
			b.count--
			return (*n).right
		}
		if (*n).right == nil {
			b.count--
			return (*n).left
		}
		//左右子树都不为空 找到右子树中的最小值节点作为当前预删除节点的后继节点
		rightMinNode := CopyNode(b.minimumr((*n).right))
		//新创建了一个节点count+1
		b.count++
		//删除右子树中的这个最小节点
		(*rightMinNode).right = b.removeMin((*n).right)
		(*rightMinNode).left = (*n).left
		b.count--
		return rightMinNode
	}
}
func NewNode(key string, value int) *node {
	return &node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}
func CopyNode(*node) *node {
	return &node{
		key:   (*node).key,
		value: (*node).value,
		left:  (*node).left,
		right: (*node).right,
	}
}

//二叉树知识点延伸：
//1.解决二叉树退化成链表的情况：平衡二叉树（红黑树）
//2.二叉树与堆的结合：treap
//3.trie字典树
