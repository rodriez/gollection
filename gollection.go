package gollection

import (
	"sort"
)

type gollection []Element

func NewCollection(elements ...interface{}) Collection {
	col := gollection{}

	for _, e := range elements {
		col = append(col, NewElement(e))
	}

	return &col
}

//Add an element to the collection
func (col *gollection) Add(element Element) Collection {
	*col = append(*col, element)

	return col
}

//Return true if the collection is empty
func (col *gollection) Empty() bool {
	return col.Length() == 0
}

//Return a new collection with filtered elements
func (col *gollection) Filter(criteria func(e Element) bool) Collection {
	filtered := gollection{}

	for _, e := range *col {
		if criteria(e) {
			filtered = append(filtered, e)
		}
	}

	return &filtered
}

//Search an element in the collection
func (col *gollection) Find(criteria func(Element) bool, offset int) Element {
	for i := offset; i < col.Length(); i++ {
		if criteria(col.Get(i)) {
			return col.Get(i)
		}
	}

	return Element{}
}

//Execute a funcion to each element of the collection
func (col *gollection) ForEach(fn func(int, Element)) {
	for i, e := range *col {
		fn(i, e)
	}
}

//Return the element corresponding to the index value
func (col *gollection) Get(i int) Element {
	return (*col)[i]
}

//Return the amount of elements of the collection
func (col *gollection) Length() int {
	return len(*col)
}

//Return a new collection with all the elements converted
func (col *gollection) Map(converter func(e Element) Element) Collection {
	mapped := gollection{}

	for _, e := range *col {
		mapped = append(mapped, converter(e))
	}

	return &mapped
}

//Reduce the collection to one value
func (col *gollection) Reduce(reductor func(acc Element, e Element) Element, startingValue Element) Element {
	for _, e := range *col {
		startingValue = reductor(startingValue, e)
	}

	return startingValue
}

//Search an element in the collection in the opposite direction to Find
func (col *gollection) ReverseFind(criteria func(Element) bool, offset int) Element {
	for i := (col.Length() - 1) - offset; i >= 0; i-- {
		if criteria(col.Get(i)) {
			return col.Get(i)
		}
	}

	return Element{}
}

//Return a new collection with all the elements ordered
func (col *gollection) Sort(criteria func(e1 Element, e2 Element) bool) Collection {
	sort.Slice(*col, func(i, j int) bool {
		return criteria(col.Get(i), col.Get(j))
	})

	return col
}
