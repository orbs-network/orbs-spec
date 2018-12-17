package hexdumps

func stubByteArrayWithLength(length int, value byte) []byte {
	res := make([]byte, length)
	for i := 0; i<length; i++ {
		res[i] = value
	}
	return res
}