package gollection_test

import (
	"fmt"
	"testing"

	"github.com/rodriez/gollection"
	"github.com/rodriez/runnest"
	"github.com/stretchr/testify/assert"
)

func TestCollection_Add(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given 1 element When is added to the collection Then collection is not empty and get the element back",
			Given: func() interface{} {
				return "John"
			},
			When: func(req interface{}) (interface{}, error) {
				element := req.(string)
				collection := gollection.NewCollection().Add(gollection.NewElement(element))

				return collection, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.False(t, collection.Empty())
				assert.Equal(t, "John", collection.Get(0).String())
			},
		},
		{
			Name: "Given 2 element When are added to the collection Then the collection is not empty and get all elements back",
			Given: func() interface{} {
				return []string{"John", "Arthur"}
			},
			When: func(req interface{}) (interface{}, error) {
				list := req.([]string)
				collection := gollection.NewCollection().
					Add(gollection.NewElement(list[0])).
					Add(gollection.NewElement(list[1]))

				return collection, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.False(t, collection.Empty())
				assert.Equal(t, "John", collection.Get(0).String())
				assert.Equal(t, "Arthur", collection.Get(1).String())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Empty(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given an empty collection Then that collection is empty",
			Given: func() interface{} {
				return nil
			},
			When: func(req interface{}) (interface{}, error) {
				collection := gollection.NewCollection()
				return collection, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.True(t, collection.Empty())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Filter(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with integers When filter odd numbers Then get a new collection with all even integers",
			Given: func() interface{} {
				return gollection.NewCollection(50, 53, 71, 2, 4, 5)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Filter(func(e gollection.Element) bool {
					return e.Int()%2 == 0
				}), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.False(t, collection.Empty())
				assert.Equal(t, 50, collection.Get(0).Int())
				assert.Equal(t, 2, collection.Get(1).Int())
				assert.Equal(t, 4, collection.Get(2).Int())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Find(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with elements When search for True without offset Then return the first element that match with criteria",
			Given: func() interface{} {
				return gollection.NewCollection(50, true, false, "true", 4, true)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Find(func(e gollection.Element) bool {
					return e.Bool()
				}, 0), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.True(t, element.Bool())
			},
		},
		{
			Name: "Given a collection with elements When search for True with an offset of 3 Then return the first element that match with criteria",
			Given: func() interface{} {
				return gollection.NewCollection(50, true, false, "true", 4, true)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Find(func(e gollection.Element) bool {
					return e.Bool()
				}, 3), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.True(t, element.Bool())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_ForEach(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 3 elements When forEach is call Then apply the function for each element",
			Given: func() interface{} {
				return gollection.NewCollection(50.5, 53.3, 71.2)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				count := 0
				collection.ForEach(func(idx int, e gollection.Element) {
					count++
				})

				return count, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				count := resp.(int)

				assert.Equal(t, 3, count)
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Length(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 3 elements When Length is call Then return 3",
			Given: func() interface{} {
				return gollection.NewCollection(50.5, 53.3, 71.2)
			},
			When: func(req interface{}) (interface{}, error) {
				return req, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Map(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 6 elements When map is called Then return a new collection with all elements converted",
			Given: func() interface{} {
				return gollection.NewCollection(50.5, 53.3, 71.2, 25, 5, 3, 2)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Map(func(e gollection.Element) gollection.Element {
					if e.Float64() > 0 {
						return gollection.NewElement(fmt.Sprintf("%v", e.Float64()))
					} else if e.Int() > 0 {
						return gollection.NewElement(fmt.Sprintf("%v", e.Int()))
					} else {
						return e
					}

				}), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, "50.5", collection.Get(0).String())
				assert.Equal(t, "53.3", collection.Get(1).String())
				assert.Equal(t, "71.2", collection.Get(2).String())
				assert.Equal(t, "25", collection.Get(3).String())
				assert.Equal(t, "5", collection.Get(4).String())
				assert.Equal(t, "3", collection.Get(5).String())
				assert.Equal(t, "2", collection.Get(6).String())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Reduce(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 6 elements When reduce is called Then return the reduced value",
			Given: func() interface{} {
				return gollection.NewCollection(50.5, 53.3, 71.2, 25, 5, 3, 2)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Reduce(func(acc, e gollection.Element) gollection.Element {
					sum := acc.Float64()

					if e.Type() == gollection.TYPE_INT {
						sum += float64(e.Int())
					} else {
						sum += e.Float64()
					}

					return gollection.NewElement(sum)
				}, gollection.NewFloat64(0)), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				total := resp.(gollection.Element)

				assert.Equal(t, float64(210), total.Float64())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_ReverseFind(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with elements When search for True without offset Then return the first element that match with criteria",
			Given: func() interface{} {
				return gollection.NewCollection(50, true, false, "true", 4, true)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.ReverseFind(func(e gollection.Element) bool {
					return e.Bool()
				}, 0), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.True(t, element.Bool())
			},
		},
		{
			Name: "Given a collection with elements When search for True with an offset of 3 Then return the first element that match with criteria",
			Given: func() interface{} {
				return gollection.NewCollection(50, true, false, "true", 4, true)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.ReverseFind(func(e gollection.Element) bool {
					return e.Bool()
				}, 3), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.True(t, element.Bool())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestCollection_Sort(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given a collection with 6 elements When reduce is called Then return the reduced value",
			Given: func() interface{} {
				return gollection.NewCollection(50.5, 53.3, 71.2, 25, 5, 3, 2)
			},
			When: func(req interface{}) (interface{}, error) {
				collection := req.(gollection.Collection)

				return collection.Sort(func(e1, e2 gollection.Element) bool {
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
		},
	}

	runnest.NewRunest(t).Run(testCases)
}
