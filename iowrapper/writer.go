package iowrapper

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	buf *bytes.Buffer
	order binary.ByteOrder
}


func (w Writer) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}

func (w Writer) WriteByte(b byte) error {
	return w.buf.WriteByte(b)
}

func (w Writer) WriteBytes(b []byte) error {
	_, err := w.Write(b)
	return err
}

func (w Writer) WriteBytesAndCount(b []byte) error {
 	err := w.WriteUInt32(uint32(len(b)))
 	if err != nil {
 		return err
	}

	_, err = w.Write(b)
	return err
}

func (w Writer) WriteBool(b bool) error {
	return binary.Write(w, w.order, b)
}

func (w Writer) WriteRune(r rune) error {
	_, err := w.buf.WriteRune(r)
	return err
}

func (w Writer) WriteInt8(i int8) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteInt16(i int16) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteInt32(i int32) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteInt64(i int64) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteUInt8(i uint8) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteUInt16(i uint16) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteUInt32(i uint32) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteUInt64(i uint64) error {
	return binary.Write(w, w.order, i)
}

func (w Writer) WriteFloat32(f float32) error {
	return binary.Write(w, w.order, f)
}

func (w Writer) WriteFloat64(f float64) error {
	return binary.Write(w, w.order, f)
}

func (w Writer) WriteComplex64(f complex64) error {
	return binary.Write(w, w.order, f)
}

func (w Writer) WriteComplex128(f complex128) error {
	return binary.Write(w, w.order, f)
}

func (w Writer) WriteString(s string) error {
	b := []byte(s)
	err := w.WriteUInt32(uint32(len(b)))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func (w Writer) Bytes() []byte {
	return w.buf.Bytes()
}