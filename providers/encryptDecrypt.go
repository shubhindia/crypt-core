package providers

import (
	"context"
	"fmt"
	"os"

	"github.com/shubhindia/crypt-core/providers/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	case "k8s":
		k8sClient, err := utils.GetKubeClient()
		if err != nil {
			return "", fmt.Errorf("failed to get kubeclient %v", err)
		}
		// Define the namespace and secret name to retrieve
		namespace := "default"
		secretName := "default"

		// Retrieve the secret from the Kubernetes cluster
		secret, err := k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to get the secret %v", err)
		}

		// Access the secret data"
		keyPhrase := string(secret.Data["token"])

		// decode the data
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

	case "k8s":
		k8sClient, err := utils.GetKubeClient()
		if err != nil {
			return "", fmt.Errorf("failed to get kubeclient %v", err)
		}
		// Define the namespace and secret name to retrieve
		namespace := "default"
		secretName := "default"

		// Retrieve the secret from the Kubernetes cluster
		secret, err := k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to get the secret %v", err)
		}

		// Access the secret data"
		keyPhrase := string(secret.Data["token"])

		// encrypt the data
		return staticEncryptAndEncode(value, keyPhrase)

	}
	return "", nil

}
