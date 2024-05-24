package ec

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/sha256"
)

const KeySize = 32

func RandomKey() [KeySize]byte {
	var key [KeySize]byte

	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}

	return key
}

func KeyToFineSize(key []byte) [KeySize]byte {
	h := sha256.New()
	h.Write(key)

	return [KeySize]byte(h.Sum(nil)[:KeySize])
}

func Enc(message []byte, key [KeySize]byte) (cypher []byte) {
	return enc(message, key)
}

func Dec(cypher []byte, key [KeySize]byte) (message []byte) {
	return dec(cypher, key)
}

func Wrap(message []byte, key []byte) (cypher []byte) {
	return Enc(message, KeyToFineSize(key))
}

func Unwrap(message []byte, key []byte) (cypher []byte) {
	return Dec(message, KeyToFineSize(key))
}

func enc(message []byte, key [KeySize]byte) []byte {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	ds := c.BlockSize()

	cypher := []byte{}
	cypherBuf := make([]byte, ds)

	addLen := 0
	if len(message)%ds != 0 {
		addLen = 1
	}

	for i := 0; i < len(message)/ds+addLen; i++ {
		messageBuf := make([]byte, ds)

		copy(messageBuf, message[i*ds:])

		c.Encrypt(cypherBuf, messageBuf)

		cypher = append(cypher, cypherBuf...)
	}

	return cypher
}

func dec(cypher []byte, key [KeySize]byte) (message []byte) {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}

	ds := c.BlockSize()

	message = []byte{}
	messageBuf := make([]byte, ds)

	for i := 0; i < len(cypher)/ds; i++ {
		cypherBuf := make([]byte, ds)

		copy(cypherBuf, cypher[i*ds:])

		c.Decrypt(messageBuf, cypherBuf)

		message = append(message, messageBuf...)
	}

	return message
}

func Hash(message []byte) []byte {
	h := sha256.New()
	h.Write(message)

	return h.Sum(nil)
}
