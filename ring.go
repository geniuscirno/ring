package ring

type Ring[T any] struct {
	elems  []T
	head   int
	tail   int
	length int
	cap    int
}

func New[T any](cap int) *Ring[T] {
	return &Ring[T]{
		elems:  make([]T, cap, cap),
		head:   0,
		tail:   0,
		length: 0,
		cap:    cap,
	}
}

func (r *Ring[T]) Push(elem T) {
	r.tail =
}

func (r *Ring[T]) Pop() (T, bool) {
	return zeroOf[T](), false
}

func (r *Ring[T]) Len() int {
	return 0
}

func zeroOf[T any]() T {
	var zero T
	return zero
}
