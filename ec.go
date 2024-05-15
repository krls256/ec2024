package ec

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Int
}

func (p Point) IsEqual(p1 Point) bool {
	return p.X.Cmp(p1.X) == 0 && p.Y.Cmp(p1.Y) == 0
}

func (p Point) Minus(mod *big.Int) Point {
	nP := p.Copy()

	nP.Y.Mul(nP.Y, minusOne)
	nP.Y.Mod(nP.Y, mod)

	return nP
}

func (p Point) Copy() Point {
	return Point{X: new(big.Int).Set(p.X), Y: new(big.Int).Set(p.Y)}
}

func (p Point) IsInf() bool {
	return p.X.Cmp(minusOne) == 0 && p.Y.Cmp(minusOne) == 0
}

func (p Point) String() string {
	if p.IsInf() {
		return "O Inf"
	}

	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

type Curve struct {
	A, B, Mod *big.Int
}

var two = big.NewInt(2)
var three = big.NewInt(3)
var minusOne = big.NewInt(-1)

func (c *Curve) SmartScalarMulPoint(k *big.Int, p Point) Point {
	Q := p

	for i := k.BitLen() - 2; i >= 0; i-- {
		Q = c.Add(Q, Q)
		
		if k.Bit(i) == 1 {
			Q = c.Add(Q, p)
		}
	}

	return Q
}

func (c *Curve) StupidScalarMulPoint(k *big.Int, p Point) Point {
	Q := p

	for i := int64(1); i < k.Int64(); i++ {
		Q = c.Add(Q, p)
	}

	return Q
}

func (c *Curve) Add(p1, p2 Point) Point {
	if p1.IsInf() {
		return p2.Copy()
	}

	if p1.X.Cmp(p2.X) == 0 {
		if p1.Y.Cmp(p2.Y) == 0 {
			return double(p1, c.A, c.Mod)
		}

		return inf()
	}

	return add(p1, p2, c.Mod)
}

func add(p1, p2 Point, mod *big.Int) Point {
	p := Point{X: new(big.Int), Y: new(big.Int)}

	p.X = otherAngleCoef(p1, p2, mod)
	p.X.Exp(p.X, two, mod)
	p.X.Sub(p.X, p1.X)
	p.X.Mod(p.X, mod)
	p.X.Sub(p.X, p2.X)
	p.X.Mod(p.X, mod)

	p.Y = otherAngleCoef(p1, p2, mod)
	p.Y.Mod(p.Y, mod)

	tmpY := new(big.Int).Sub(p.X, p1.X)
	tmpY = tmpY.Mod(tmpY, mod)
	p.Y.Mul(p.Y, tmpY)
	p.Y.Mul(p.Y, minusOne)
	p.Y.Sub(p.Y, p1.Y)

	p.Y.Mod(p.Y, mod)

	return p
}

func double(p1 Point, a, mod *big.Int) Point {
	p := Point{X: new(big.Int), Y: new(big.Int)}

	tmp := selfAngleCoef(p1, a, mod)

	p.X.Exp(tmp, two, mod)

	p.X.Sub(p.X, p1.X)
	p.X.Sub(p.X, p1.X)
	p.X.Mod(p.X, mod)

	tmpY := new(big.Int).Sub(p.X, p1.X)
	p.Y.Mul(tmp, tmpY)
	p.Y.Mod(p.Y, mod)
	p.Y.Mul(p.Y, minusOne)

	p.Y.Sub(p.Y, p1.Y)
	p.Y.Mod(p.Y, mod)

	return p
}

func inf() Point {
	return Point{X: big.NewInt(-1), Y: big.NewInt(-1)}
}

func selfAngleCoef(p1 Point, a, mod *big.Int) *big.Int {
	tmpXMul, tmpXDiv := new(big.Int), new(big.Int)

	tmpXMul.Exp(p1.X, two, mod)
	tmpXMul.Mul(tmpXMul, three)
	tmpXMul.Add(tmpXMul, a)

	tmpXDiv.Mul(p1.Y, two)
	tmpXDiv.ModInverse(tmpXDiv, mod)

	coef := new(big.Int).Mul(tmpXMul, tmpXDiv)
	coef.Mod(coef, mod)

	return coef
}

func otherAngleCoef(p1, p2 Point, mod *big.Int) *big.Int {
	tmpXMul, tmpXDiv := new(big.Int), new(big.Int)
	tmpXMul.Sub(p1.Y, p2.Y)
	tmpXMul.Mod(tmpXMul, mod)

	tmpXDiv.Sub(p1.X, p2.X)
	tmpXDiv.Mod(tmpXDiv, mod)

	tmpXDiv.ModInverse(tmpXDiv, mod)

	coef := new(big.Int).Mul(tmpXMul, tmpXDiv)
	coef.Mod(coef, mod)

	return coef
}
