package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
)

// GetTransactionID generates a unique transaction ID for a document based on its encrypted content
func GetTransactionID(encryptedContent []byte) (string, error) {
    // Hash the encrypted content
    hash := sha256.Sum256(encryptedContent)

    // Convert the hash to a hexadecimal string
    txID := hex.EncodeToString(hash[:])

    if txID == "" {
        return "", errors.New("failed to get transaction ID")
    }

    return txID, nil
}

func GenerateBlockchainTransactionID() int {
    return rand.Int()
} 