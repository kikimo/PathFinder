package main

import "github.com/kikimo/PathFinder/finder"

func main() {
	f1 := "/home/alex/src/ldbc_snb_datagen_spark/sf10/graphs/csv/raw/singular-projected-fk/dynamic/Person_knows_Person/part-00000-c74e39ba-8a8f-4874-a271-dd35c59458cc-c000.csv"
	g := finder.NewGraph(f1)
	g.RandomAnyShortestPath(1000)
}
