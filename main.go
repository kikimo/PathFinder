package main

import (
	"fmt"
	"time"

	"github.com/kikimo/PathFinder/finder"
)

func main() {
	// f1 := "/home/alex/src/ldbc_snb_datagen_spark/sf10/graphs/csv/raw/singular-projected-fk/dynamic/Person_knows_Person/part-00000-c74e39ba-8a8f-4874-a271-dd35c59458cc-c000.csv"
	f1 := "/Users/wenlinwu/tmp/part-00000-c74e39ba-8a8f-4874-a271-dd35c59458cc-c000.csv"
	g := finder.NewGraph(f1)
	// g.RandomAnyShortestPath(1000)
	start := time.Now()
	g.FindAnyShortestPathBFS(28587302366215, 19791209354607)
	dur := time.Since(start)
	fmt.Printf("dur: %d\n", dur.Milliseconds())
}
