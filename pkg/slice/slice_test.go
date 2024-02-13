package slice

import (
	"fmt"
	"testing"
)

func TestSlice_Append(t *testing.T) {
	lenMsgPattern := "slice length must be equals %d, actual %d"
	capMsgPattern := "slice capacity must be equals %d, actual %d"

	sl := New[string](0, 0)

	if sl.Len() != 0 {
		t.Fatalf(lenMsgPattern, 0, sl.Len())
	}
	if sl.Cap() != 0 {
		t.Fatalf(capMsgPattern, 0, sl.Cap())
	}

	sl.Append("1")

	if sl.Len() != 1 {
		t.Fatalf(lenMsgPattern, 1, sl.Len())
	}
	if sl.Cap() != 1 {
		t.Fatalf(capMsgPattern, 1, sl.Cap())
	}

	sl.Append("2")

	if sl.Len() != 2 {
		t.Fatalf(lenMsgPattern, 2, sl.Len())
	}
	if sl.Cap() != 2 {
		t.Fatalf(capMsgPattern, 2, sl.Cap())
	}

	sl.Append("3")

	if sl.Len() != 3 {
		t.Fatalf(lenMsgPattern, 3, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf(capMsgPattern, 4, sl.Cap())
	}

	sl.Append("4")

	if sl.Len() != 4 {
		t.Fatalf(lenMsgPattern, 4, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf(capMsgPattern, 4, sl.Cap())
	}

	sl.Append("5")

	if sl.Len() != 5 {
		t.Fatalf(lenMsgPattern, 5, sl.Len())
	}
	if sl.Cap() != 8 {
		t.Fatalf(capMsgPattern, 8, sl.Cap())
	}

	for i := 0; i < 5; i++ {
		if (*sl)[i] != fmt.Sprintf("%d", i+1) {
			t.Fatalf(
				"item of slice with index %d  (value %v) is not mathched expected value %d",
				i, (*sl)[i], i+1,
			)
		}
	}
}

func TestSlice_Prepend(t *testing.T) {
	lenMsgPattern := "slice length must be equals %d, actual %d"
	capMsgPattern := "slice capacity must be equals %d, actual %d"

	sl := New[string](0, 0)

	if sl.Len() != 0 {
		t.Fatalf(lenMsgPattern, 0, sl.Len())
	}
	if sl.Cap() != 0 {
		t.Fatalf(capMsgPattern, 0, sl.Cap())
	}

	sl.Prepend("1")

	if sl.Len() != 1 {
		t.Fatalf(lenMsgPattern, 1, sl.Len())
	}
	if sl.Cap() != 1 {
		t.Fatalf(capMsgPattern, 1, sl.Cap())
	}

	sl.Prepend("2")

	if sl.Len() != 2 {
		t.Fatalf(lenMsgPattern, 2, sl.Len())
	}
	if sl.Cap() != 2 {
		t.Fatalf(capMsgPattern, 2, sl.Cap())
	}

	sl.Prepend("3")

	if sl.Len() != 3 {
		t.Fatalf(lenMsgPattern, 3, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf(capMsgPattern, 4, sl.Cap())
	}

	sl.Prepend("4")

	if sl.Len() != 4 {
		t.Fatalf(lenMsgPattern, 4, sl.Len())
	}
	if sl.Cap() != 4 {
		t.Fatalf(capMsgPattern, 4, sl.Cap())
	}

	sl.Prepend("5")

	if sl.Len() != 5 {
		t.Fatalf(lenMsgPattern, 5, sl.Len())
	}
	if sl.Cap() != 8 {
		t.Fatalf(capMsgPattern, 8, sl.Cap())
	}

	for i := 5; i > 0; i-- {
		if (*sl)[5-i] != fmt.Sprintf("%d", i) {
			t.Fatalf(
				"item of slice with index %d (value %v) is not mathched expected value %d",
				5-i, (*sl)[5-i], i,
			)
		}
	}
}
