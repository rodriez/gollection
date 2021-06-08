package gollection

const SEARCH_TO_LEFT = -1
const SEARCH_END = 0
const SEARCH_TO_RIGHT = 1

type Collection interface {
	//Add an element to the collection
	Add(e Element) Collection

	//Return true if the collection is empty
	Empty() bool

	//Returns a new collection with filtered elements
	//according to the criteria
	Filter(criteria func(e *Element) bool) Collection

	//Finds and returns the first element in the collection that satisfies the criteria
	//If no element satisfie it, then it returns nil
	Find(criteria func(e *Element) bool, offset int) *Element

	//Execute the passed funcion to each element of the collection
	ForEach(exec func(idx int, e *Element))

	//Returns the element corresponding to the index value
	Get(idx int) *Element

	//Return the amount of elements of the collection
	Length() int

	//Applies the converter to each element in the collection
	//and return the new collection
	Map(converter func(e *Element) Element) Collection

	//Returns the next element corresponding to the index value
	//if the index not exist return nil
	Next(idx int) *Element

	//Returns the previous element corresponding to the index value
	//if the index not exist return nil
	Prev(idx int) *Element

	//Reduce the collection to a value, applying reducer to each element
	Reduce(reducer func(acc *Element, e *Element) Element, startingValue Element) *Element

	//Finds and returns the first element in the collection according to the criteria
	//starting from the end of the collection
	//If no element satisfie it, then it returns nil
	ReverseFind(criteria func(e *Element) bool, offset int) *Element

	//Returns a new collection with all the elements ordered
	//according to the criteria
	Sort(criteria func(e1 *Element, e2 *Element) bool) Collection

	//Returns a new collection with the elements between from and to
	SubCollection(from, to int) Collection

	//Given a ordered collection BinaryFind search and returns the first element that satisfies the criteria
	//If the collection if empty or any element in the collection satisfies the criteria returns nil
	//also if the returned sense doesn't match with any supported direction.\n
	//crtieria must indicates the sense of the search:
	//returns SEARCH_TO_LEFT if you want to continue searching at the bottom of the collection
	//returns SEARCH_TO_RIGHT if you want to keep looking at the top of the collection
	//returns SEARCH_END when you find what you are looking for
	BinaryFind(criteria func(*Element) int) *Element
}
