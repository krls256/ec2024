package ec

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncDec(t *testing.T) {
	key := RandomKey()

	message := []byte("lalala")

	require.Equal(t, message, dec(enc(message, key), key)[:len(message)])
}
