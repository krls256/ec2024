package tests

import (
	"ec"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type CurveParams struct {
	P   *big.Int
	A   *big.Int
	B   *big.Int
	X   *big.Int
	Y   *big.Int
	Ord *big.Int
}

func (p *CurveParams) Curve() *ec.Curve {
	return &ec.Curve{
		A:   p.A,
		B:   p.B,
		Mod: p.P,
	}
}

func (p *CurveParams) Point2D() ec.Point2D {
	return ec.Point2D{
		X: p.X,
		Y: p.Y,
	}
}

func (p *CurveParams) Point3D() ec.Point3D {
	return ec.Point3D{
		X: p.X,
		Y: p.Y,
		Z: big.NewInt(1),
	}
}

func Must[T any](t T, ok bool) T {
	if !ok {
		panic(t)
	}

	return t
}

// https://neuromancer.sk/std/nist/P-192
var TestCases = []CurveParams{
	{ // 192
		P:   Must(new(big.Int).SetString("fffffffffffffffffffffffffffffffeffffffffffffffff", 16)),
		A:   Must(new(big.Int).SetString("fffffffffffffffffffffffffffffffefffffffffffffffc", 16)),
		B:   Must(new(big.Int).SetString("64210519e59c80e70fa7e9ab72243049feb8deecc146b9b1", 16)),
		X:   Must(new(big.Int).SetString("188da80eb03090f67cbf20eb43a18800f4ff0afd82ff1012", 16)),
		Y:   Must(new(big.Int).SetString("07192b95ffc8da78631011ed6b24cdd573f977a11e794811", 16)),
		Ord: Must(new(big.Int).SetString("ffffffffffffffffffffffff99def836146bc9b1b4d22831", 16)),
	},

	{ // 224
		P:   Must(new(big.Int).SetString("ffffffffffffffffffffffffffffffff000000000000000000000001", 16)),
		A:   Must(new(big.Int).SetString("fffffffffffffffffffffffffffffffefffffffffffffffffffffffe", 16)),
		B:   Must(new(big.Int).SetString("b4050a850c04b3abf54132565044b0b7d7bfd8ba270b39432355ffb4", 16)),
		X:   Must(new(big.Int).SetString("b70e0cbd6bb4bf7f321390b94a03c1d356c21122343280d6115c1d21", 16)),
		Y:   Must(new(big.Int).SetString("bd376388b5f723fb4c22dfe6cd4375a05a07476444d5819985007e34", 16)),
		Ord: Must(new(big.Int).SetString("ffffffffffffffffffffffffffff16a2e0b8f03e13dd29455c5c2a3d", 16)),
	},

	{ // 256
		P:   Must(new(big.Int).SetString("ffffffff00000001000000000000000000000000ffffffffffffffffffffffff", 16)),
		A:   Must(new(big.Int).SetString("ffffffff00000001000000000000000000000000fffffffffffffffffffffffc", 16)),
		B:   Must(new(big.Int).SetString("5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", 16)),
		X:   Must(new(big.Int).SetString("6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", 16)),
		Y:   Must(new(big.Int).SetString("4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5", 16)),
		Ord: Must(new(big.Int).SetString("ffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551", 16)),
	},

	{ // 384
		P:   Must(new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffeffffffff0000000000000000ffffffff", 16)),
		A:   Must(new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffeffffffff0000000000000000fffffffc", 16)),
		B:   Must(new(big.Int).SetString("b3312fa7e23ee7e4988e056be3f82d19181d9c6efe8141120314088f5013875ac656398d8a2ed19d2a85c8edd3ec2aef", 16)),
		X:   Must(new(big.Int).SetString("aa87ca22be8b05378eb1c71ef320ad746e1d3b628ba79b9859f741e082542a385502f25dbf55296c3a545e3872760ab7", 16)),
		Y:   Must(new(big.Int).SetString("3617de4a96262c6f5d9e98bf9292dc29f8f41dbd289a147ce9da3113b5f0b8c00a60b1ce1d7e819d7a431d7c90ea0e5f", 16)),
		Ord: Must(new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffc7634d81f4372ddf581a0db248b0a77aecec196accc52973", 16)),
	},

	{ // 521
		P:   Must(new(big.Int).SetString("01ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)),
		A:   Must(new(big.Int).SetString("01fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc", 16)),
		B:   Must(new(big.Int).SetString("0051953eb9618e1c9a1f929a21a0b68540eea2da725b99b315f3b8b489918ef109e156193951ec7e937b1652c0bd3bb1bf073573df883d2c34f1ef451fd46b503f00", 16)),
		X:   Must(new(big.Int).SetString("00c6858e06b70404e9cd9e3ecb662395b4429c648139053fb521f828af606b4d3dbaa14b5e77efe75928fe1dc127a2ffa8de3348b3c1856a429bf97e7e31c2e5bd66", 16)),
		Y:   Must(new(big.Int).SetString("011839296a789a3bc0045c8a5fb42c7d1bd998f54449579b446817afbd17273e662c97ee72995ef42640c550b9013fad0761353c7086a272c24088be94769fd16650", 16)),
		Ord: Must(new(big.Int).SetString("01fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa51868783bf2f966b7fcc0148f709a5d03bb5c9b8899c47aebb6fb71e91386409", 16)),
	},
}

func Test_2DCorrectness(t *testing.T) {
	for _, tc := range TestCases {
		c := tc.Curve()
		p := tc.Point2D()

		pNext := c.ScalarMulPoint2D(tc.Ord, p)

		require.True(t, pNext.IsInf())
	}
}

func Test_3DCorrectness(t *testing.T) {
	for _, tc := range TestCases {
		c := tc.Curve()
		p := tc.Point3D()

		pNext := c.ScalarMulPoint2D(tc.Ord, p)

		require.True(t, pNext.IsInf())
	}
}

func Benchmark_3D_192(b *testing.B) {
	tc := TestCases[0]

	c := tc.Curve()
	p := tc.Point3D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_2D_192(b *testing.B) {
	tc := TestCases[0]

	c := tc.Curve()
	p := tc.Point2D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_3D_224(b *testing.B) {
	tc := TestCases[1]

	c := tc.Curve()
	p := tc.Point3D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_2D_224(b *testing.B) {
	tc := TestCases[1]

	c := tc.Curve()
	p := tc.Point2D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_3D_256(b *testing.B) {
	tc := TestCases[2]

	c := tc.Curve()
	p := tc.Point3D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_2D_256(b *testing.B) {
	tc := TestCases[2]

	c := tc.Curve()
	p := tc.Point2D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_3D_384(b *testing.B) {
	tc := TestCases[3]

	c := tc.Curve()
	p := tc.Point3D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_2D_384(b *testing.B) {
	tc := TestCases[3]

	c := tc.Curve()
	p := tc.Point2D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_3D_521(b *testing.B) {
	tc := TestCases[4]

	c := tc.Curve()
	p := tc.Point3D()

	c.ScalarMulPoint2D(tc.Ord, p)
}

func Benchmark_2D_521(b *testing.B) {
	tc := TestCases[4]

	c := tc.Curve()
	p := tc.Point2D()

	c.ScalarMulPoint2D(tc.Ord, p)
}
