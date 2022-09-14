package main

import (
	"guochao.com/datastruct/07_graph/cgraph"
)

func main() {
	// n := cgraph.NewUND(5, 6)
	// n.CreateUDN()
	// n.PrintG()

	g := cgraph.NewALGraph(5, 6)
	g.CreateUDG()
	g.PrintG()

}
