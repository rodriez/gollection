package gollection

func NewStringCollection(elements ...string) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewInt64Collection(elements ...int64) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewInt32Collection(elements ...int32) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewInt16Collection(elements ...int16) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewInt8Collection(elements ...int8) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewIntCollection(elements ...int) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewFloat64Collection(elements ...float64) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewFloat32Collection(elements ...float32) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewByteCollection(elements ...byte) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

func NewBoolCollection(elements ...bool) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}
