package providers

import (
	"fmt"
	"os"
)

func DecodeAndDecrypt(encoded string, provider string) ([]byte, error) {

	switch provider {

	case "static":
		keyPhrase := os.Getenv("KEYPHRASE")
		if keyPhrase == "" {
			return nil, fmt.Errorf("keyphrase not found")
		}
		staticDecodeAndDecrypt(encoded, keyPhrase)

	}

	return nil, nil

}

func EncryptAndEncode(value string, provider string) (string, error) {

	switch provider {
	case "static":
		keyPhrase := os.Getenv("KEYPHRASE")
		if keyPhrase == "" {
			return "", fmt.Errorf("keyphrase not found")
		}
		staticDecodeAndDecrypt(value, keyPhrase)

	}
	return "", nil

}
