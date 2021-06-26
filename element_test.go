package gollection_test

import (
	"testing"

	"github.com/rodriez/gollection"
	"github.com/rodriez/runnest"
	"github.com/stretchr/testify/assert"
)

func TestElement_String(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given string value When The constructor is called Then the type is string and the String() method return the value",
			Given: func() interface{} {
				return "A very large name"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_STRING, element.Type())
				assert.Equal(t, "A very large name", element.String())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return empty string",
			Given: func() interface{} {
				return 87987.9
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_STRING, element.Type())
				assert.Equal(t, "", element.String())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Int64(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given int64 value When The constructor is called Then the type is int64 and the Int64() method return the value",
			Given: func() interface{} {
				return int64(547865)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_INT64, element.Type())
				assert.Equal(t, int64(547865), element.Int64())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_INT64, element.Type())
				assert.Equal(t, int64(0), element.Int64())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Int32(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given int32 value When The constructor is called Then the type is int32 and the Int32() method return the value",
			Given: func() interface{} {
				return int32(547865)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_INT32, element.Type())
				assert.Equal(t, int32(547865), element.Int32())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_INT32, element.Type())
				assert.Equal(t, int32(0), element.Int32())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Int16(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given int16 value When The constructor is called Then the type is int16 and the Int16() method return the value",
			Given: func() interface{} {
				return int16(5478)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_INT16, element.Type())
				assert.Equal(t, int16(5478), element.Int16())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_INT16, element.Type())
				assert.Equal(t, int16(0), element.Int16())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Int8(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given int8 value When The constructor is called Then the type is int8 and the Int8() method return the value",
			Given: func() interface{} {
				return int8(54)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_INT8, element.Type())
				assert.Equal(t, int8(54), element.Int8())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_INT8, element.Type())
				assert.Equal(t, int8(0), element.Int8())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Int(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given int value When The constructor is called Then the type is int and the Int() method return the value",
			Given: func() interface{} {
				return int(54)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_INT, element.Type())
				assert.Equal(t, int(54), element.Int())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "797.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_INT, element.Type())
				assert.Equal(t, int(0), element.Int())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Float64(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given float64 value When The constructor is called Then the type is float64 and the Float64() method return the value",
			Given: func() interface{} {
				return float64(547865)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_FLOAT64, element.Type())
				assert.Equal(t, float64(547865), element.Float64())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_FLOAT64, element.Type())
				assert.Equal(t, float64(0), element.Float64())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Float32(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given float32 value When The constructor is called Then the type is float32 and the Float32() method return the value",
			Given: func() interface{} {
				return float32(547865)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_FLOAT32, element.Type())
				assert.Equal(t, float32(547865), element.Float32())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_FLOAT32, element.Type())
				assert.Equal(t, float32(0), element.Float32())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Byte(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given byte value When The constructor is called Then the type is byte and the Byte() method return the value",
			Given: func() interface{} {
				return byte(54)
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_BYTE, element.Type())
				assert.Equal(t, byte(54), element.Byte())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_BYTE, element.Type())
				assert.Equal(t, byte(0), element.Byte())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Bool(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given bool value When The constructor is called Then the type is bool and the Bool() method return the value",
			Given: func() interface{} {
				return true
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.Equal(t, gollection.TYPE_BOOL, element.Type())
				assert.Equal(t, true, element.Bool())
			},
		},
		{
			Name: "Given random value When The constructor is called Then the type is not string and the String() method return 0",
			Given: func() interface{} {
				return "87987.9"
			},
			When: func(req interface{}) (interface{}, error) {
				return gollection.NewElement(req), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				element := resp.(gollection.Element)

				assert.NotEqual(t, gollection.TYPE_BOOL, element.Type())
				assert.Equal(t, false, element.Bool())
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Value(t *testing.T) {
	type Mock struct{}

	runnest.TestCase{
		Name: "Given random value When The constructor is called Then the Value() method return the value",
		Given: func() interface{} {
			return Mock{}
		},
		When: func(req interface{}) (interface{}, error) {
			return gollection.NewElement(req), nil
		},
		Then: func(t *testing.T, resp interface{}, e error) {
			element := resp.(gollection.Element)

			assert.Equal(t, Mock{}, element.Value())
		},
	}.Run(t)
}

func TestElement_Fill(t *testing.T) {
	type Mock struct {
		Value string
	}
	testCases := []runnest.TestCase{
		{
			Name: "Given random value When Fill() is called Then copy the value in the interface passed in",
			Given: func() interface{} {
				return Mock{"A simple string"}
			},
			When: func(req interface{}) (interface{}, error) {
				var mock Mock

				element := gollection.NewElement(req)
				element.Fill(&mock)

				return mock, nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				assert.Equal(t, Mock{"A simple string"}, resp)
			},
		},
		{
			Name: "Given random value When Fill() is called with a wrong type value Then return incompatible types error",
			Given: func() interface{} {
				return "A simple string"
			},
			When: func(req interface{}) (interface{}, error) {
				var mock Mock
				element := gollection.NewElement(req)

				return nil, element.Fill(&mock)
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				assert.Error(t, e, "element value not compatible with type string")
			},
		},
		{
			Name: "Given random value When Fill() is called with a non pointer value Then return non pointer error",
			Given: func() interface{} {
				return "A simple string"
			},
			When: func(req interface{}) (interface{}, error) {
				var mock Mock
				element := gollection.NewElement(req)

				return nil, element.Fill(mock)
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				assert.Error(t, e, "non-pointer gollection_test.Mock{}")
			},
		},
		{
			Name: "Given random value When Fill() is called with a non pointer value Then return non pointer error",
			Given: func() interface{} {
				return "A simple string"
			},
			When: func(req interface{}) (interface{}, error) {
				element := gollection.NewElement(req)

				return nil, element.Fill(nil)
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				assert.Error(t, e, "non-pointer gollection_test.Mock{}")
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}

func TestElement_Empty(t *testing.T) {
	testCases := []runnest.TestCase{
		{
			Name: "Given an empty element When IsEmpty is called Then should return true",
			Given: func() interface{} {
				return gollection.Empty()
			},
			When: func(req interface{}) (interface{}, error) {
				e := req.(gollection.Element)
				return e.ISEmpty(), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				isEmpty := resp.(bool)

				assert.True(t, isEmpty)
			},
		},
		{
			Name: "Given an non empty element When IsEmpty is called Then should return false",
			Given: func() interface{} {
				return gollection.NewString("Something")
			},
			When: func(req interface{}) (interface{}, error) {
				e := req.(gollection.Element)
				return e.ISEmpty(), nil
			},
			Then: func(t *testing.T, resp interface{}, e error) {
				isEmpty := resp.(bool)

				assert.False(t, isEmpty)
			},
		},
	}

	runnest.NewRunest(t).Run(testCases)
}
