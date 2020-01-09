package iowrapper

import (
	"bytes"
	"encoding/binary"
)

func NewReader(buf *bytes.Buffer, order binary.ByteOrder) *Reader {
	return &Reader{buf: buf, order: order}
}

func NewWriter(buf *bytes.Buffer, order binary.ByteOrder) *Writer {
	return &Writer{buf: buf, order: order}
}