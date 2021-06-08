package gollection

type concurrentGollection struct {
	Collection
	Threads int
}

func NewConcurrentCollection(threads int, elements ...interface{}) Collection {
	if threads > len(elements) {
		threads = len(elements)
	}

	col := gollection{}

	for _, e := range elements {
		col = append(col, NewElement(e))
	}

	return &concurrentGollection{
		Collection: &col,
		Threads:    threads,
	}
}

//Add an element to the collection
func (cc *concurrentGollection) Add(e Element) Collection {
	cc.Collection.Add(e)

	return cc
}

//Return a new collection with filtered elements
func (cc *concurrentGollection) Filter(criteria func(e *Element) bool) Collection {
	size := cc.Length()
	iterations, elements := cc.createChannels(size)
	defer close(elements)

	for i := 0; i < cc.Threads; i++ {
		go cc.filterWorker(iterations, elements, criteria)
	}

	cc.setIterations(iterations, size)
	return cc.readElements(elements, size)
}

func (cc *concurrentGollection) filterWorker(iterations <-chan int, output chan<- *Element, criteria func(e *Element) bool) {
	for i := range iterations {
		if e := cc.Get(i); criteria(e) {
			output <- e
		}

		output <- nil
	}
}

//Return a new collection with all the elements converted
func (cc *concurrentGollection) Map(converter func(e *Element) Element) Collection {
	size := cc.Length()
	iterations, elements := cc.createChannels(size)
	defer close(elements)

	for i := 0; i < cc.Threads; i++ {
		go cc.mapWorker(iterations, elements, converter)
	}

	cc.setIterations(iterations, size)
	return cc.readElements(elements, size)
}

func (cc *concurrentGollection) mapWorker(iterations <-chan int, output chan<- *Element, converter func(e *Element) Element) {
	for i := range iterations {
		e := cc.Get(i)
		ne := converter(e)
		output <- &ne
	}
}

func (cc *concurrentGollection) createChannels(size int) (chan int, chan *Element) {
	iterations := make(chan int, size)
	elements := make(chan *Element, size)

	return iterations, elements
}

func (cc *concurrentGollection) setIterations(iterations chan<- int, count int) {
	for i := 0; i < count; i++ {
		iterations <- i
	}
	close(iterations)
}

func (cc *concurrentGollection) readElements(elements <-chan *Element, count int) Collection {
	col := NewConcurrentCollection(cc.Threads)
	for i := 0; i < count; i++ {
		if e := <-elements; e != nil {
			col.Add(*e)
		}
	}

	return col
}

//Returns a new collection with all the elements ordered
//according to the criteria
func (cc *concurrentGollection) Sort(criteria func(e1 *Element, e2 *Element) bool) Collection {
	sorted := cc.Collection.Sort(criteria)

	return &concurrentGollection{
		Collection: sorted,
		Threads:    cc.Threads,
	}
}

//Returns a new collection with the elements between from and to
func (cc *concurrentGollection) SubCollection(from, to int) Collection {
	sub := cc.Collection.SubCollection(from, to)

	return &concurrentGollection{
		Collection: sub,
		Threads:    cc.Threads,
	}
}
