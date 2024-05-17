package ec

import (
	"fmt"
	"math/big"
)

type Point2D struct {
	X, Y *big.Int
}

func (p Point2D) Double(c *Curve) Point {
	return double2d(p, c.A, c.Mod)
}

func (p Point2D) Add(point Point, c *Curve) Point {
	p1 := point.(Point2D) // Bad pattern but it's ok for lab

	if p.X.Cmp(p1.X) == 0 {
		return inf2d()
	}

	return add2d(p, p1, c.Mod)
}

func (p Point2D) Copy() Point {
	return p.copy()
}

func (p Point2D) IsEqual(point Point) bool {
	p1 := point.(Point2D) // Bad pattern but it's ok for lab

	return p.isEqual(p1)
}

func (p Point2D) isEqual(p1 Point2D) bool {
	return p.X.Cmp(p1.X) == 0 && p.Y.Cmp(p1.Y) == 0
}

func (p Point2D) minus(mod *big.Int) Point2D {
	nP := p.copy()

	nP.Y.Mul(nP.Y, minusOne)
	nP.Y.Mod(nP.Y, mod)

	return nP
}

func (p Point2D) copy() Point2D {
	return Point2D{X: new(big.Int).Set(p.X), Y: new(big.Int).Set(p.Y)}
}

func (p Point2D) IsInf() bool {
	return p.X.Cmp(minusOne) == 0 && p.Y.Cmp(minusOne) == 0
}

func (p Point2D) String() string {
	if p.IsInf() {
		return "O Inf"
	}

	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func add2d(p1, p2 Point2D, mod *big.Int) Point2D {
	p := Point2D{X: new(big.Int), Y: new(big.Int)}

	p.X = otherAngleCoef2d(p1, p2, mod)
	p.X.Exp(p.X, two, mod)
	p.X.Sub(p.X, p1.X)
	p.X.Mod(p.X, mod)
	p.X.Sub(p.X, p2.X)
	p.X.Mod(p.X, mod)

	p.Y = otherAngleCoef2d(p1, p2, mod)
	p.Y.Mod(p.Y, mod)

	tmpY := new(big.Int).Sub(p.X, p1.X)
	tmpY = tmpY.Mod(tmpY, mod)
	p.Y.Mul(p.Y, tmpY)
	p.Y.Mul(p.Y, minusOne)
	p.Y.Sub(p.Y, p1.Y)

	p.Y.Mod(p.Y, mod)

	return p
}

func double2d(p1 Point2D, a, mod *big.Int) Point2D {
	p := Point2D{X: new(big.Int), Y: new(big.Int)}

	tmp := selfAngleCoef2d(p1, a, mod)

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

func inf2d() Point2D {
	return Point2D{X: big.NewInt(-1), Y: big.NewInt(-1)}
}

func selfAngleCoef2d(p1 Point2D, a, mod *big.Int) *big.Int {
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

func otherAngleCoef2d(p1, p2 Point2D, mod *big.Int) *big.Int {
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
