package payhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateSignature generating hash string from params that validate webhook signature
func GenerateSignature(apiKey string, webhookID string, webhookEvent string, jsonBody string) string {
	signatureString := fmt.Sprintf("id=%s\nevent=%s\ndata=%s", webhookID, webhookEvent, jsonBody)
	data := []byte(signatureString)

	hash := hmac.New(sha256.New, []byte(apiKey))
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}
