package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func GetMd5(key string) (string, error) {
	if key == "" {
		return "", errors.New("param is nil")
	}
	m5 := md5.New()
	m5.Write([]byte(key))
	sliceRet := m5.Sum(nil)
	strRet := hex.EncodeToString(sliceRet)
	return strRet, nil
}
