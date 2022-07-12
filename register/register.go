package register

import (
	"archive/zip"
	"hash/crc32"
	"os"
	"strings"
	"time"
)

const (
	FileName  = "Custom.mxtpro"
	EntryName = "Pro.key"
)

func Register(licenseKey, path string) {
	encryptBytes := EncryptBytes(0x787, []byte(licenseKey))
	encode := VariantBase64Encode([]byte(encryptBytes))

	var fileName = ""
	if path == "" {
		fileName = FileName
	} else if strings.HasSuffix(path, "/") {
		fileName = path + FileName
	} else {
		fileName = path + "/" + FileName
	}

	file, _ := os.Create(fileName)
	zipWriter := zip.NewWriter(file)
	dosDate, dosTime := timeToMsDosTime(time.Now())
	header := zip.FileHeader{
		Name:               EntryName,
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
	file.Close()
}

func timeToMsDosTime(t time.Time) (fDate uint16, fTime uint16) {
	fDate = uint16(t.Day() + int(t.Month())<<5 + (t.Year()-1980)<<9)
	fTime = uint16(t.Second()/2 + t.Minute()<<5 + t.Hour()<<11)
	return
}
