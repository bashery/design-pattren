package strain

type Ints []int

func (n Ints) Keep() Ints {
	return n
}

func (n Ints) Discard() Ints {

	return n
}
