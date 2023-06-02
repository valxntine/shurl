package service

import (
	"crypto/rand"
	"encoding/base64"
)

func createShortUrl() (string, error) {
	// here we use 6 bytes of entropy as a simple solution, could easily replace with some different logic to determine length of the short url
	var data [6]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}
