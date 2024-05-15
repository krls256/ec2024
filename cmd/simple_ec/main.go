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

	ord, ok := new(big.Int).SetString("3", 10)
	if !ok {
		panic("cant parse")
	}

	c := ec.Curve{
		A:   a,
		B:   b,
		Mod: p,
	}

	point := ec.Point{
		X: x,
		Y: y,
	}

	for i := range 100 {
		ord.SetInt64(int64(i + 1))

		if !c.StupidScalarMulPoint(ord, point).IsEqual(c.SmartScalarMulPoint(ord, point)) {
			fmt.Println(i)
		}
	}

}
