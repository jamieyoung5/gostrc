package dlx_test

import (
	"github.com/jamieyoung5/gostrc/dlx"
	"reflect"
	"sort"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	cols := []string{"A", "B", "C"}
	m := dlx.NewMatrix(cols)

	if m.Root == nil {
		t.Fatal("Matrix root should not be nil")
	}

	count := 0
	curr := m.Root.Right
	for curr != m.Root {
		count++
		if curr.Column == nil {
			t.Errorf("Header node at index %d has nil Column", count)
		}
		curr = curr.Right
	}

	if count != 3 {
		t.Errorf("Expected 3 columns, got %d", count)
	}
}

func TestAppendRow(t *testing.T) {
	cols := []string{"A", "B", "C"}
	m := dlx.NewMatrix(cols)

	m.AppendRow([]string{"A", "C"})

	headerA := m.Root.Right
	if headerA.Column.ID != "A" {
		t.Fatalf("Expected first column A, got %s", headerA.Column.ID)
	}
	if headerA.Column.Size != 1 {
		t.Errorf("Column A size expected 1, got %d", headerA.Column.Size)
	}

	headerB := headerA.Right
	if headerB.Column.ID != "B" {
		t.Fatalf("Expected second column B, got %s", headerB.Column.ID)
	}
	if headerB.Column.Size != 0 {
		t.Errorf("Column B size expected 0, got %d", headerB.Column.Size)
	}

	headerC := headerB.Right
	if headerC.Column.ID != "C" {
		t.Fatalf("Expected third column C, got %s", headerC.Column.ID)
	}
	if headerC.Column.Size != 1 {
		t.Errorf("Column C size expected 1, got %d", headerC.Column.Size)
	}
}

func TestSolveExactCover(t *testing.T) {

	cols := []string{"1", "2", "3", "4"}
	m := dlx.NewMatrix(cols)

	m.AppendRow([]string{"1", "2"})
	m.AppendRow([]string{"3", "4"})
	m.AppendRow([]string{"1", "3"})
	m.AppendRow([]string{"2", "4"})

	var solutions [][]*dlx.Node
	m.Search(nil, &solutions, 10)

	if len(solutions) != 2 {
		t.Fatalf("Expected 2 solutions, got %d", len(solutions))
	}

	var stringSolutions [][]string

	for _, sol := range solutions {
		var currentSolRows []string
		for _, node := range sol {
			rowCols := getRowColumnIDs(node)
			sig := ""
			for i, colID := range rowCols {
				if i > 0 {
					sig += "-"
				}
				sig += colID
			}
			currentSolRows = append(currentSolRows, sig)
		}
		sort.Strings(currentSolRows)
		stringSolutions = append(stringSolutions, currentSolRows)
	}

	expectedSol1 := []string{"1-2", "3-4"}
	expectedSol2 := []string{"1-3", "2-4"}

	found1 := false
	found2 := false

	for _, sol := range stringSolutions {
		if reflect.DeepEqual(sol, expectedSol1) {
			found1 = true
		}
		if reflect.DeepEqual(sol, expectedSol2) {
			found2 = true
		}
	}

	if !found1 {
		t.Error("Did not find solution {1,2}, {3,4}")
	}
	if !found2 {
		t.Error("Did not find solution {1,3}, {2,4}")
	}
}

func TestNoSolution(t *testing.T) {
	cols := []string{"1", "2", "3"}
	m := dlx.NewMatrix(cols)
	m.AppendRow([]string{"1", "2"})
	m.AppendRow([]string{"2", "3"})

	var solutions [][]*dlx.Node
	m.Search(nil, &solutions, 10)

	if len(solutions) != 0 {
		t.Errorf("Expected 0 solutions, got %d", len(solutions))
	}
}

func TestKnuthSimpleExample(t *testing.T) {
	cols := []string{"A", "B", "C"}
	m := dlx.NewMatrix(cols)
	m.AppendRow([]string{"A"})
	m.AppendRow([]string{"B"})
	m.AppendRow([]string{"C"})
	m.AppendRow([]string{"A", "B"})

	var solutions [][]*dlx.Node
	m.Search(nil, &solutions, 10)

	if len(solutions) != 2 {
		t.Fatalf("Expected 1 solution, got %d", len(solutions))
	}

	sol := solutions[0]
	if len(sol) != 3 {
		t.Errorf("Expected solution to have 3 rows, got %d", len(sol))
	}
}

func getRowColumnIDs(n *dlx.Node) []string {
	var ids []string
	ids = append(ids, n.Column.ID)

	for curr := n.Right; curr != n; curr = curr.Right {
		ids = append(ids, curr.Column.ID)
	}

	sort.Strings(ids)
	return ids
}
