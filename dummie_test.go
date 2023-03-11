package dummie

import (
	"testing"
)

func TestInflateInt8(t *testing.T) {
	var intData int8
	err := Inflate(&intData)
	if err != nil {
		t.Fatalf("Dummie inflate data error: %+v.", err.Error())
	}
	if intData != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}

type PrimitiveData struct {
	Bool       bool
	Int8       int8
	Int16      int16
	Int32      int32
	Int        int
	Int64      int64
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint       uint
	Uint64     uint64
	String     string
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
}

func TestInflatePrimitiveData(t *testing.T) {
	data := PrimitiveData{}
	Inflate(&data)

	verifyPrimitiveData(t, &data)
}

type PrimitivePointerData struct {
	Bool       *bool
	Int8       *int8
	Int16      *int16
	Int32      *int32
	Int        *int
	Int64      *int64
	Uint8      *uint8
	Uint16     *uint16
	Uint32     *uint32
	Uint       *uint
	Uint64     *uint64
	String     *string
	Float32    *float32
	Float64    *float64
	Complex64  *complex64
	Complex128 *complex128
}

func TestInflatePrimitivePointerData(t *testing.T) {
	data := PrimitivePointerData{}
	Inflate(&data)

	verifyPrimitivePointerData(t, &data)
}

type ComplexData struct {
	Array        []PrimitiveData
	PointerArray []*PrimitivePointerData
}

func TestInflateComplexData(t *testing.T) {
	data := ComplexData{}

	Inflate(&data)

	if len(data.Array) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyPrimitiveData(t, &data.Array[0])
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyPrimitivePointerData(t, data.PointerArray[0])
}

var mockData = PrimitiveData{
	Bool:       false,
	String:     "123456",
	Int:        24,
	Int8:       int8(36),
	Int16:      int16(50),
	Int32:      int32(31),
	Int64:      int64(1234),
	Uint:       24,
	Uint8:      uint8(36),
	Uint16:     uint16(50),
	Uint32:     uint32(31),
	Uint64:     uint64(1234),
	Float32:    float32(3.0),
	Float64:    float64(41.0),
	Complex64:  complex(float32(123.0), -float32(234.4)),
	Complex128: complex(float64(555), float64(666)),
}

func TestOverrideByType(t *testing.T) {
	configuration := NewConfiguration()
	configuration.Override("", mockData.String).
		Override(true, mockData.Bool).
		Override(1, mockData.Int).
		Override(int8(1), mockData.Int8).
		Override(int16(1), mockData.Int16).
		Override(int32(1), mockData.Int32).
		Override(int64(1), mockData.Int64).
		Override(uint(1), mockData.Uint).
		Override(uint8(1), mockData.Uint8).
		Override(uint16(1), mockData.Uint16).
		Override(uint32(1), mockData.Uint32).
		Override(uint64(1), mockData.Uint64).
		Override(float32(1), mockData.Float32).
		Override(float64(1), mockData.Float64).
		Override(complex(float32(1), 1), mockData.Complex64).
		Override(complex(float64(1), float64(1)), mockData.Complex128)
	data := ComplexData{}

	InflateWithConfiguration(&data, configuration)

	if len(data.Array) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyOverridePrimitiveData(t, data.Array[0])
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyOverridePrimitivePoinerData(t, *data.PointerArray[0])
}

func TestOverrideStruct(t *testing.T) {
	configuration := NewConfiguration()
	configuration.Override(PrimitiveData{}, mockData)
	data := ComplexData{}

	InflateWithConfiguration(&data, configuration)

	if len(data.Array) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyOverridePrimitiveData(t, data.Array[0])
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyPrimitivePointerData(t, data.PointerArray[0])
}

func TestOverrideByPath(t *testing.T) {
	configuration := NewConfiguration()
	configuration.OverrideWithPath("Array", mockData).
		OverrideWithPath("String", mockData.String).
		OverrideWithPath("Bool", mockData.Bool).
		Override(1, mockData.Int).
		Override(int8(1), mockData.Int8).
		Override(int16(1), mockData.Int16).
		Override(int32(1), mockData.Int32).
		Override(int64(1), mockData.Int64).
		Override(uint(1), mockData.Uint).
		Override(uint8(1), mockData.Uint8).
		Override(uint16(1), mockData.Uint16).
		Override(uint32(1), mockData.Uint32).
		Override(uint64(1), mockData.Uint64).
		Override(float32(1), mockData.Float32).
		Override(float64(1), mockData.Float64).
		Override(complex(float32(1), 1), mockData.Complex64).
		Override(complex(float64(1), float64(1)), mockData.Complex128)
	data := ComplexData{}

	InflateWithConfiguration(&data, configuration)

	if len(data.Array) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}

	verifyOverridePrimitiveData(t, data.Array[0])
	verifyOverridePrimitivePoinerData(t, *data.PointerArray[0])
}

func verifyOverridePrimitiveData(t *testing.T, primitiveData PrimitiveData) {
	if primitiveData.Bool != mockData.Bool {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.String != mockData.String {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int != mockData.Int {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int8 != mockData.Int8 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int16 != mockData.Int16 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int32 != mockData.Int32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int64 != mockData.Int64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Uint != mockData.Uint {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Uint8 != mockData.Uint8 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Uint16 != mockData.Uint16 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Uint32 != mockData.Uint32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Uint64 != mockData.Uint64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Float32 != mockData.Float32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Float64 != mockData.Float64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Complex64 != mockData.Complex64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Complex128 != mockData.Complex128 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}

func verifyOverridePrimitivePoinerData(t *testing.T, primitivePointerData PrimitivePointerData) {
	if *primitivePointerData.String != mockData.String {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Bool != mockData.Bool {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int != mockData.Int {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int8 != mockData.Int8 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int16 != mockData.Int16 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int32 != mockData.Int32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int64 != mockData.Int64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Uint != mockData.Uint {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Uint8 != mockData.Uint8 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Uint16 != mockData.Uint16 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Uint32 != mockData.Uint32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Uint64 != mockData.Uint64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Float32 != mockData.Float32 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Float64 != mockData.Float64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Complex64 != mockData.Complex64 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Complex128 != mockData.Complex128 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}

func verifyPrimitiveData(t *testing.T, data *PrimitiveData) {
	if data.Bool != true {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Int8 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Int16 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Int32 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Int64 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Uint8 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Uint16 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Uint32 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Uint != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Uint64 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.String != "String" {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Float32 != 1.0 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Float64 != 1.0 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Complex64 != complex(1, 1) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if data.Complex128 != complex(float64(1), float64(1)) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}

func verifyPrimitivePointerData(t *testing.T, data *PrimitivePointerData) {
	if *data.Bool != true {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Int8 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Int16 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Int32 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Int64 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Uint8 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Uint16 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Uint32 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Uint != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Uint64 != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.String != "String" {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Float32 != 1.0 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Float64 != 1.0 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Complex64 != complex(1, 1) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Complex128 != complex(float64(1), float64(1)) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}
