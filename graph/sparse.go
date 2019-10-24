//稀疏图-邻接表
package graph

import (
	"github.com/brewlin/go-stl/tree/rb"
)

type SparseGraph struct {
	g        map[int]*rb.Set
	n, size  int
	directed bool
}

//NewSparseGraph 创建一个有向无向图
func NewSparseGraph(n int, directed bool) SparseGraph {
	g := map[int]*rb.Set{}
	return SparseGraph{
		n:        n,
		size:     0,
		directed: directed,
		g:        g,
	}
}

func (g *SparseGraph) V() int {
	return g.n
}
func (g *SparseGraph) E() int {
	return g.size
}

//AddEdge 添加一条关系
func (g *SparseGraph) AddEdge(v, w int) {
	g.g[v].Add(w)
	//如果是无向图，则需要记录相反关系
	if v != w && !g.directed {
		g.g[w].Add(v)
	}
	g.size++
}

//HasEdge 检查是否存在该边关系
func (g *SparseGraph) HasEdge(v, w int) bool {
	return g.g[v].Has(w)
}
