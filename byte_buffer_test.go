package gobytebuffer

import (
	"encoding/binary"
	"testing"
)

func TestSerializer_ReadByte_WriteByte(t *testing.T) {
	value := byte(254)

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteByte(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadByte()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadShort_WriteShort(t *testing.T) {
	value := 10

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteShort(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadShort()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadInt_WriteInt(t *testing.T) {
	value := 10

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteInt(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadInt()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadLong_WriteLong(t *testing.T) {
	value := 29914

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteLong(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadLong()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadFloat32_WriteFloat32(t *testing.T) {
	value := float32(3.981)

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteFloat32(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadFloat32()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadFloat64_WriteFloat64(t *testing.T) {
	value := 10.512

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteFloat64(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadFloat64()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadBool_WriteBool(t *testing.T) {
	value := true

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteBool(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadBool()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}

func TestSerializer_ReadString_WriteString(t *testing.T) {
	value := "awesome text!"

	for _, endianness := range []binary.ByteOrder{binary.LittleEndian, binary.BigEndian} {
		writer := ByteBuffer{
			Endianness: endianness,
		}

		writer.WriteString(value)

		reader := ByteBuffer{
			Endianness: endianness,
		}

		reader.Write(writer.GetBytes())

		rValue := reader.ReadString()

		if rValue != value {
			t.Errorf("Expected value of %v, but got %v, endiannesss: %s", value, rValue, endianness)
		}
	}
}