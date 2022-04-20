package provider

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func hashForState(value string) string {
	if value == "" {
		return ""
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(value)))
	return hex.EncodeToString(hash[:])
}
