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
	FLoat64    float64
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
	FLoat64    *float64
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
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	verifyPrimitiveData(t, &data.Array[0])
	verifyPrimitivePointerData(t, data.PointerArray[0])
}

func TestOverrideByType(t *testing.T) {
	configuration := NewConfiguration()
	stringValue := "123456"
	intValue := 24
	int16Value := int16(50)
	int32Value := int32(31)
	int64Value := int64(1234)
	int8Value := int8(36)
	configuration.Override("", stringValue).
		Override(int8(1), int8Value).
		Override(1, intValue).
		Override(int16(1), int16Value).
		Override(int32(1), int32Value).
		Override(int64(1), int64Value)
	data := ComplexData{}

	InflateWithConfiguration(&data, configuration)

	if len(data.Array) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if len(data.PointerArray) != 1 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	primitiveData := data.Array[0]
	if primitiveData.String != stringValue {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int != intValue {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int8 != int8Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int16 != int16Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int32 != int32Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if primitiveData.Int64 != int64Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	primitivePointerData := data.PointerArray[0]
	if *primitivePointerData.String != stringValue {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int != intValue {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int8 != int8Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int16 != int16Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int32 != int32Value {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *primitivePointerData.Int64 != int64Value {
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
	if data.FLoat64 != 1.0 {
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
	if *data.FLoat64 != 1.0 {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Complex64 != complex(1, 1) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
	if *data.Complex128 != complex(float64(1), float64(1)) {
		t.Fatal("Dummie didn't fill the correct data.")
	}
}
