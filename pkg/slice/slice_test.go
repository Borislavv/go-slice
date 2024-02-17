package slice

import (
	"fmt"
	"testing"
)

func TestSlice_New(t *testing.T) {
	s := []string{"first", "second", "third", "forth", "fifth"}

	sl := New[string](0, 0, s)

	if sl.Len() != 5 {
		t.Fatalf("expected length '5' but actual '%d'", sl.Len())
	} else if sl.Cap() != 5 {
		t.Fatalf("expected capacity '8' but actual '%d'", sl.Cap())
	}

	for k, v := range s {
		if (*sl)[k] != v {
			t.Fatalf("value is not matched by index '%d', expected '%v', actual '%v'", k, v, (*sl)[k])
		}
	}

	sl = New[string](10, 20, s)

	if sl.Len() != 15 {
		t.Fatalf("expected length '5' but actual '%d'", sl.Len())
	} else if sl.Cap() != 20 {
		t.Fatalf("expected capacity '8' but actual '%d'", sl.Cap())
	}

	for k, v := range s {
		if (*sl)[k+10] != v {
			t.Fatalf("value is not matched by index '%d', expected '%v', actual '%v'", k, v, (*sl)[k+10])
		}
	}
}

func TestSlice_Append(t *testing.T) {
	lenMsgPattern := "slice length must be equals '%d', actual '%d'"
	capMsgPattern := "slice capacity must be equals '%d', actual '%d'"

	sl := New[string](0, 0, nil)

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
				"item of slice with index '%d' (value '%v') is not mathched expected value '%d'",
				i, (*sl)[i], i+1,
			)
		}
	}
}

func TestSlice_Prepend(t *testing.T) {
	lenMsgPattern := "slice length must be equals '%d', actual '%d'"
	capMsgPattern := "slice capacity must be equals '%d', actual '%d'"

	sl := New[string](0, 0, nil)

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
				"item of slice with index '%d' (value '%v') is not mathched expected value '%d'",
				5-i, (*sl)[5-i], i,
			)
		}
	}
}

func TestSlice_Pop(t *testing.T) {
	sl := New[string](0, 0, []string{"first", "second", "third", "forth", "fifth"})

	fifth, ok := sl.Pop()
	if !ok {
		t.Fatal("failed to pop element of slice")
	}
	if fifth != "fifth" {
		t.Fatalf("poped element mismatched, expected 'fifth', actual '%v'", fifth)
	}
	if sl.Len() != 4 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 4)
	}

	forth, ok := sl.Pop()
	if !ok {
		t.Fatal("failed to pop element of slice")
	}
	if forth != "forth" {
		t.Fatalf("poped element mismatched, expected 'forth', actual '%v'", forth)
	}
	if sl.Len() != 3 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 3)
	}
}

func TestSlice_Shift(t *testing.T) {
	sl := New[string](0, 0, []string{"first", "second", "third", "forth", "fifth"})

	first, ok := sl.Shift()
	if !ok {
		t.Fatal("failed to shift element of slice")
	}
	if first != "first" {
		t.Fatalf("shifted element mismatched, expected 'first', actual '%v'", first)
	}
	if sl.Len() != 4 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 4)
	}

	second, ok := sl.Shift()
	if !ok {
		t.Fatal("failed to pop element of slice")
	}
	if second != "second" {
		t.Fatalf("shifted element mismatched, expected 'second', actual '%v'", second)
	}
	if sl.Len() != 3 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 3)
	}
}

func TestSlice_Sub(t *testing.T) {
	sl := New[string](0, 0, []string{"first", "second", "third", "forth", "fifth"})

	second, ok := sl.Sub(1)
	if !ok {
		t.Fatal("failed to sub element of slice")
	}
	if second != "second" {
		t.Fatalf("cutted element mismatched, expected 'second', actual '%v'", second)
	}
	if sl.Len() != 4 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 4)
	}

	forth, ok := sl.Sub(2)
	if !ok {
		t.Fatal("failed to sub element of slice")
	}
	if forth != "forth" {
		t.Fatalf("subbed element mismatched, expected 'forth', actual '%v'", forth)
	}
	if sl.Len() != 3 {
		t.Fatalf("length of the target slice was not decreased, actual '%d', expected '%d'", sl.Len(), 3)
	}
}

func TestSlice_Get(t *testing.T) {
	s := []string{"first", "second", "third", "forth", "fifth"}

	sl := New[string](0, 0, s)

	for k, v := range s {
		val, ok := sl.Get(k)
		if !ok {
			t.Fatalf("failed to fetch element by index '%d' from slice", k)
		}
		if val != v {
			t.Fatalf("mismatch fetched value, expected '%v', actual '%v'", v, val)
		}
	}

	v, ok := sl.Get(10)
	if ok {
		t.Fatalf("found unexists element for index '10' with value '%v'", v)
	}
}

func TestSlice_Has(t *testing.T) {
	s := []string{"first", "second", "third", "forth", "fifth"}

	sl := New[string](0, 0, s)

	for k, _ := range s {
		ok := sl.Has(k)
		if !ok {
			t.Fatalf("has: failed to fetch element by index '%d' from slice", k)
		}
	}

	ok := sl.Has(10)
	if ok {
		t.Fatalf("has: found unexists element for index '10'")
	}
}
