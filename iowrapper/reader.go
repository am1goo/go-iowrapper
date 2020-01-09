package iowrapper

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Reader struct {
	buf *bytes.Buffer
	order binary.ByteOrder
}

func (r Reader) Read(p []byte) (n int, err error) {
	return r.buf.Read(p)
}

func (r Reader) ReadBytes(size int) ([]byte, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	b := make([]byte, size)
	_, err := r.Read(b)
	return b, err
}

func (r Reader) ReadBytesAndCount() ([]byte, error) {
	size, err := r.ReadUInt32()
	if err != nil {
		return nil, err
	}

	b := make([]byte, size)
	_, err = r.Read(b)
	return b, err
}

func (r Reader) ReadRune() (rune, error) {
	i, _, err := r.buf.ReadRune()
	return i, err
}

func (r Reader) ReadInt8() (int8, error) {
	i := new(int8)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadBool() (bool, error) {
	b := new(bool)
	err := binary.Read(r, r.order, b)
	return *b, err
}

func (r Reader) ReadInt16() (int16, error) {
	i := new(int16)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadInt32() (int32, error) {
	i := new(int32)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadInt64() (int64, error) {
	i := new(int64)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadUInt8() (uint8, error) {
	i := new(uint8)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadUInt16() (uint16, error) {
	i := new(uint16)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadUInt32() (uint32, error) {
	i := new(uint32)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadUInt64() (uint64, error) {
	i := new(uint64)
	err := binary.Read(r, r.order, i)
	return *i, err
}

func (r Reader) ReadFloat32() (float32, error) {
	f := new(float32)
	err := binary.Read(r, r.order, f)
	return *f, err
}

func (r Reader) ReadFloat64() (float64, error) {
	f := new(float64)
	err := binary.Read(r, r.order, f)
	return *f, err
}

func (r Reader) ReadComplex64() (complex64, error) {
	f := new(complex64)
	err := binary.Read(r, r.order, f)
	return *f, err
}

func (r Reader) ReadComplex128() (complex128, error) {
	f := new(complex128)
	err := binary.Read(r, r.order, f)
	return *f, err
}

func (r Reader) ReadString() (string, error) {
	size, err := r.ReadUInt32()
	if err != nil {
		return "", err
	}

	b := make([]byte, size)
	_, err = r.Read(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}