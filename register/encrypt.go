package register

func EncryptBytes(key int, byteArr []byte) string {
	length := len(byteArr)
	encryptBytes := make([]byte, length, length*2)

	for i := 0; i < length; i++ {
		integer := int(byteArr[i]) ^ ((key >> 8) & 0xff)

		encryptBytes[i] = uint8(integer)

		tmp := encryptBytes[length-1]
		key = int(tmp)&key | 0x482D
	}
	return string(encryptBytes)
}
