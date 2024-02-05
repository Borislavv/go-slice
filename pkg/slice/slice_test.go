package slice

import "testing"

func TestSlice_Append(t *testing.T) {
	sl := New[string](0, 0)

	if sl.Len() != 0 {
		t.Fatalf("slice length must be equals %d, actual %d", 0, sl.Len())
	}
	if sl.Cap() != 0 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 0, sl.Cap())
	}

	sl.Append("1")

	if sl.Len() != 1 {
		t.Fatalf("slice length must be equals %d, actual %d", 1, sl.Len())
	}
	if sl.Cap() != 1 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 1, sl.Cap())
	}

	sl.Append("2")

	if sl.Len() != 2 {
		t.Fatalf("slice length must be equals %d, actual %d", 2, sl.Len())
	}
	if sl.Cap() != 2 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 2, sl.Cap())
	}

	sl.Append("3")

	if sl.Len() != 3 {
		t.Fatalf("slice length must be equals %d, actual %d", 3, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 4, sl.Cap())
	}

	sl.Append("4")

	if sl.Len() != 4 {
		t.Fatalf("slice length must be equals %d, actual %d", 4, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 4, sl.Cap())
	}

	sl.Append("5")

	if sl.Len() != 5 {
		t.Fatalf("slice length must be equals %d, actual %d", 5, sl.Len())
	}
	if sl.Cap() != 8 {
		t.Fatalf("slice capacity must be equals %d, actual %d", 8, sl.Cap())
	}
}
