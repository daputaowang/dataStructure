package cgraph

const MVNum = 5

var visited []bool

// 图的邻接矩阵存储均表示（二维数组）
type AMGraph struct { // arc matrix
	vexs           []string // 顶点表
	arcs           [][]int  // 邻接矩阵
	vexnum, arcnum int      // 图的点前点数和边数
}

// 图的邻接表存储表示（链表数组）
type ArcNode struct { // 边结点
	adjvex  int // 该边所指向的顶点在顶点表中的位置
	nextarc *ArcNode
	info    string
}

type VNode struct { // 顶点信息
	data     string   // 顶点信息
	firstarc *ArcNode // 指向第一条依附该顶点的边的指针
}

type ALGraph struct { // arc list
	vertices       []VNode // 顶点表
	vexnum, arcnum int     // 图的当前顶点数和弧数
}
