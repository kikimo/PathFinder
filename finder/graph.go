package finder

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	From    int64
	To      int64
	Reverse bool
	Parent  *Edge
}

func (e *Edge) Equal(o *Edge) bool {
	if e.Reverse == o.Reverse {
		return e.From == o.From && e.To == o.To
	}
	return e.From == o.To && e.To == o.From
}

func (e *Edge) IsDuplicate(o *Edge) bool {
	for p := e; p != nil; p = p.Parent {
		if p.Equal(o) {
			return true
		}
	}
	return false
}

type Graph struct {
	forwardEdges  map[int64]map[int64]struct{}
	backwardEdges map[int64]map[int64]struct{}
	persons       map[int64]struct{}
}

func loadData(dataFile string) [][]string {
	cnt, err := os.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(cnt)))
	r.Comma = '|'
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	records = records[1:]
	return records

}

func loadGraph(dataFiles ...string) *Graph {
	recs := [][]string{}
	for _, f := range dataFiles {
		recs = append(recs, loadData(f)...)
	}
	fmt.Printf("total: %d\n", len(recs))
	fmt.Printf("first: %s\n", recs[0][0])
	forwardEdges := map[int64]map[int64]struct{}{}
	backwardEdges := map[int64]map[int64]struct{}{}
	persons := map[int64]struct{}{}
	for _, known := range recs {
		from, err := strconv.ParseInt(known[3], 10, 64)
		if err != nil {
			panic(err)
		}
		to, err := strconv.ParseInt(known[4], 10, 64)
		if err != nil {
			panic(err)
		}
		persons[from] = struct{}{}
		persons[to] = struct{}{}
		if forwardEdges[from] == nil {
			forwardEdges[from] = make(map[int64]struct{})
		}
		forwardEdges[from][to] = struct{}{}
		if backwardEdges[to] == nil {
			backwardEdges[to] = make(map[int64]struct{})
		}
		backwardEdges[to][from] = struct{}{}
	}

	g := &Graph{
		backwardEdges: backwardEdges,
		forwardEdges:  forwardEdges,
		persons:       persons,
	}
	return g
}

func NewGraph(dataFiles ...string) *Graph {
	return loadGraph(dataFiles...)
}
