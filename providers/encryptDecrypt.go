package providers

import (
	"fmt"
	"os"
)

func DecodeAndDecrypt(encoded string, provider string) (string, error) {

	switch provider {

	case "static":
		keyPhrase := os.Getenv("KEYPHRASE")
		if keyPhrase == "" {
			return "", fmt.Errorf("keyphrase not found")
		}

		decoded, err := staticDecodeAndDecrypt(encoded, keyPhrase)
		if err != nil {
			return "", err
		}
		return decoded, nil

	}

	return "", nil

}

func EncryptAndEncode(value string, provider string) (string, error) {

	switch provider {
	case "static":
		keyPhrase := os.Getenv("KEYPHRASE")
		if keyPhrase == "" {
			return "", fmt.Errorf("keyphrase not found")
		}
		return staticEncryptAndEncode(value, keyPhrase)

	}
	return "", nil

}
