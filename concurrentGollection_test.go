package gollection_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rodriez/gollection"
	"github.com/rodriez/runnest"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentCollection_Filter(t *testing.T) {
	runnest.TestCase{
		Name: "Given a collection with integers When filter odd numbers Then get a new collection with all even integers",
		Given: func() interface{} {
			return gollection.NewConcurrentCollection(3, "a", "b", "c", "d", "e", "f")
		},
		When: func(req interface{}) (interface{}, error) {
			collection := req.(gollection.Collection)
			filtered := collection.Filter(func(e *gollection.Element) bool {
				return strings.Contains("aeiou", e.String())
			})

			return filtered, nil
		},
		Then: func(t *testing.T, resp interface{}, e error) {
			collection := resp.(gollection.Collection)

			assert.False(t, collection.Empty())
			assert.Equal(t, 2, collection.Length())
		},
	}.Run(t)
}

func TestConcurrentCollection_Map(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 6 elements When map is called Then return a new collection with all elements converted",
			Given: func() interface{} {
				return gollection.NewConcurrentCollection(4, 50.5, 53.3, 71.2, 25, 5, 3, 2)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Map(func(e *gollection.Element) gollection.Element {
					if e.Float64() > 0 {
						return gollection.NewElement(fmt.Sprintf("%v", e.Float64()))
					} else if e.Int() > 0 {
						return gollection.NewElement(fmt.Sprintf("%v", e.Int()))
					} else {
						return *e
					}

				}), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 7, collection.Length())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestConcurrentCollection_Sort(t *testing.T) {
	runnest.TestCase{
		Name: "Given a collection with 6 elements When sort is called Then return the sorted collection",
		Given: func() interface{} {
			return gollection.NewConcurrentCollection(2, 50.5, 53.3, 71.2, 25, 5, 3, 2)
		},
		When: func(req interface{}) (interface{}, error) {
			collection := req.(gollection.Collection)

			return collection.Sort(func(e1, e2 *gollection.Element) bool {
				return (e1.Float64() > e2.Float64() && e1.Int() == e2.Int()) || (e1.Int() > e2.Int() && e1.Float64() == e2.Float64())
			}), nil
		},
		Then: func(t *testing.T, resp interface{}, e error) {
			collection := resp.(gollection.Collection)

			assert.Equal(t, float64(71.2), collection.Get(0).Float64())
			assert.Equal(t, float64(53.3), collection.Get(1).Float64())
			assert.Equal(t, float64(50.5), collection.Get(2).Float64())
			assert.Equal(t, 25, collection.Get(3).Int())
			assert.Equal(t, 5, collection.Get(4).Int())
			assert.Equal(t, 3, collection.Get(5).Int())
			assert.Equal(t, 2, collection.Get(6).Int())
		},
	}.Run(t)
}

func TestConcurrentCollection_SubCollection(t *testing.T) {
	runnest.TestCase{
		Name: "Given a collection with 6 elements When SubCollection is called Then return a new collection with wanted values",
		Given: func() interface{} {
			return gollection.NewConcurrentCollection(2, 50.5, 53.3, 71.2, 25, 5, 3, 2)
		},
		When: func(req interface{}) (interface{}, error) {
			collection := req.(gollection.Collection)

			return collection.SubCollection(3, 5), nil
		},
		Then: func(t *testing.T, resp interface{}, e error) {
			collection := resp.(gollection.Collection)

			assert.Equal(t, 25, collection.Get(0).Int())
			assert.Equal(t, 5, collection.Get(1).Int())
		},
	}.Run(t)
}
