package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
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
	forwardEdges  map[int64]int64
	backwardEdges map[int64]int64
	persons       map[int64]struct{}
}

func (g *Graph) FindAllShortestPath(from, to int64) {
	// TODO
}

func (g *Graph) FindAnyShortestPath(from, to int64) int {
	if from == to {
		return 0
	}
	open := [][]Edge{{Edge{From: -1, To: from}}, {Edge{From: -1, To: to}}}
	steps := 0
	for len(open[0]) != 0 && len(open[1]) != 0 {
		steps += 1
		fmt.Printf("steps: %d, open[0]: %d, open[1]: %d\n", steps, len(open[0]), len(open[1]))
		which := 0
		if len(open[0]) > len(open[1]) {
			which = 1
		}
		// expand open[which]
		newOpen := []Edge{}
		dup := map[int64]struct{}{}
		for i, _ := range open[which] {
			e := &open[which][i]
			if f, ok := g.forwardEdges[e.To]; ok {
				if _, ok := dup[f]; ok {
					continue
				}
				ne := Edge{From: e.To, To: f, Parent: e}
				if e.IsDuplicate(&ne) {
					continue
				}
				dup[f] = struct{}{}
				newOpen = append(newOpen, ne)
			}
			if f, ok := g.backwardEdges[e.To]; ok {
				if _, ok := dup[f]; ok {
					continue
				}
				ne := Edge{From: e.To, To: f, Parent: e, Reverse: true}
				if e.IsDuplicate(&ne) {
					continue
				}
				dup[f] = struct{}{}
				newOpen = append(newOpen, ne)
			}
		}
		for _, e := range open[1-which] {
			if _, ok := dup[e.To]; ok {
				return steps
			}
		}
		open[which] = newOpen
	}
	return -1
}

func (g *Graph) randomAnyShortestPath() {
	persons := make([]int64, 0, len(g.persons))
	for p, _ := range g.persons {
		persons = append(persons, p)
	}
	for i := 0; i < 100; i++ {
		fromi := rand.Intn(len(g.persons))
		toi := rand.Intn(len(g.persons))
		from := persons[fromi]
		to := persons[toi]
		fmt.Printf("search path from %d to %d\n", from, to)
		steps := g.FindAnyShortestPath(from, to)
		fmt.Printf("steps: %d\n", steps)
	}
}

func main() {
	knowFile := "/home/alex/src/ldbc_snb_datagen_spark/sf10/graphs/csv/raw/singular-projected-fk/dynamic/Person_knows_Person/part-00000-c74e39ba-8a8f-4874-a271-dd35c59458cc-c000.csv"
	cnt, err := os.ReadFile(knowFile)
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
	fmt.Printf("total: %d\n", len(records))
	fmt.Printf("first: %s\n", records[0][0])
	forwardEdges := map[int64]int64{}
	backwardEdges := map[int64]int64{}
	persons := map[int64]struct{}{}
	for _, known := range records {
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
		forwardEdges[from] = to
		backwardEdges[to] = from
	}

	g := Graph{
		backwardEdges: backwardEdges,
		forwardEdges:  forwardEdges,
		persons:       persons,
	}
	var from int64 = 17592186094270
	var to int64 = 9995
	steps := g.FindAnyShortestPath(from, to)
	fmt.Printf("find from %d to %d steps: %d\n", from, to, steps)
	g.randomAnyShortestPath()
}
