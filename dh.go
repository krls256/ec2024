package ec

import (
	"crypto/rand"
	"math/big"
)

func NewDiffieHellman(c *Curve, p Point, ord *big.Int) DiffieHellman {
	dh := DiffieHellman{
		c:   c,
		p:   p,
		ord: ord,
	}

	dh.generatePrivateKey()

	return dh
}

type DiffieHellman struct {
	c   *Curve
	p   Point
	ord *big.Int

	privateKey *big.Int
}

func (dh *DiffieHellman) generateOrdScalar() *big.Int {
	scalar, err := rand.Int(rand.Reader, new(big.Int).Sub(dh.ord, two))
	if err != nil {
		panic(err) // Come on its lab don't want to create error system
	}

	return scalar.Add(scalar, two)
}

func (dh *DiffieHellman) generateModScalar() *big.Int {
	scalar, err := rand.Int(rand.Reader, new(big.Int).Sub(dh.c.Mod, two))
	if err != nil {
		panic(err) // Come on its lab don't want to create error system
	}

	return scalar.Add(scalar, two)
}

func (dh *DiffieHellman) generatePrivateKey() {
	dh.privateKey = dh.generateOrdScalar()
}

func (dh *DiffieHellman) PublicKey() Point {
	return dh.c.ScalarMulPoint2D(dh.privateKey, dh.p)
}

func (dh *DiffieHellman) SharedSecret(pk Point) Point {
	return dh.c.ScalarMulPoint2D(dh.privateKey, pk)
}

func (dh *DiffieHellman) MessageEnc(pk Point, rawText []byte) EncryptedMessage {
	ss := dh.SharedSecret(pk)

	Sx, _ := ss.Coords2D(dh.c)

	symKey := RandomKey()

	cypherText := Enc(rawText, symKey)
	wrappedKey := Wrap(symKey[:], Sx.Bytes())

	return EncryptedMessage{
		CypherText: cypherText,
		WrappedKey: wrappedKey,
		PublicKey:  dh.PublicKey(),
	}
}

func (dh *DiffieHellman) MessageDec(m EncryptedMessage) []byte {
	ss := dh.SharedSecret(m.PublicKey)

	Sx, _ := ss.Coords2D(dh.c)

	key := Unwrap(m.WrappedKey, Sx.Bytes())

	return Dec(m.CypherText, [32]byte(key))
}

type EncryptedMessage struct {
	CypherText []byte
	WrappedKey []byte
	PublicKey  Point
}

func (dh *DiffieHellman) hash(message []byte) *big.Int {
	hB := Hash(message)

	h := new(big.Int)

	return h.SetBytes(hB).Mod(h, dh.c.Mod)
}

func (dh *DiffieHellman) Sign(message []byte) (r, s *big.Int) {
	k := dh.generateModScalar()
	kInf := new(big.Int).ModInverse(k, dh.ord)

	r, _ = dh.c.ScalarMulPoint2D(k, dh.p).Coords2D(dh.c)
	if r.Cmp(big.NewInt(0)) == 0 {
		panic("wrong x")
	}

	s = new(big.Int).Mul(dh.privateKey, r)
	s.Add(s, dh.hash(message)).Mul(s, kInf).Mod(s, dh.ord)

	return r, s
}

func (dh *DiffieHellman) Verify(message []byte, r, s *big.Int, pk Point) bool {
	sInv := new(big.Int).ModInverse(s, dh.ord)

	u1 := new(big.Int).Mul(sInv, dh.hash(message))
	u1.Mod(u1, dh.ord)

	u2 := new(big.Int).Mul(sInv, r)
	u2.Mod(u2, dh.ord)

	u1P := dh.c.ScalarMulPoint2D(u1, dh.p)
	u2Q := dh.c.ScalarMulPoint2D(u2, pk)

	x, _ := dh.c.Add(u1P, u2Q).Coords2D(dh.c)

	x.Mod(x, dh.c.Mod)

	return x.Cmp(r) == 0
}
