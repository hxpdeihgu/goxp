package goxp


import (
	"encoding/hex"
	"crypto/md5"
	"crypto/sha1"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
)

func _md5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func _sha1(s string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(s))
	cipherStr := sha1Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func _sha256(s string) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(s))
	cipherStr := sha256Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func _sha512(s string) string {
	sha512Ctx := sha512.New()
	sha512Ctx.Write([]byte(s))
	cipherStr := sha512Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}


type AesEncrypt struct {
	Key string
}

func (this *AesEncrypt) getKey() []byte {
	keyLen := len(this.Key)
	if keyLen < 16 {
		panic("aes key error!")
	}
	arrKey := []byte(this.Key)
	if keyLen >= 32 {
		return arrKey[:32]
	}
	if keyLen >= 24 {
		return arrKey[:24]
	}
	return arrKey[:16]
}

//加密字符串
func (this *AesEncrypt) Encrypt(strMesg string) ([]byte, error) {
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	
	if err != nil {
		return nil, err
	}
	
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}
//加密字符串
func (this *AesEncrypt) EncryptBytes(strMesg []byte) ([]byte, error) {
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	
	if err != nil {
		return nil, err
	}
	
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, strMesg)
	return encrypted, nil
}

//aes解密字符串
func (this *AesEncrypt) Decrypt(src []byte) (string,error) {
	defer func() {
		if e := recover(); e != nil {
			err := e.(error)
			Error(err)
		}
	}()
	
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	var aesBlockDecrypter cipher.Block
	
	decrypted := make([]byte, len(src))
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
//aes解密字符串
func (this *AesEncrypt) DecryptBytes(src []byte) ([]byte,error) {
	defer func() {
		if e := recover(); e != nil {
			err := e.(error)
			Error(err)
		}
	}()
	
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	var aesBlockDecrypter cipher.Block
	
	decrypted := make([]byte, len(src))
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	
	if err != nil {
		return nil, err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}


