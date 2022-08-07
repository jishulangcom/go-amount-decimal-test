package test

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func RandBigIntList(min, max int64, n int) []*big.Int {
	list := make([]*big.Int, n)
	for i := range list {
		list[i] = RandBigInt(min, max)
	}
	return list
}

func RandBigInt(min, max int64) *big.Int {
	i := min + int64(rand.Int())*(max-min)
	fmt.Println("dddd:", i)
	return new(big.Int).SetInt64(i)
}

func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandFloat64List(min, max float64, n int) []float64 {
	list := make([]float64, n)
	for i := range list {
		list[i] = RandFloat64(min, max)
	}
	return list
}

func RandFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandFloat32List(min, max float32, n int) []float32 {
	list := make([]float32, n)
	for i := range list {
		list[i] = RandFloat32(min, max)
	}
	return list
}

func RandIntList(min, max int, n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = RandInt(min, max)
	}
	return list
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min == 0 && max == 0 {
		return 0
	}
	return rand.Intn(max-min) + min
}

func RandUint(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min == 0 && max == 0 {
		return 0
	}
	return rand.Intn(max-min) + min
}
