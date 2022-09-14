package cgraph

import "fmt"

// undirected graph 无向图
// 无向图邻接表存储方式（链表数组）
/* 有向图的邻接表表存储方式中，只存出度，所以不需要处理顶点j，也就是不需要下面这两行
p2 := &ArcNode{i, g.vertices[j].firstarc, ""}
g.vertices[j].firstarc = p2 */
/* 有向图的逆邻接表存储方式中，只存入度，所以不需要处理顶点i，也就是不需要下面这两行
p1 := &ArcNode{j, g.vertices[i].firstarc, ""}
g.vertices[i].firstarc = p1 */
// 网的邻接表存储方式，只需要在上述基础上将ArcNode的info置为权重即可

func NewALGraph(v, a int) *ALGraph {
	return &ALGraph{
		vertices: make([]VNode, MVNum),
		vexnum:   v,
		arcnum:   a,
	}
}

func (g *ALGraph) LocateVex(v string) int {
	for i := 0; i < g.vexnum; i++ {
		if v == g.vertices[i].data {
			return i
		}
	}
	return -1
}

func (g *ALGraph) PrintG() {
	for i, node := range g.vertices {
		tmp := fmt.Sprintf("%d|%s:", i, node.data)
		p := node.firstarc
		for p != nil {
			tmp = fmt.Sprintf("%s->%d", tmp, p.adjvex)
			p = p.nextarc
		}
		fmt.Println(tmp)
	}
}

func (g *ALGraph) CreateUDG() {
	fmt.Printf("输入总顶点数与总边数（vexnum arcnum）:")
	fmt.Scanf("%d %d", &g.vexnum, &g.arcnum) // 输入总顶点数，总边数
	fmt.Println("依次输入顶点的值：")
	for i := 0; i < g.vexnum; i++ { // 输入各点构造表头结点表
		fmt.Scanf("%s", &g.vertices[i].data) // 输入顶点值
		g.vertices[i].firstarc = nil         // 初始化表头结点的指针域，golang可以省略
	}
	fmt.Println("输入一条边所依附的顶点（v1 v2）:")
	for k := 0; k < g.arcnum; k++ { // 输入各边，构造邻接表
		v1, v2 := "", ""
		fmt.Scanf("%s %s", &v1, &v2) // 输入边的两个顶点的数据域
		i := g.LocateVex(v1)
		j := g.LocateVex(v2)
		p1 := &ArcNode{j, g.vertices[i].firstarc, ""} // 无向图中i的邻接点序号为j，采用头插法
		g.vertices[i].firstarc = p1
		p2 := &ArcNode{i, g.vertices[j].firstarc, ""} // 无向图中j的邻接点序号为i，采用头插法
		g.vertices[j].firstarc = p2
	}
}

func (g *ALGraph) visit(v int) {
	fmt.Println(v)
}

func (g *ALGraph) InitVisited() {
	visited = make([]bool, MVNum)
}

// 邻接表的深度优先遍历
func (g *ALGraph) DFS(v int) {
	g.visit(v)
	visited[v] = true
	p := g.vertices[v].firstarc
	for p != nil {
		if !visited[p.adjvex] {
			g.DFS(p.adjvex)
		}
		p = p.nextarc
	}
}
