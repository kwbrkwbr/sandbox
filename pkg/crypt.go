package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func CryptCFB(v, k []byte) ([]byte, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(c, iv)
	r := make([]byte, len(v))
	cfb.XORKeyStream(r, v)
	return r, nil
}

func DecryptCFB(v, k []byte) ([]byte, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBDecrypter(c, iv)
	r := make([]byte, len(v))
	cfb.XORKeyStream(r, v)
	return r, nil
}

func CryptCBC(v, k []byte) ([]byte, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	cbcIV := make([]byte, aes.BlockSize)
	if _, err := rand.Read(cbcIV); err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(c, iv)
	pad := cbcPad(v)
	pad = append(cbcIV, pad...) // 先頭にIVをつける
	b := make([]byte, len(pad))
	cbc.CryptBlocks(b, pad)
	return b, nil
}

func DecryptCBC(v, k []byte) ([]byte, error) {
	c, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	// IVと複合対象を分ける
	cbcIV := v[:aes.BlockSize]
	v = v[aes.BlockSize:]
	cbc := cipher.NewCBCDecrypter(c, cbcIV)

	r := make([]byte, len(v))
	cbc.CryptBlocks(r, v)
	padded := cbcUnPad(r)
	return padded, nil
}

// 末尾に足りないサイズを足りないサイズ数の数字でpadding
func cbcPad(b []byte) []byte {
	padSize := aes.BlockSize - (len(b) % aes.BlockSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(b, pad...)
}

// 末尾のサイズ数の数字をそのサイズ分削る
func cbcUnPad(b []byte) []byte {
	padSize := int(b[len(b)-1])
	r := b[:len(b)-padSize]
	return r
}
