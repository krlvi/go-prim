# Go-prim

Finds the Minimum spanning tree (MST) of a grpah.

# Usage

1) Have Go installed

3) Have an undirected graph in a dot/graphviz format

4) For Generating MST do `go run main.go <mygraph.dot>` which outputs a new graph T to standard out.

If you wish to visualize the output:

5) Have Graphviz installed

6) Do a `go run main.go <mygraph.dot> | dot -Tpdf > mst.pdf`

# Example 

`go run main.go input.dot | dot -Tpdf > mst.pdf`
