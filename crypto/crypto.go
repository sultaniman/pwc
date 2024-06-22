package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// TODO(thanks): Encryption was adopted from https://gist.github.com/enyachoke/5c60f5eebed693d9b4bacddcad693b47

const (
	NumIterations    = 10000
	SaltLengthBytes  = 8
	PassphraseLength = 16
	KeyLength        = 32

	// IVKeyLength initialization vector in bytes
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// See Section 8.2 & Section 8.2.1, 12 bytes = 96 bits will use
	// deterministic methods to derive a key for passphrase.
	IVKeyLength = 12
)

type Message struct {
	Encrypted string
	Plaintext string
}

func NewMessage(plaintext string, encrypted string) *Message {
	return &Message{
		Encrypted: encrypted,
		Plaintext: plaintext,
	}
}

func (m *Message) DeriveKey(passphrase string, salt []byte) ([]byte, []byte, error) {
	// http://www.ietf.org/rfc/rfc2898.txt 4.1 Salt
	if salt == nil {
		salt = make([]byte, SaltLengthBytes)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key := pbkdf2.Key([]byte(passphrase), salt, NumIterations, KeyLength, sha256.New)
	return key, salt, nil
}

func (m *Message) Encrypt(passphrase string) (string, error) {
	key, salt, err := m.DeriveKey(passphrase, nil)
	iv := make([]byte, IVKeyLength)

	_, err = rand.Read(iv)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesCGM, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	encryptedBytes := aesCGM.Seal(nil, iv, []byte(m.Plaintext), nil)
	m.Encrypted = fmt.Sprintf("%x:%x:%x", salt, iv, encryptedBytes)
	return m.Encrypted, nil
}

func (m *Message) Decrypt(passphrase string) (string, error) {
	parts := strings.Split(m.Encrypted, ":")
	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	iv, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(parts[2])
	if err != nil {
		return "", err
	}

	key, _, err := m.DeriveKey(passphrase, salt)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesCGM, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	decryptedBytes, err := aesCGM.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	m.Plaintext = string(decryptedBytes)
	return m.Plaintext, nil
}

// RandomPassphrase godoc
// Generate random passphrase to encrypt generated data
// Returns first 16 characters from random passphrase.
func (m *Message) RandomPassphrase() (string, error) {
	passphraseBytes := make([]byte, KeyLength)
	_, err := rand.Read(passphraseBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(passphraseBytes)[:PassphraseLength], nil
}
