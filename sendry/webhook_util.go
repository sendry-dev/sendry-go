package sendry

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// VerifyWebhookSignature verifies the HMAC-SHA256 signature of a webhook payload.
//
// The signature header format is: "t=<unix_timestamp>,v1=<hex_digest>"
//
// This function parses the header, recomputes the HMAC over
// "<timestamp>.<payload>" using the provided secret, and performs a
// constant-time comparison to prevent timing attacks.
//
// Returns true if the signature is valid, false otherwise.
//
// Example:
//
//	valid := sendr.VerifyWebhookSignature(
//	    string(body),
//	    r.Header.Get("Sendr-Signature"),
//	    "whsec_your_signing_secret",
//	)
//	if !valid {
//	    http.Error(w, "invalid signature", http.StatusUnauthorized)
//	    return
//	}
func VerifyWebhookSignature(payload, signature, secret string) bool {
	if payload == "" || signature == "" || secret == "" {
		return false
	}

	// Parse "t=<timestamp>,v1=<hex>"
	var timestamp, v1 string
	for _, part := range strings.Split(signature, ",") {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "t=") {
			timestamp = strings.TrimPrefix(part, "t=")
		} else if strings.HasPrefix(part, "v1=") {
			v1 = strings.TrimPrefix(part, "v1=")
		}
	}

	if timestamp == "" || v1 == "" {
		return false
	}

	// Decode the expected hex digest
	expected, err := hex.DecodeString(v1)
	if err != nil {
		return false
	}

	// Recompute: HMAC-SHA256(secret, timestamp + "." + payload)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write([]byte(payload))
	computed := mac.Sum(nil)

	// Constant-time comparison
	return hmac.Equal(computed, expected)
}
