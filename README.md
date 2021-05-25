# gobytebuffer

Simple byte buffer implementation written in GO

## Example usage
```go
writer := gobytebuffer.ByteBuffer{
	Endianness: binary.BigEndian,
}

writer.WriteInt(4321)
writer.WriteString("example message")
writer.WriteBool(true)


reader := gobytebuffer.ByteBuffer{
	Endianness: binary.BigEndian,
}

reader.Write(writer.GetBytes())

fmt.Println(reader.ReadInt()) // 4321
fmt.Println(reader.ReadString()) // example message
fmt.Println(reader.ReadBool()) // true
```