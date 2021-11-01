package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 封包
func Encode(message string) ([]byte, error) {
	length := int32(len(message)) // 32位 占4字节
	packet := new(bytes.Buffer)
	// 包头：message的长度
	err := binary.Write(packet, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 包体：message
	err = binary.Write(packet, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return packet.Bytes(), nil
}

// 解包
func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4) // 读取前4个字节
	if err != nil {
		return "", err
	}
	lengthBuff := bytes.NewBuffer(lengthByte) // 用这个字节切片 创建一个用于读取数据的buffer；

	var length int32
	err = binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// 当前buffer可以读取的字节数 小于 头部规定的数据，说明数据丢失，返回error
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	packet := make([]byte, 4+length)
	_, err = reader.Read(packet)
	if err != nil {
		return "", err
	}
	return string(packet[4:]), nil
}
