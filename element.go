package gollection

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	TYPE_STRING  = "string"
	TYPE_INT64   = "int64"
	TYPE_INT32   = "int32"
	TYPE_INT16   = "int16"
	TYPE_INT8    = "int8"
	TYPE_INT     = "int"
	TYPE_FLOAT64 = "float64"
	TYPE_FLOAT32 = "float32"
	TYPE_BYTE    = "uint8"
	TYPE_BOOL    = "bool"
)

type Element struct {
	typ   string
	value interface{}
}

func NewElement(e interface{}) Element {
	return Element{
		typ:   reflect.TypeOf(e).Name(),
		value: e,
	}
}

func NewString(e string) Element {
	return Element{
		typ:   TYPE_STRING,
		value: e,
	}
}

func NewInt64(e int64) Element {
	return Element{
		typ:   TYPE_INT64,
		value: e,
	}
}

func NewInt32(e int32) Element {
	return Element{
		typ:   TYPE_INT32,
		value: e,
	}
}

func NewInt16(e int16) Element {
	return Element{
		typ:   TYPE_INT16,
		value: e,
	}
}

func NewInt8(e int8) Element {
	return Element{
		typ:   TYPE_INT8,
		value: e,
	}
}

func NewInt(e int) Element {
	return Element{
		typ:   TYPE_INT,
		value: e,
	}
}

func NewFloat64(e float64) Element {
	return Element{
		typ:   TYPE_FLOAT64,
		value: e,
	}
}

func NewFloat32(e float32) Element {
	return Element{
		typ:   TYPE_FLOAT32,
		value: e,
	}
}

func NewByte(e byte) Element {
	return Element{
		typ:   TYPE_BYTE,
		value: e,
	}
}

func NewBool(e bool) Element {
	return Element{
		typ:   TYPE_BOOL,
		value: e,
	}
}

func (e Element) Type() string {
	return e.typ
}

func (e Element) String() string {
	if e.Type() == TYPE_STRING {
		return e.value.(string)
	}

	return ""
}

func (e Element) Int64() int64 {
	if e.Type() == TYPE_INT64 {
		return e.value.(int64)
	}

	return 0
}

func (e Element) Int32() int32 {
	if e.Type() == TYPE_INT32 {
		return e.value.(int32)
	}

	return 0
}

func (e Element) Int16() int16 {
	if e.Type() == TYPE_INT16 {
		return e.value.(int16)
	}

	return 0
}

func (e Element) Int8() int8 {
	if e.Type() == TYPE_INT8 {
		return e.value.(int8)
	}

	return 0
}

func (e Element) Int() int {
	if e.Type() == TYPE_INT {
		return e.value.(int)
	}

	return 0
}

func (e Element) Float64() float64 {
	if e.Type() == TYPE_FLOAT64 {
		return e.value.(float64)
	}

	return 0
}

func (e Element) Float32() float32 {
	if e.Type() == TYPE_FLOAT32 {
		return e.value.(float32)
	}

	return 0
}

func (e Element) Byte() byte {
	if e.Type() == TYPE_BYTE {
		return e.value.(byte)
	}

	return 0
}

func (e Element) Bool() bool {
	if e.Type() == TYPE_BOOL {
		return e.value.(bool)
	}

	return false
}

func (e Element) Value() interface{} {
	return e.value
}

func (e Element) Fill(target interface{}) error {
	if target == nil {
		return errors.New("nil target")
	}

	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", value.Type())
	}

	value = value.Elem()
	if reflect.TypeOf(e.value) != value.Type() {
		return fmt.Errorf("element value not compatible with type %v", value.Type())
	}

	elementValue := reflect.ValueOf(e.value)
	value.Set(elementValue)

	return nil
}
