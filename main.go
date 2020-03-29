package main

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"os"
	"strconv"
	"strings"
)

func main() {
	out, err := work()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func work() (string, error) {
	astGraph, err := readDotFile(os.Args[1])
	if err != nil {
		return "", err
	}
	source, nodes, err := parseUnidirectedGraph(astGraph)
	if err != nil {
		return "", err
	}
	dest := simple.NewWeightedUndirectedGraph(0, 0)
	path.Prim(dest, source)

	edges := dest.WeightedEdges()

	var sb strings.Builder

	sb.WriteString("graph T {\n")
	sb.WriteString("\trankdir=\"LR\";\n")
	for _, n := range nodes {
		sb.WriteString(fmt.Sprint("\t", n, ";\n"))
	}
	for edges.Next() {
		e := edges.WeightedEdge()
		sb.WriteString(fmt.Sprint(
			"\t",
			nodes[e.From().ID()], " -- ", nodes[e.To().ID()],
			fmt.Sprintf("[weight=%d,\tlabel=%d]", int(e.Weight()), int(e.Weight())),
			";\n"))
	}
	sb.WriteString("}\n")
	return sb.String(), nil
}

func readDotFile(inputPath string) (*ast.Graph, error) {
	file, err := dot.ParseFile(inputPath)
	if err != nil {
		return nil, err
	}
	if len(file.Graphs) != 1 {
		return nil, errors.New("invalid dot file")
	}
	return file.Graphs[0], nil
}

func parseUnidirectedGraph(astGraph *ast.Graph) (*simple.WeightedUndirectedGraph, map[int64]string, error) {
	nodes := map[string]graph.Node{}
	intStringNodes := map[int64]string{}
	g := simple.NewWeightedUndirectedGraph(0, 0)
	for _, s := range astGraph.Stmts {
		switch s.(type) {
		case *ast.NodeStmt:
			n := s.(*ast.NodeStmt)
			nodes[n.String()] = g.NewNode()
			intStringNodes[nodes[n.String()].ID()] = n.String()
			g.AddNode(nodes[n.String()])
		case *ast.EdgeStmt:
			e := s.(*ast.EdgeStmt)
			weight, err := strconv.ParseFloat(e.Attrs[0].Val, 64)
			if err != nil {
				return nil, nil, err
			}
			g.SetWeightedEdge(g.NewWeightedEdge(nodes[e.From.String()], nodes[e.To.Vertex.String()], weight))
		}
	}
	return g, intStringNodes, nil
}
