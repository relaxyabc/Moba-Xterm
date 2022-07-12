package register

import (
	"archive/zip"
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"mobaxterm/entity"
	"os"
	"regexp"
	"testing"
	"time"
)

func TestGenerateKey(t *testing.T) {
	license := entity.NewLicense("wanna", 1, "21.5", false, false)
	licenseKey := license.GetLicenseKey()
	fmt.Printf("licenseKey : %s \n", licenseKey)
	encryptBytes := EncryptBytes(0x787, []byte(licenseKey))
	fmt.Printf("encryptBytes : %s \n", encryptBytes)
	encode := VariantBase64Encode([]byte(encryptBytes))
	fmt.Printf("encode: %s \n", encode)
}

func TestGenZip(t *testing.T) {
	license := entity.NewLicense("wanna", 1, "21.5", false, false)
	licenseKey := license.GetLicenseKey()
	encryptBytes := EncryptBytes(0x787, []byte(licenseKey))
	encode := VariantBase64Encode([]byte(encryptBytes))
	fmt.Printf("encode: %s \n", encode)

	create, _ := os.Create("Custom.mxtpro")
	zipWriter := zip.NewWriter(create)
	dosDate, dosTime := timeToMsDosTime(time.Now())
	header := zip.FileHeader{
		Name:               "Pro.key",
		Method:             zip.Store,
		ReaderVersion:      10,
		Flags:              2048,
		CRC32:              crc32.ChecksumIEEE([]byte(encode)),
		UncompressedSize64: uint64(len(encode)),
		CompressedSize64:   uint64(len(encode)),
		ModifiedTime:       dosTime,
		ModifiedDate:       dosDate,
	}
	f, _ := zipWriter.CreateRaw(&header)
	f.Write([]byte(encode))

	zipWriter.Close()
	create.Close()
}

// 实在是不会了,通过 java 版本 比对 字节数组来实现吧
func TestByteCode(t *testing.T) {
	// java byteCode

	//byteArr := []byte{80, 75, 3, 4, 10, 0, 0, 8, 0, 0, 217, 144, 187, 84, 91, 202, 211, 118, 36, 0, 0, 0, 36, 0, 0, 0, 7, 0, 0, 0, 80, 114, 111, 46, 107, 101, 121}
	goBtye := []byte{80, 75, 3, 4, 10, 0, 0, 8, 0, 0, 0, 0, 0, 0, 91, 202, 211, 118, 36, 0, 0, 0, 36, 0, 0, 0, 7, 0, 0, 0, 80, 114, 111, 46, 107, 101, 121}

	byteArr1 := []byte{80, 75, 1, 2, 10, 0, 10, 0, 0, 8, 0, 0, 217, 144, 187, 84, 91, 202, 211, 118, 36, 0, 0, 0, 36, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 80, 114, 111, 46, 107, 101, 121, 80, 75, 5, 6, 0, 0, 0, 0, 1, 0, 1, 0, 53, 0, 0, 0, 73, 0, 0, 0, 0, 0}

	encode := "2s2PpYiJpQje513a5tme5tXf+13a4tGerh3a"

	var zipBuffer *bytes.Buffer = new(bytes.Buffer)
	zipBuffer.Write(goBtye)
	zipBuffer.Write([]byte(encode))

	zipBuffer.Write(byteArr1)

	ioutil.WriteFile("Custom.mxtpro", zipBuffer.Bytes(), 0644)
}

func TestVersion(t *testing.T) {
	expression, _ := regexp.Compile("^\\d+\\.\\d+$")
	expression.MatchString("2.0")
	fmt.Println("222")
}
