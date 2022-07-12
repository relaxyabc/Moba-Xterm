package register

import (
	"math/big"
	"strings"
)

const Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

var TableCache = make(map[int]string, 128)

func init() {
	split := strings.Split(Table, "")
	for idx, val := range split {
		TableCache[idx] = val
	}
}

func VariantBase64Encode(byteArr []byte) string {
	builder := strings.Builder{}
	length := len(byteArr)
	blockCount := length / 3
	leftBytes := length % 3
	for idx := 0; idx < blockCount; idx++ {
		start := 3 * idx
		subBytes := byteArr[start : start+3]
		intVal := convertByteToInt(subBytes)
		block := TableCache[intVal&0x3f]
		block += TableCache[(intVal>>6)&0x3f]
		block += TableCache[(intVal>>12)&0x3f]
		block += TableCache[(intVal>>18)&0x3f]
		builder.WriteString(block)
	}
	if leftBytes == 0 {
		// nothing to do
	} else if leftBytes == 1 {
		subBytes := byteArr[3*blockCount : length]
		block := rightMoveSix(subBytes)
		builder.WriteString(block)
	} else {
		subBytes := byteArr[3*blockCount : length]
		intVal := convertByteToInt(subBytes)
		block := rightMoveSix(subBytes)
		block += TableCache[(intVal>>12)&0x3f]
	}
	return builder.String()
}

func rightMoveSix(byteArr []byte) string {
	intVal := convertByteToInt(byteArr)
	block := TableCache[intVal&0x3f]
	block += TableCache[(intVal>>6)&0x3f]
	return block
}

//z := new(big.Int)
//z.SetBytes(byteArr)
//return int(z.Uint64())
func convertByteToInt(byteArr []byte) int {
	half := 2
	length := len(byteArr)
	for idx := 0; idx < length/half; idx++ {
		temp := byteArr[idx]
		byteArr[idx] = byteArr[length-idx-1]
		byteArr[length-idx-1] = temp
	}
	z := new(big.Int)
	z.SetBytes(byteArr)
	return int(z.Uint64())
}
