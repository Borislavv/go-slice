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
	itLen := len(items)
	slLen := s.Len()
	slCap := s.Cap()
	nwLen := slLen + itLen

	if slCap >= nwLen {
		*s = append(*s, items...)
		copy((*s)[itLen:], (*s)[:slLen])
		copy(*s, items)
	} else {
		nwCap := slCap
		for {
			nwCap = nwCap * 2
			if nwCap == 0 {
				nwCap = 1
			}
			if nwCap >= nwLen {
				nwSl := make([]T, itLen, nwCap)
				copy(nwSl, items)
				*s = append(nwSl, *s...)
				return
			}
		}
	}
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

func (s *Slice[T]) Iter() chan T {
	ch := make(chan T)
	defer close(ch)
	go func() {
		for _, v := range *s {
			ch <- v
		}
	}()
	return ch
}

func (s *Slice[T]) Map(callable func(k int, v T)) {
	for k, v := range *s {
		callable(k, v)
	}
}

func (s *Slice[T]) Len() int {
	return len(*s)
}

func (s *Slice[T]) Cap() int {
	return cap(*s)
}
