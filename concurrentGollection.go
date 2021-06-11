package gollection

import (
	"errors"

	"github.com/rodriez/goncu"
)

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
	resp, _ := goncu.NewPool(cc.Threads).
		DO(func(n int) (interface{}, error) {
			if e := cc.Get(n); criteria(e) {
				return e, nil
			}

			return nil, errors.New("filtered")
		}).
		Run(cc.Length())

	return NewConcurrentCollection(cc.Threads, resp.Hits...)
}

//Return a new collection with all the elements converted
func (cc *concurrentGollection) Map(converter func(e *Element) Element) Collection {
	resp, _ := goncu.NewPool(cc.Threads).
		DO(func(n int) (interface{}, error) {
			return converter(cc.Get(n)), nil
		}).
		Run(cc.Length())

	return NewConcurrentCollection(cc.Threads, resp.Hits...)
}

//Return a new collection with all the elements converted
func (cc *concurrentGollection) ForEach(exec func(idx int, e *Element)) {
	goncu.NewPool(cc.Threads).
		DO(func(n int) (interface{}, error) {
			exec(n, cc.Get(n))
			return nil, nil
		}).
		Run(cc.Length())
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
