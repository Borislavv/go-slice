package slice

type Slice[T any] []T

func New[T any](len int, cap int) *Slice[T] {
	var s Slice[T] = make([]T, len, cap)
	return &s
}

func (s *Slice[T]) Append(items ...T) {
	*s = append(*s, items...)
}

func (s *Slice[T]) Prepend(items ...T) {
	*s = append(items, *s...)
}

func (s *Slice[T]) Pop() (item T, ok bool) {
	if len(*s) == 0 {
		return item, false
	}
	item = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item, true
}

func (s *Slice[T]) Shift() (item T, ok bool) {
	if len(*s) == 0 {
		return item, false
	}
	item = (*s)[0]
	*s = (*s)[1:]
	return item, true
}

func (s *Slice[T]) Sub(idx int) (item T, ok bool) {
	if idx < 0 || len(*s) == 0 || len(*s) < idx {
		return item, false
	}
	item = (*s)[idx]
	*s = append((*s)[:idx], (*s)[idx+1:]...)
	return item, ok
}

func (s *Slice[T]) Get(idx int) (item T, ok bool) {
	if len(*s) == 0 || len(*s)-1 < idx {
		return item, false
	}
	return (*s)[idx], true
}

func (s *Slice[T]) Has(idx int) (ok bool) {
	return len(*s)-1 >= idx
}

func (s *Slice[T]) Remove(idx int) (ok bool) {
	if idx < 0 || len(*s) == 0 || len(*s) < idx {
		return true
	}
	*s = append((*s)[:idx], (*s)[idx+1:]...)
	return true
}

func (s *Slice[T]) Len() int {
	return len(*s)
}

func (s *Slice[T]) Cap() int {
	return cap(*s)
}
