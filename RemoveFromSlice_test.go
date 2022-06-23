package goTemplateTools

import "testing"

func TestRemoveFromSLice(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromSlice(A, 1)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "A" {
		t.Errorf("Expected %s got %s\n", []string{"A", "C"}, A)
	}
	if A[1] != "C" {
		t.Errorf("Expected %s got %s\n", []string{"A", "C"}, A)
	}

}

func TestRemoveFromSLice2(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromSlice(A, 2)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "A" {
		t.Errorf("Expected %s got %s\n", []string{"A", "B"}, A)
	}
	if A[1] != "B" {
		t.Errorf("Expected %s got %s\n", []string{"A", "B"}, A)
	}

}

func TestRemoveFromSLice3(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromSlice(A, 0)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "C" {
		t.Errorf("Expected %s got %s\n", []string{"C", "B"}, A)
	}
	if A[1] != "B" {
		t.Errorf("Expected %s got %s\n", []string{"C", "B"}, A)
	}

}

func TestRemoveFromSLice4(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A"}
	A = RemoveFromSlice(A, 0)
	if len(A) != 0 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}

}

func TestRemoveFromOrderedSlice(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromOrderedSlice(A, 1)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "A" {
		t.Errorf("Expected %s got %s\n", []string{"A", "C"}, A)
	}
	if A[1] != "C" {
		t.Errorf("Expected %s got %s\n", []string{"A", "C"}, A)
	}

}

func TestRemoveFromOrderedSlice2(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromOrderedSlice(A, 2)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "A" {
		t.Errorf("Expected %s got %s\n", []string{"A", "B"}, A)
	}
	if A[1] != "B" {
		t.Errorf("Expected %s got %s\n", []string{"A", "B"}, A)
	}

}

func TestRemoveFromOrderedSlice3(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A", "B", "C"}
	A = RemoveFromOrderedSlice(A, 0)
	if len(A) != 2 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}
	if A[0] != "B" {
		t.Errorf("Expected %s got %s\n", []string{"B", "C"}, A)
	}
	if A[1] != "C" {
		t.Errorf("Expected %s got %s\n", []string{"B", "C"}, A)
	}

}

func TestRemoveFromOrderedSlice4(t *testing.T) {

	// Basic Slice -- remove from middle
	A := []string{"A"}
	A = RemoveFromOrderedSlice(A, 0)
	if len(A) != 0 {
		t.Errorf("Expected Length of 2 got %d\n", len(A))
	}

}
