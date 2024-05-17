package ec

import (
	"math/big"
)

type Point3D struct {
	X, Y, Z *big.Int
}

// https://www.nayuki.io/page/elliptic-curve-point-addition-in-projective-coordinates
func (p Point3D) Double(c *Curve) Point {
	if p.IsInf() {
		return p
	}

	if p.Y.Cmp(zero) == 0 {
		return inf3d()
	}

	T1 := big.NewInt(1)
	T2 := big.NewInt(1)

	T := new(big.Int).Add(T1.Mul(T1, p.X).Mul(T1, p.X).Mul(T1, three), T2.Mul(T2, c.A).Mul(T2, p.Z).Mul(T2, p.Z))
	T.Mod(T, c.Mod)

	U := big.NewInt(1)
	U.Mul(U, two).Mul(U, p.Y).Mul(U, p.Z).Mod(U, c.Mod)

	V := big.NewInt(1)
	V.Mul(V, two).Mul(V, U).Mul(V, p.X).Mul(V, p.Y).Mod(V, c.Mod)

	W1 := new(big.Int).Exp(T, two, c.Mod)
	W2 := new(big.Int).Mul(two, V)
	W2.Mod(W2, c.Mod)

	W := new(big.Int).Sub(W1, W2)
	W.Mod(W, c.Mod)
	W.Add(W, c.Mod)
	W.Mod(W, c.Mod)

	X := new(big.Int).Mul(U, W)
	X.Mod(X, c.Mod)

	Y1 := new(big.Int).Sub(V, W)
	Y1.Mod(Y1, c.Mod)
	Y1.Add(Y1, c.Mod)
	Y1.Mod(Y1, c.Mod)
	Y1.Mul(Y1, T).Mod(Y1, c.Mod)

	Y2 := new(big.Int).Mul(U, p.Y)
	Y2.Mul(Y2, Y2).Mul(two, Y2).Mod(Y2, c.Mod)

	Y := new(big.Int).Sub(Y1, Y2)
	Y.Mod(Y, c.Mod)
	Y.Add(Y, c.Mod)
	Y.Mod(Y, c.Mod)

	Z := new(big.Int).Exp(U, three, c.Mod)

	return Point3D{
		X: X,
		Y: Y,
		Z: Z,
	}
}

func (p1 Point3D) Add(point Point, c *Curve) Point {
	p2 := point.(Point3D)

	if p1.IsInf() {
		return p2.Copy()
	}

	U1 := new(big.Int).Mul(p2.Y, p1.Z)
	U1.Mod(U1, c.Mod)

	U2 := new(big.Int).Mul(p2.Z, p1.Y)
	U2.Mod(U2, c.Mod)

	V1 := new(big.Int).Mul(p2.X, p1.Z)
	V1.Mod(V1, c.Mod)

	V2 := new(big.Int).Mul(p2.Z, p1.X)
	V2.Mod(V2, c.Mod)

	if V1.Cmp(V2) == 0 {
		if U1.Cmp(U2) != 0 {
			return inf3d()
		}

		return p1.Double(c)
	}

	U := new(big.Int).Sub(U1, U2)
	U.Add(U, c.Mod)
	U.Mod(U, c.Mod)

	V := new(big.Int).Sub(V1, V2)
	V.Add(V, c.Mod)
	V.Mod(V, c.Mod)

	W := new(big.Int).Mul(p1.Z, p2.Z)
	W.Mod(W, c.Mod)

	A1 := new(big.Int).Exp(U, two, c.Mod)
	A1.Mul(A1, W)
	A1.Mod(A1, c.Mod)
	A1.Sub(A1, new(big.Int).Exp(V, three, c.Mod))
	A1.Mod(A1, c.Mod)

	A2 := new(big.Int).Exp(V, two, c.Mod)
	A2.Mul(A2, V2)
	A2.Mod(A2, c.Mod)
	A2.Mul(A2, two)
	A2.Mod(A2, c.Mod)

	A := new(big.Int).Sub(A1, A2)
	A.Mod(A, c.Mod)
	A.Add(A, c.Mod)
	A.Mod(A, c.Mod)

	X := new(big.Int).Mul(V, A)
	X.Mod(X, c.Mod)

	Y1 := new(big.Int).Exp(V, two, c.Mod)
	Y1.Mul(Y1, V2)
	Y1.Mod(Y1, c.Mod)
	Y1.Sub(Y1, A)
	Y1.Mod(Y1, c.Mod)
	Y1.Add(Y1, c.Mod)
	Y1.Mod(Y1, c.Mod)
	Y1.Mul(Y1, U)
	Y1.Mod(Y1, c.Mod)

	Y2 := new(big.Int).Exp(V, three, c.Mod)
	Y2.Mul(Y2, U2)
	Y2.Mod(Y2, c.Mod)

	Y := new(big.Int).Sub(Y1, Y2)
	Y.Add(Y, c.Mod)
	Y.Mod(Y, c.Mod)

	Z := new(big.Int).Exp(V, three, c.Mod)
	Z.Mul(Z, W)
	Z.Mod(Z, c.Mod)

	return Point3D{
		X: X,
		Y: Y,
		Z: Z,
	}
}

func (p Point3D) Copy() Point {
	return Point3D{X: new(big.Int).Set(p.X), Y: new(big.Int).Set(p.Y), Z: new(big.Int).Set(p.Z)}
}

func (p Point3D) IsInf() bool {
	return p.X.Cmp(zero) == 0 && p.Y.Cmp(one) == 0 && p.Z.Cmp(zero) == 0
}

func (p Point3D) IsEqual(point Point) bool {
	p1 := point.(Point3D) // Bad pattern but it's ok for lab

	return p.X.Cmp(p1.X) == 0 && p.Y.Cmp(p1.Y) == 0 && p.Z.Cmp(p1.Z) == 0
}

func (p Point3D) To2D(c *Curve) Point2D {
	mul := new(big.Int).ModInverse(p.Z, c.Mod)

	X := new(big.Int).Mul(p.X, mul)
	X.Mod(X, c.Mod)

	Y := new(big.Int).Mul(p.Y, mul)
	Y.Mod(Y, c.Mod)

	return Point2D{
		X: X,
		Y: Y,
	}
}

func inf3d() Point3D {
	return Point3D{X: zero, Y: one, Z: zero}
}
