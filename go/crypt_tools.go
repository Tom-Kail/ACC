package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CryptReq struct {
	Data        string `form:"data"`
	EncryptType string `form:"encryptType"`
	Key         string `form:"key"`
}

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func md5Encrypt(content string, bitLen int) ([]byte, error) {
	if bitLen != 16 && bitLen != 32 {
		return nil, errors.New("位长只能是16位或32位")
	}

	h := md5.New()
	h.Write([]byte(content)) // 需要加密的字符串为 sharejs.com
	longCipher := hex.EncodeToString(h.Sum(nil))
	if bitLen == 32 {
		return []byte(strings.ToUpper(longCipher)), nil
	} else {
		shortCipher := longCipher[8:24]
		return []byte(strings.ToUpper(shortCipher)), nil
	}
}

// 3DES加密
func tripleDesEncrypt(origDataStr, keyStr string) ([]byte, error) {
	origData := []byte(origDataStr)
	key := []byte(keyStr)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func tripleDesDecrypt(cryptedStr, keyStr string) ([]byte, error) {
	crypted := []byte(cryptedStr)
	key := []byte(keyStr)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func desEncrypt(origDataStr, keyStr string) ([]byte, error) {
	origData := []byte(origDataStr)
	key := []byte(keyStr)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//des解密

func desDecrypt(cryptedStr, keyStr string) ([]byte, error) {
	crypted := []byte(cryptedStr)
	key := []byte(keyStr)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	//origData := make([]byte, len(crypted))
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func getKey(strKey string) []byte {
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func aesEncrypt(strMesg string, strKey string) ([]byte, error) {
	if len(strKey) < 16 {
		return nil, errors.New("The key length cannot be less than 16")
	}
	key := getKey(strKey)
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

//解密字符串
func aesDecrypt(origDataStr string, strKey string) (rst []byte, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src := []byte(origDataStr)
	key := getKey(strKey)
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}

func base64Encode(originDataStr string) ([]byte, error) {
	tmpCipher := base64.StdEncoding.EncodeToString([]byte(originDataStr))
	return []byte(tmpCipher), nil
}

func base64Decode(originDataStr string) ([]byte, error) {
	tmpCipher, cryErr := base64.StdEncoding.DecodeString(originDataStr)
	if cryErr != nil {
		return nil, errors.New("Base64字符串格式不正确")
	}
	return tmpCipher, cryErr
}

func sha1Encrypt(origData string) ([]byte, error) {
	t := sha1.New()
	_, err := io.WriteString(t, origData)
	if err != nil {
		return nil, err
	} else {
		return []byte(fmt.Sprintf("%x", t.Sum(nil))), nil
	}
}

func EncryptData(req CryptReq) ([]byte, error) {
	cipher := []byte{}
	err := errors.New("no error")
	err = nil
	switch req.EncryptType {
	case "md5_16":
		cipher, err = md5Encrypt(req.Data, 16)
	case "md5_32":
		cipher, err = md5Encrypt(req.Data, 32)
	case "sha1":
		cipher, err = sha1Encrypt(req.Data)
	case "aes":
		MD5Key, genKeyErr := md5Encrypt(req.Key, 16)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		tmpCipher, cryErr := aesEncrypt(req.Data, string(MD5Key))
		if cryErr != nil {
			return nil, cryErr
		}
		cipherStr := base64.StdEncoding.EncodeToString(tmpCipher)
		cipher = []byte(cipherStr)

	case "des":
		MD5Key, genKeyErr := md5Encrypt(req.Key, 16)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		tmpCipher, cryErr := desEncrypt(req.Data, string(MD5Key[0:8]))
		if cryErr != nil {
			return nil, cryErr
		}
		cipherStr := base64.StdEncoding.EncodeToString(tmpCipher)
		cipher = []byte(cipherStr)
	case "3des":
		MD5Key, genKeyErr := md5Encrypt(req.Key, 32)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		tmpCipher, cryErr := tripleDesEncrypt(req.Data, string(MD5Key[0:24]))
		if cryErr != nil {
			return nil, cryErr
		}
		cipherStr := base64.StdEncoding.EncodeToString(tmpCipher)
		cipher = []byte(cipherStr)
	case "base64":
		cipher, err = base64Encode(req.Data)
	default:
		return nil, errors.New("encrypt type is not supported")
	}
	if err != nil {
		return nil, err
	} else {
		return cipher, nil
	}
}

func DecryptData(req CryptReq) ([]byte, error) {
	clearText := []byte{}
	err := errors.New("no error")
	err = nil
	switch req.EncryptType {
	case "aes":

		tmpCipher, cryErr := base64.StdEncoding.DecodeString(req.Data)
		if cryErr != nil {
			return nil, cryErr
		}
		MD5Key, genKeyErr := md5Encrypt(req.Key, 16)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		clearText, err = aesDecrypt(string(tmpCipher), string(MD5Key))
	case "des":
		tmpCipher, cryErr := base64.StdEncoding.DecodeString(req.Data)
		if cryErr != nil {
			return nil, cryErr
		}
		MD5Key, genKeyErr := md5Encrypt(req.Key, 16)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		clearText, err = desDecrypt(string(tmpCipher), string(MD5Key[0:8]))
	case "3des":
		tmpCipher, cryErr := base64.StdEncoding.DecodeString(req.Data)
		if cryErr != nil {
			return nil, cryErr
		}
		MD5Key, genKeyErr := md5Encrypt(req.Key, 32)
		if genKeyErr != nil {
			return nil, genKeyErr
		}
		clearText, err = tripleDesDecrypt(string(tmpCipher), string(MD5Key[0:24]))
	case "base64":
		clearText, err = base64Decode(req.Data)
	default:
		return nil, errors.New("decrypt type is not supported")
	}
	if err != nil {
		return nil, err
	} else {
		return clearText, nil

	}
}
