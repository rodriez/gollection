package gollection_test

import (
	"testing"

	"github.com/rodriez/gollection"
	"github.com/rodriez/runnest"
	"github.com/stretchr/testify/assert"
)

func TestConstructors(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Testing NewStringCollection",
			Given: func() interface{} {
				return []string{"str1", "str2", "str3"}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]string)
				return gollection.NewStringCollection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewInt64Collection",
			Given: func() interface{} {
				return []int64{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]int64)
				return gollection.NewInt64Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewInt32Collection",
			Given: func() interface{} {
				return []int32{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]int32)
				return gollection.NewInt32Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewInt16Collection",
			Given: func() interface{} {
				return []int16{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]int16)
				return gollection.NewInt16Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewInt8Collection",
			Given: func() interface{} {
				return []int8{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]int8)
				return gollection.NewInt8Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewIntCollection",
			Given: func() interface{} {
				return []int{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]int)
				return gollection.NewIntCollection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewFloat64Collection",
			Given: func() interface{} {
				return []float64{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]float64)
				return gollection.NewFloat64Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewFloat32Collection",
			Given: func() interface{} {
				return []float32{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]float32)
				return gollection.NewFloat32Collection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewByteCollection",
			Given: func() interface{} {
				return []byte{1, 2, 3}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]byte)
				return gollection.NewByteCollection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
		{
			Name: "Testing NewBoolCollection",
			Given: func() interface{} {
				return []bool{true, false, true}
			},
			When: func(req interface{}) (interface{}, error) {
				arr := req.([]bool)
				return gollection.NewBoolCollection(arr...), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				collection := resp.(gollection.Collection)

				assert.Equal(t, 3, collection.Length())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}
