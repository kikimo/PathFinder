package finder

import (
	"testing"
)

func TestFinderBFS(t *testing.T) {
	f1 := "/home/alex/src/ldbc_snb_datagen_spark/sf01/graphs/csv/raw/singular-projected-fk/dynamic/Person_knows_Person/part-00000-ec1ad1b9-1460-43a3-b25a-2916a98f8ec9-c000.csv"
	g := NewGraph(f1)
	var from int64 = 24189255812361
	var to int64 = 37383395345729

	steps := g.FindAnyShortestPathBFS(from, to)
	// edge: 24189255812361, 15393162789420, 24189255812167, 37383395345729
	t.Logf("find from %d to %d steps: %d\n", from, to, steps)
}

func TestFinderTwoWayBFS(t *testing.T) {
	f1 := "/home/alex/src/ldbc_snb_datagen_spark/sf01/graphs/csv/raw/singular-projected-fk/dynamic/Person_knows_Person/part-00000-ec1ad1b9-1460-43a3-b25a-2916a98f8ec9-c000.csv"
	g := NewGraph(f1)
	var from int64 = 24189255812361
	var to int64 = 37383395345729

	steps := g.FindAnyShortestPath(from, to)
	// edge: 24189255812361, 15393162789420, 24189255812167, 37383395345729
	t.Logf("find from %d to %d steps: %d\n", from, to, steps)
}
