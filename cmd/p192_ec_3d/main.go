package main

import (
	"ec"
	"fmt"
	"math/big"
)

func main() {
	p, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffeffffffffffffffff", 16)
	if !ok {
		panic("cant parse")
	}

	a, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffefffffffffffffffc", 16)
	if !ok {
		panic("cant parse")
	}

	b, ok := new(big.Int).SetString("64210519e59c80e70fa7e9ab72243049feb8deecc146b9b1", 16)
	if !ok {
		panic("cant parse")
	}

	x, ok := new(big.Int).SetString("188da80eb03090f67cbf20eb43a18800f4ff0afd82ff1012", 16)
	if !ok {
		panic("cant parse")
	}

	y, ok := new(big.Int).SetString("07192b95ffc8da78631011ed6b24cdd573f977a11e794811", 16)
	if !ok {
		panic("cant parse")
	}

	ord, ok := new(big.Int).SetString("ffffffffffffffffffffffff99def836146bc9b1b4d22831", 16)
	if !ok {
		panic("cant parse")
	}

	c := ec.Curve{
		A:   a,
		B:   b,
		Mod: p,
	}

	point := ec.Point3D{
		X: x,
		Y: y,
		Z: big.NewInt(1),
	}

	fmt.Println(c.SmartScalarMulPoint(ord, point))
}
