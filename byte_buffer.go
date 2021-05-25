package gobytebuffer

import (
	"bytes"
	"encoding/binary"
	"math"
)

type ByteBuffer struct {
	Endianness binary.ByteOrder
	Buffer bytes.Buffer
}

func (bbuf *ByteBuffer) GetBytes() []byte {
	return bbuf.Buffer.Bytes()
}

func (bbuf *ByteBuffer) Reset() {
	bbuf.Buffer.Reset()
}

func (bbuf *ByteBuffer) Read(len int) []byte {
	return bbuf.Buffer.Next(len)
}

func (bbuf *ByteBuffer) Write(value []byte) {
	bbuf.Buffer.Write(value)
}

func (bbuf *ByteBuffer) ReadByte() byte {
	return bbuf.Read(1)[0]
}

func (bbuf *ByteBuffer) WriteByte(value byte) {
	bbuf.Write([]byte{value})
}

func (bbuf *ByteBuffer) ReadShort() int {
	value := bbuf.Read(2)

	return int(bbuf.Endianness.Uint16(value))
}

func (bbuf *ByteBuffer) WriteShort(value int) {
	b := make([]byte, 2)

	bbuf.Endianness.PutUint16(b, uint16(value))
	bbuf.Write(b)
}

func (bbuf *ByteBuffer) ReadInt() int {
	value := bbuf.Read(4)

	return int(bbuf.Endianness.Uint32(value))
}

func (bbuf *ByteBuffer) WriteInt(value int) {
	b := make([]byte, 4)

	bbuf.Endianness.PutUint32(b, uint32(value))
	bbuf.Write(b)
}

func (bbuf *ByteBuffer) ReadLong() int {
	value := bbuf.Read(8)

	return int(bbuf.Endianness.Uint64(value))
}

func (bbuf *ByteBuffer) WriteLong(value int) {
	b := make([]byte, 8)

	bbuf.Endianness.PutUint64(b, uint64(value))
	bbuf.Write(b)
}

func (bbuf *ByteBuffer) ReadFloat32() float32 {
	value := bbuf.Read(4)

	return math.Float32frombits(bbuf.Endianness.Uint32(value))
}

func (bbuf *ByteBuffer) WriteFloat32(value float32) {
	b := make([]byte, 4)

	bbuf.Endianness.PutUint32(b, math.Float32bits(value))
	bbuf.Write(b)
}

func (bbuf *ByteBuffer) ReadFloat64() float64 {
	value := bbuf.Read(8)

	return math.Float64frombits(bbuf.Endianness.Uint64(value))
}

func (bbuf *ByteBuffer) WriteFloat64(value float64) {
	b := make([]byte, 8)

	bbuf.Endianness.PutUint64(b, math.Float64bits(value))
	bbuf.Write(b)
}

func (bbuf *ByteBuffer) ReadBool() bool {
	value := bbuf.ReadByte()

	if value == 0 {
		return false
	}

	return true
}

func (bbuf *ByteBuffer) WriteBool(value bool) {
	if value {
		bbuf.WriteByte(1)
	} else {
		bbuf.WriteByte(0)
	}
}

func (bbuf *ByteBuffer) ReadString() string {
	return string(bbuf.Read(bbuf.ReadShort()))
}

func (bbuf *ByteBuffer) WriteString(value string) {
	bbuf.WriteShort(len(value))
	bbuf.Write([]byte(value))
}