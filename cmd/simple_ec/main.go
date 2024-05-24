package main

import (
	"ec"
	"fmt"
	"math/big"
)

func main() {
	p, ok := new(big.Int).SetString("97", 10)
	if !ok {
		panic("cant parse")
	}

	a, ok := new(big.Int).SetString("2", 10)
	if !ok {
		panic("cant parse")
	}

	b, ok := new(big.Int).SetString("3", 10)
	if !ok {
		panic("cant parse")
	}

	x, ok := new(big.Int).SetString("17", 10)
	if !ok {
		panic("cant parse")
	}

	y, ok := new(big.Int).SetString("10", 10)
	if !ok {
		panic("cant parse")
	}

	//ord, ok := new(big.Int).SetString("51", 10)
	//if !ok {
	//	panic("cant parse")
	//}

	c := ec.Curve{
		A:   a,
		B:   b,
		Mod: p,
	}

	p1 := ec.Point3D{
		X: x,
		Y: y,
		Z: big.NewInt(1),
	}

	p2 := ec.Point2D{
		X: x,
		Y: y,
	}

	for i := range 10 {
		ord := big.NewInt(int64(i + 1))

		r1 := c.ScalarMulPoint2D(ord, p1)
		r1T := r1.(ec.Point3D)
		r12D := r1T.To2D(&c)

		r2 := c.ScalarMulPoint2D(ord, p2)

		if !r2.IsEqual(r12D) {
			fmt.Println(ord, r1, r12D, r2)
		}
	}
}
