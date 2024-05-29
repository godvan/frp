package protocol

import (
	"bytes"
	"encoding/binary"
)

func EncodePacket(packet []byte) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, uint32(len(packet)))
	buf.Write(packet)
	return buf.Bytes()
}

func DecodePacket(data []byte) ([]byte, error) {
	var length uint32
	err := binary.Read(bytes.NewReader(data), binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	return data[4:], nil
}
