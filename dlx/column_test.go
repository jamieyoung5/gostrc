package dlx_test

import (
	"github.com/jamieyoung5/gostrc/dlx"
	"testing"
)

func TestColumn_Cover_Simple(t *testing.T) {
	cols := []string{"A", "B"}
	m := dlx.NewMatrix(cols)
	m.AppendRow([]string{"A", "B"})

	colA := m.Root.Right.Column
	colB := colA.Right.Column

	if colB.Size != 1 {
		t.Fatalf("Setup failed: Expected Column B size 1, got %d", colB.Size)
	}

	colA.Cover()

	if m.Root.Right != colB.Node {
		t.Errorf("Header link broken: Root.Right should point to B, points to %v", m.Root.Right.Column.ID)
	}
	if colB.Left != m.Root {
		t.Errorf("Header link broken: B.Left should point to Root")
	}

	if colB.Down != colB.Node {
		t.Error("Vertical link broken: Column B should be empty (Down -> Header)")
	}
	if colB.Up != colB.Node {
		t.Error("Vertical link broken: Column B should be empty (Up -> Header)")
	}

	if colB.Size != 0 {
		t.Errorf("Size update failed: Expected B size 0, got %d", colB.Size)
	}
}

func TestColumn_Uncover_Simple(t *testing.T) {
	cols := []string{"A", "B"}
	m := dlx.NewMatrix(cols)
	m.AppendRow([]string{"A", "B"})

	colA := m.Root.Right.Column
	colB := colA.Right.Column

	colA.Cover()

	colA.Uncover()

	if m.Root.Right != colA.Node {
		t.Error("Header restoration failed: Root.Right should be A")
	}
	if colA.Right != colB.Node {
		t.Error("Header restoration failed: A.Right should be B")
	}
	if colB.Left != colA.Node {
		t.Error("Header restoration failed: B.Left should be A")
	}

	if colB.Down == colB.Node {
		t.Error("Vertical restoration failed: Column B is empty")
	}
	rowNodeB := colB.Down
	if rowNodeB.Up != colB.Node {
		t.Error("Vertical restoration failed: Node Up pointer incorrect")
	}
	if rowNodeB.Down != colB.Node {
		t.Error("Vertical restoration failed: Node Down pointer incorrect")
	}

	if colB.Size != 1 {
		t.Errorf("Size restoration failed: Expected B size 1, got %d", colB.Size)
	}
}

func TestColumn_Cover_ComplexInteractions(t *testing.T) {
	cols := []string{"A", "B", "C"}
	m := dlx.NewMatrix(cols)
	m.AppendRow([]string{"A", "B"})
	m.AppendRow([]string{"A", "C"})
	m.AppendRow([]string{"B", "C"})

	colA := m.Root.Right.Column
	colB := colA.Right.Column
	colC := colB.Right.Column

	if colB.Size != 2 || colC.Size != 2 {
		t.Fatal("Setup failed: Columns B and C should have size 2")
	}

	colA.Cover()

	if colB.Size != 1 {
		t.Errorf("Complex Cover: Expected B size 1, got %d", colB.Size)
	}

	remainingNodeB := colB.Down
	if remainingNodeB == colB.Node {
		t.Fatal("Complex Cover: B became empty unexpected")
	}
	if remainingNodeB.Right.Column != colC {
		t.Error("Complex Cover: Remaining node in B is not from Row 3 (should link to C)")
	}

	if colC.Size != 1 {
		t.Errorf("Complex Cover: Expected C size 1, got %d", colC.Size)
	}
	remainingNodeC := colC.Down
	if remainingNodeC.Left.Column != colB {
		t.Error("Complex Cover: Remaining node in C is not from Row 3 (should link to B)")
	}

	colA.Uncover()

	if colB.Size != 2 {
		t.Errorf("Complex Uncover: B size mismatch, expected 2, got %d", colB.Size)
	}
	if colC.Size != 2 {
		t.Errorf("Complex Uncover: C size mismatch, expected 2, got %d", colC.Size)
	}
}
