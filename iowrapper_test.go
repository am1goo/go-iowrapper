package iowrapper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

type teststruct struct {
	bytes []byte
	rune rune
	int8 int8
	int16 int16
	int32 int32
	int64 int64
	uint8 uint8
	uint16 uint16
	uint32 uint32
	uint64 uint64
	float32 float32
	float64 float64
	complex64 complex64
	complex128 complex128
	string string
}

type testerrors []error

func (t *testerrors) add(err error) {
	*t = append(*t, err)
}

func (t testerrors) Error() string {
	if len(t) > 0 {
		return t[0].Error()
	} else {
		return ""
	}
}


func checkError(err error, errs *testerrors) {
	if err != nil {
		errs.add(err)
	}
}

func checkErrors(errs *testerrors) error {
	if *errs == nil {
		return nil
	} else if len(*errs) > 0 {
		return *errs
	} else {
		return nil
	}
}

func (t *teststruct) Read(r *Reader) error {
	errs := new(testerrors)

	var err error
	t.bytes, err = r.ReadBytesAndCount()
	checkError(err, errs)
	t.rune, err = r.ReadRune()
	checkError(err, errs)
	t.int8, err = r.ReadInt8()
	checkError(err, errs)
	t.int16, err = r.ReadInt16()
	checkError(err, errs)
	t.int32, err = r.ReadInt32()
	checkError(err, errs)
	t.int64, err = r.ReadInt64()
	checkError(err, errs)
	t.uint8, err = r.ReadUInt8()
	checkError(err, errs)
	t.uint16, err = r.ReadUInt16()
	checkError(err, errs)
	t.uint32, err =  r.ReadUInt32()
	checkError(err, errs)
	t.uint64, err =  r.ReadUInt64()
	checkError(err, errs)
	t.float32, err = r.ReadFloat32()
	checkError(err, errs)
	t.float64, err = r.ReadFloat64()
	checkError(err, errs)
	t.complex64, err = r.ReadComplex64()
	checkError(err, errs)
	t.complex128, err = r.ReadComplex128()
	checkError(err, errs)
	t.string, err = r.ReadString()
	checkError(err, errs)

	return checkErrors(errs)
}

func (t *teststruct) Write(w *Writer) error {
	errs := new(testerrors)

	err := w.WriteBytesAndCount(t.bytes)
	checkError(err, errs)
	err = w.WriteRune(t.rune)
	checkError(err, errs)
	err = w.WriteInt8(t.int8)
	checkError(err, errs)
	err = w.WriteInt16(t.int16)
	checkError(err, errs)
	err = w.WriteInt32(t.int32)
	checkError(err, errs)
	err = w.WriteInt64(t.int64)
	checkError(err, errs)
	err = w.WriteUInt8(t.uint8)
	checkError(err, errs)
	err = w.WriteUInt16(t.uint16)
	checkError(err, errs)
	err = w.WriteUInt32(t.uint32)
	checkError(err, errs)
	err = w.WriteUInt64(t.uint64)
	checkError(err, errs)
	err = w.WriteFloat32(t.float32)
	checkError(err, errs)
	err = w.WriteFloat64(t.float64)
	checkError(err, errs)
	err = w.WriteComplex64(t.complex64)
	checkError(err, errs)
	err = w.WriteComplex128(t.complex128)
	checkError(err, errs)
	err = w.WriteString(t.string)
	checkError(err, errs)

	return checkErrors(errs)
}

func createRandomBytes(len int) []byte {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(rand.Intn(math.MaxUint8))
	}
	return b
}

var testvar = teststruct{
	bytes:      createRandomBytes(99),
	rune:       '五',
	int8:       int8(rand.Intn(math.MaxInt8)),
	int16:      int16(rand.Intn(math.MaxInt16)),
	int32:      rand.Int31(),
	int64:      rand.Int63(),
	uint8:      uint8(rand.Intn(math.MaxUint8)),
	uint16:     uint16(rand.Intn(math.MaxUint16)),
	uint32:     rand.Uint32(),
	uint64:     rand.Uint64(),
	float32:    rand.Float32(),
	float64:    rand.Float64(),
	complex64:  complex(rand.Float32(), rand.Float32()),
	complex128: complex(rand.Float64(), rand.Float64()),
	string:     "This Is A Test String",
}

func TestAverage(t *testing.T) {
	var wb bytes.Buffer
	wr := NewWriter(&wb, binary.LittleEndian)
	err := testvar.Write(wr)
	if err != nil {
		t.Error(err)
	}

	var resultvar teststruct
	rb := bytes.NewBuffer(wb.Bytes())
	rw := NewReader(rb, binary.LittleEndian)
	err = resultvar.Read(rw)
	if err != nil {
		t.Error(err)
	}

	assertEquals(t, resultvar.bytes, testvar.bytes)
	assertEquals(t, resultvar.rune, testvar.rune)
	assertEquals(t, resultvar.int8, testvar.int8)
	assertEquals(t, resultvar.int16, testvar.int16)
	assertEquals(t, resultvar.int32, testvar.int32)
	assertEquals(t, resultvar.int64, testvar.int64)
	assertEquals(t, resultvar.uint8, testvar.uint8)
	assertEquals(t, resultvar.uint16, testvar.uint16)
	assertEquals(t, resultvar.uint32, testvar.uint32)
	assertEquals(t, resultvar.uint64, testvar.uint64)
	assertEquals(t, resultvar.float32, testvar.float32)
	assertEquals(t, resultvar.float64, testvar.float64)
	assertEquals(t, resultvar.complex64, testvar.complex64)
	assertEquals(t, resultvar.complex128, testvar.complex128)
	assertEquals(t, resultvar.string, testvar.string)
}

func assertEquals(t *testing.T, actual interface{}, expect interface{}) {
	actualType := reflect.TypeOf(actual)
	expectType := reflect.TypeOf(expect)
	if actualType != expectType {
		t.Error(fmt.Sprintf("type mismatch, actual=%v, expect=%v", actualType, expectType))
	}

	switch actualType.Kind() {
	case reflect.Slice:
		actualArray := sliceToArray(actual)
		expectArray := sliceToArray(expect)
		assertEquals(t, actualArray, expectArray)
		break
	default:
		if !reflect.DeepEqual(actual, expect) {
			t.Error(fmt.Sprintf("value mismatch, actual=%v, expect=%v", actual, expect))
		}
		break
	}
}

func sliceToArray(in interface{}) interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() != reflect.Slice {
		panic("not a slice")
	}
	t := reflect.ArrayOf(s.Len(), s.Type().Elem())
	a := reflect.New(t).Elem()
	reflect.Copy(a, s)
	return a.Interface()
}
