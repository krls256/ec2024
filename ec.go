package ec

import (
	"math/big"
)

type Point interface {
	Double(c *Curve) Point
	Add(p Point, c *Curve) Point

	Copy() Point

	IsInf() bool
	IsEqual(Point) bool
}

type Curve struct {
	A, B, Mod *big.Int
}

var two = big.NewInt(2)
var three = big.NewInt(3)
var minusOne = big.NewInt(-1)
var zero = big.NewInt(0)
var one = big.NewInt(1)
var eight = big.NewInt(8)
var four = big.NewInt(4)

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

	if p1.IsEqual(p2) {
		return p1.Double(c)
	}

	return p1.Add(p2, c)
}
