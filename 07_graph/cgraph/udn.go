package cgraph

// undigragh net 无向网 udn
// 无向网邻接矩阵存储方式（二维数组）
// 将w置为1，矩阵初始化值置为0，则为“图”
// 去掉g.arcs[j][i] = g.arcs[i][j]则为“有向”

import (
	"fmt"
	"math"
)

func NewUND(v, a int) *AMGraph {
	arcs := make([][]int, MVNum)
	for i := range arcs {
		arcs[i] = make([]int, MVNum)
	}
	return &AMGraph{
		vexs:   make([]string, MVNum),
		arcs:   arcs,
		vexnum: v,
		arcnum: a,
	}
}

func (g *AMGraph) LocateVex(v string) int {
	for i := 0; i < g.vexnum; i++ {
		if v == g.vexs[i] {
			return i
		}
	}
	return -1
}

func (g *AMGraph) PrintG() {
	for _, row := range g.arcs {
		fmt.Println(row)
	}
}

func (g *AMGraph) CreateUDN() {
	fmt.Printf("输入总顶点数与总边数（vexnum arcnum）:")
	fmt.Scanf("%d %d", &g.vexnum, &g.arcnum)
	fmt.Println("依次输入点的信息：")
	const maxInt = math.MaxInt64
	for i := 0; i < g.vexnum; i++ {
		fmt.Scanf("%s", &g.vexs[i])
	}
	for i := 0; i < g.vexnum; i++ {
		for j := 0; j < g.vexnum; j++ {
			g.arcs[i][j] = maxInt
		}
	}
	fmt.Println("输入一条边所依附的顶点以及边的权值（v1 v2 w）:")
	for k := 0; k < g.arcnum; k++ {
		v1, v2, w := "", "", 0
		fmt.Scanf("%s %s %d", &v1, &v2, &w)
		i := g.LocateVex(v1)
		j := g.LocateVex(v2)
		g.arcs[i][j] = w
		g.arcs[j][i] = g.arcs[i][j]
	}
}

func (g *AMGraph) visit(v int) {
	fmt.Println(v)
}

func (g *AMGraph) InitVisited() {
	visited = make([]bool, MVNum)
}

// 邻接矩阵的深度优先遍历
func (g *AMGraph) DFS(v int) {
	g.visit(v)
	visited[v] = true
	for w := 0; w < g.vexnum; w++ {
		if g.arcs[v][w] != 0 && !visited[w] {
			g.DFS(w)
		}
	}
}
