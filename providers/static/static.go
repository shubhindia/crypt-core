package static

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"

	"github.com/shubhindia/crypt-core/providers/utils"
)

func DecodeAndDecrypt(encoded string, keyPhrase string) []byte {
	ciphered, _ := base64.StdEncoding.DecodeString(encoded)
	hashedPhrase := utils.MdHashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		log.Fatalln(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Fatalln(err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return originalText

}

func EncryptAndEncode(value string, keyPhrase string) (string, error) {

	aesBlock, err := aes.NewCipher([]byte(utils.MdHashing(keyPhrase)))
	if err != nil {
		return "", err
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, []byte(value), nil)

	encoded := base64.StdEncoding.EncodeToString(cipheredText)

	return encoded, nil
}