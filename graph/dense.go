//稠密图-邻接矩阵
package graph

type DenseGraph struct {
	g        map[int]map[int]bool
	n, size  int
	directed bool
}

//NewSparseGraph 创建一个有向无向图
func NewDenseGraph(n int, directed bool) DenseGraph {
	g := map[int]map[int]bool{}
	return DenseGraph{
		n:        n,
		size:     0,
		directed: directed,
		g:        g,
	}
}

func (g *DenseGraph) V() int {
	return g.n
}
func (g *DenseGraph) E() int {
	return g.size
}

//AddEdge 添加一条关系
func (g *DenseGraph) AddEdge(v, w int) {
	if g.HasEdge(v, w) {
		return
	}
	g.g[v][w] = true
	if !g.directed {
		g.g[w][v] = true
	}
	g.size++
}

//HasEdge 检查是否存在该边关系
func (g *DenseGraph) HasEdge(v, w int) bool {
	return g.g[v][w]
}
