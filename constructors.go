package gollection

//Return a collection of string elements
func NewStringCollection(elements ...string) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of int64 elements
func NewInt64Collection(elements ...int64) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of int32 elements
func NewInt32Collection(elements ...int32) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of int16 elements
func NewInt16Collection(elements ...int16) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of int8 elements
func NewInt8Collection(elements ...int8) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of int elements
func NewIntCollection(elements ...int) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of float64 elements
func NewFloat64Collection(elements ...float64) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of float32 elements
func NewFloat32Collection(elements ...float32) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of byte elements
func NewByteCollection(elements ...byte) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}

//Return a collection of bool elements
func NewBoolCollection(elements ...bool) Collection {
	col := NewCollection()

	for _, e := range elements {
		col.Add(NewElement(e))
	}

	return col
}
