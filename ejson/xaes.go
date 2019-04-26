package ejson

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"utils/slog"
)

const aesKey = "Qi=+-ho!&(wniyeK" //16 bytes  AS-128

func Encrypt(origData []byte) ([]byte, error) {

	key := []byte(aesKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		slog.Debug("Server Encrypt NewCipher error")
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)

	//slog.Debug("Server Encrypt length %v", len(crypted))
	return crypted, nil
}

func Decrypt(crypted []byte) ([]byte, error) {

	key := []byte(aesKey)

	//slog.Debug("Server Decrypt length %v", len(crypted))
	block, err := aes.NewCipher(key)
	if err != nil {
		slog.Debug("Server Decrypt NewCipher error")
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)

	origData, err = pKCS5UnPadding(origData)
	if err != nil {
		slog.Debug("Server Decrypt error")
		return nil, err
	}

	return origData, nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("pKCS5UnPadding error!")
	}

	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)], nil
}

func SNXDecode(src []byte) []byte {
	var i int
	dest := make([]byte, len(src))
	copy(dest, src)

	dest[8] = '-'
	dest[13] = '-'
	dest[18] = '-'
	dest[23] = '-'

	for i = 0; i < 3; i++ {
		dest[5+i] = src[2+i]
		dest[2+i] = src[5+i]
		dest[15+i] = src[26+i]
		dest[26+i] = src[15+i]
	}

	return dest
}
