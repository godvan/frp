package protocol

type Packet struct {
	Type    uint8
	Payload []byte
}

func (p *Packet) Encode() []byte {
	return EncodePacket(p.Payload)
}

func Decode(data []byte) (*Packet, error) {
	payload, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	return &Packet{
		Type:    0, // 根据实际类型设置
		Payload: payload,
	}, nil
}
