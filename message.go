package tower

type Message struct {
	DataLen uint32 // message's length
	Id      uint32 // message's id
	Data    []byte // message's content
}

//NewMsgPackage create a message package instance
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:      id,
		Data:    data,
	}
}

// GetDataLen get message data's length
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetMsgId
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// GetData get message content
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetDataLen
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// SetMsgId
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

// SetData set message's data content
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
