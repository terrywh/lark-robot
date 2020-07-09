package toolkit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
)

var crc64_table *crc64.Table

func init() {
	crc64_table = crc64.MakeTable(crc64.ISO)
}
func doHash(methods []string, data string) (string, error) {
	if len(methods) < 1 {
		return "", errors.New("未知 HASH 方法")
	}

	var h hash.Hash
	switch methods[1] {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "crc32":
		h = crc32.NewIEEE()
	case "crc64":
		h = crc64.New(crc64_table)
	default:
		return "", errors.New("未知 HASH 方法")
	}
	io.WriteString(h, data)
	return fmt.Sprintf("%0x", h.Sum(nil)), nil
}
