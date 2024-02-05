package slice

type Interface[T any] interface {
	New(len int, cap int) *Slice[T]
	Append(items ...T)
	Prepend(items ...T)
	Pop() (item T, ok bool)
	Shift() (item T, ok bool)
	Sub(idx int) (item T, ok bool)
	Get(idx int) (item T, ok bool)
	Has(idx int) (ok bool)
	Remove(idx int) (ok bool)
	Len() int
	Cap() int
}
