package finder

import (
	"fmt"
	"math/rand"
	"time"
)

func printPath(e *Edge) {
	path := []*Edge{}
	for e != nil {
		path = append(path, e)
		e = e.Parent
	}
	sz := len(path)
	fmt.Printf("edge: ")
	for i := sz - 1; i >= 0; i-- {
		if path[i].From == -1 {
			continue
		}
		fmt.Printf("%d, ", path[i].From)
	}
	fmt.Printf("%d\n", path[0].To)
}

func (g *Graph) FindAnyShortestPath(from, to int64) int {
	if from == to {
		return 0
	}
	open := [][]Edge{{Edge{From: -1, To: from}}, {Edge{From: -1, To: to}}}
	steps := 0
	dupNode := map[int64]struct{}{}
	for len(open[0]) != 0 && len(open[1]) != 0 {
		steps += 1
		// fmt.Printf("steps: %d, open[0]: %d, open[1]: %d\n", steps, len(open[0]), len(open[1]))
		which := 0
		if len(open[0]) > len(open[1]) {
			which = 1
		}
		// expand open[which]
		newOpen := []Edge{}
		dup := map[int64]*Edge{}
		for i := range open[which] {
			e := &open[which][i]
			if fs, ok := g.forwardEdges[e.To]; ok {
				for f := range fs {
				    if _, ok := dupNode[f]; ok {
						continue
					}
					// dupNode[f] = struct{}{}
					if _, ok := dup[f]; ok {
						continue
					}
					ne := Edge{From: e.To, To: f, Parent: e}
					if e.IsDuplicate(&ne) {
						continue
					}
					dup[f] = &ne
					newOpen = append(newOpen, ne)
				}
			}
			if fs, ok := g.backwardEdges[e.To]; ok {
				for f := range fs {
				    if _, ok := dupNode[f]; ok {
						continue
					}
					// dupNode[f] = struct{}{}
					if _, ok := dup[f]; ok {
						continue
					}
					ne := Edge{From: e.To, To: f, Parent: e, Reverse: true}
					if e.IsDuplicate(&ne) {
						continue
					}
					dup[f] = &ne
					newOpen = append(newOpen, ne)
				}
			}
		}
		for _, e := range open[1-which] {
			if _, ok := dup[e.To]; ok {
				// printPath(&e)
				// printPath(dup[e.To])
				return steps
			}
		}
		open[which] = newOpen
	}
	return -1
}

func (g *Graph) FindAnyShortestPathBFS(from, to int64) int {
	steps := 0
	if from == to {
		return 0
	}
	open := []Edge{{From: -1, To: from}}
	dupNode := map[int64]struct{}{}
	for len(open) != 0 {
		newOpen := []Edge{}
		dup := map[int64]struct{}{}
		steps += 1
		// fmt.Printf("open size: %d\n", len(open))
		for i := range open {
			e := &open[i]
			if fs, ok := g.forwardEdges[e.To]; ok {
				for f := range fs {
				    if _, ok := dupNode[f]; ok {
						continue
					}
					dupNode[f] = struct{}{}

					ne := Edge{From: e.To, To: f, Parent: e}
					if e.IsDuplicate(&ne) {
						continue
					}
					if f == to {
						// printPath(&ne)
						return steps
					}
					if _, ok := dup[f]; ok {
						continue
					}
					dup[f] = struct{}{}
					newOpen = append(newOpen, ne)
				}
			}
			if fs, ok := g.backwardEdges[e.To]; ok {
				for f := range fs {
				    if _, ok := dupNode[f]; ok {
						continue
					}
					dupNode[f] = struct{}{}

					ne := Edge{From: e.To, To: f, Parent: e, Reverse: true}
					if e.IsDuplicate(&ne) {
						continue
					}
					if f == to {
						// printPath(&ne)
						return steps
					}
					if _, ok := dup[f]; ok {
						continue
					}
					dup[f] = struct{}{}
					newOpen = append(newOpen, ne)
				}
			}
		}
		open = newOpen
	}
	return -1
}

func (g *Graph) RandomAnyShortestPath(loop int) {
	persons := make([]int64, 0, len(g.persons))
	maxSteps := 0
	for p := range g.persons {
		persons = append(persons, p)
	}
	for i := 0; i < loop; i++ {
		fromi := rand.Intn(len(g.persons))
		toi := rand.Intn(len(g.persons))
		from := persons[fromi]
		to := persons[toi]
		// fmt.Printf("search path from %d to %d\n", from, to)
		// steps := g.FindAnyShortestPath(from, to)
		start := time.Now()
		steps := g.FindAnyShortestPathBFS(from, to)
		dur1 := time.Since(start)
		start = time.Now()
		steps2 := g.FindAnyShortestPath(from, to)
		dur2 := time.Since(start)
		if steps != steps2 {
			panic(fmt.Sprintf("path from %d to %d, steps bfs: %d, dur1 bfs: %d, steps2: %d, dur2: %d",
				from, to, steps, dur1.Milliseconds(), steps2, dur2.Milliseconds()))
		}
		if steps > maxSteps {
			maxSteps = steps
		}
		fmt.Printf("path from %d to %d, steps bfs: %d, dur1 bfs: %d, steps2: %d, dur2: %d\n",
			from, to, steps, dur1.Milliseconds(), steps2, dur2.Milliseconds())
		// fmt.Printf("finnal steps: %d, dur: %d\n", steps, dur.Microseconds())
	}
	fmt.Printf("max steps: %d\n", maxSteps)
}
