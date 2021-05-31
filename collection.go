package gollection

type Collection interface {
	//Add an element to the collection
	Add(e Element) Collection

	//Return true if the collection is empty
	Empty() bool

	//Return a new collection with filtered elements
	Filter(criteria func(e *Element) bool) Collection

	//Search an element in the collection
	Find(criteria func(e *Element) bool, offset int) *Element

	//Execute a funcion to each element of the collection
	ForEach(exec func(idx int, e *Element))

	//Return the element corresponding to the index value
	Get(idx int) *Element

	//Return the amount of elements of the collection
	Length() int

	//Return a new collection with all the elements converted
	Map(converter func(e *Element) Element) Collection

	//Return the next element corresponding to the index value
	//if the index not exist return nil
	Next(idx int) *Element

	//Return the previous element corresponding to the index value
	//if the index not exist return nil
	Prev(idx int) *Element

	//Reduce the collection to one value
	Reduce(reductor func(acc *Element, e *Element) Element, startingValue Element) *Element

	//Search an element in the collection in the opposite direction to Find
	ReverseFind(criteria func(e *Element) bool, offset int) *Element

	//Return a new collection with all the elements ordered
	Sort(criteria func(e1 *Element, e2 *Element) bool) Collection
}
