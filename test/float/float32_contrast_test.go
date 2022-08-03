package floattest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/jishulangcom/go-amount-decimal-test/test"
	"github.com/shopspring/decimal"
	"testing"
)

var (
	float32Min     float32     = 0
	float32Max     float32     = 99.9999999999999999999999999999999
	float32Nub     int         = 10000
	float32Decimal int         = amountdecimal.COIN_ETH_USDT_DECIMAL
	float32Coin    interface{} = amountdecimal.COIN_ETH_USDT
)

func TestFloat32(t *testing.T) {
	list := test.RandFloat32List(float32Min, float32Max, float32Nub)
	for _, v := range list {
		v2 := test.RandFloat32(float32Min, float32Max)
		Float32Add(v, v2)
		Float32Sub(v, v2)
		Float32Mul(v, v2)
		Float32Div(v, v2)
	}
}

func Float32Add(v float32, v2 float32) {
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)

	res, err := Float32CalculationAmountDecimal(test.Add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := Float32CalculationAmountDecimal2(test.Add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := Float32CalculationDecimal(test.Add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float32Sub(v float32, v2 float32) {
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)

	res, err := Float32CalculationAmountDecimal(test.Sub, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := Float32CalculationAmountDecimal2(test.Sub, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := Float32CalculationDecimal(test.Sub, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float32Mul(v float32, v2 float32) {
	fmt.Printf("\n\r ----------%v * %v ------------------------\n", v, v2)

	res, err := Float32CalculationAmountDecimal(test.Mul, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := Float32CalculationAmountDecimal2(test.Mul, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := Float32CalculationDecimal(test.Mul, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float32Div(v float32, v2 float32) {
	fmt.Printf("\n\r ----------%v / %v ------------------------\n", v, v2)

	res, err := Float32CalculationAmountDecimal(test.Div, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		fmt.Println("calculation_amountdecimal err:", err)
		panic(err)
	}

	res2, err2 := Float32CalculationAmountDecimal2(test.Div, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		fmt.Println("calculation_amountdecimal2 err:", err2)
		panic(err2)
	}

	res3, err3 := Float32CalculationDecimal(test.Div, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float32CalculationAmountDecimal(f uint8, v float32, v2 float32) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.New(v).Add(v2).ToString(float32Decimal)
	case test.Sub:
		return amountdecimal.New(v).Sub(v2).ToString(float32Decimal)
	case test.Mul:
		return amountdecimal.New(v).Mul(v2).ToString(float32Decimal)
	case test.Div:
		return amountdecimal.New(v).Div(v2).ToString(float32Decimal)
	}

	return "", nil
}

func Float32CalculationAmountDecimal2(f uint8, v float32, v2 float32) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.NewFloat32(v).AddFloat32(v2).ToString(float32Decimal)
	case test.Sub:
		return amountdecimal.NewFloat32(v).SubFloat32(v2).ToString(float32Decimal)
	case test.Mul:
		return amountdecimal.NewFloat32(v).MulFloat32(v2).ToString(float32Decimal)
	case test.Div:
		return amountdecimal.NewFloat32(v).DivFloat32(v2).ToString(float32Decimal)
	}

	return "", nil
}

func Float32CalculationDecimal(f uint8, v float32, v2 float32) (string, error) {
	var res string

	v2d := decimal.NewFromFloat(float64(v2))

	// 计算方法所对应的操作
	switch f {
	case test.Add:
		res = decimal.NewFromFloat(float64(v)).Add(v2d).StringFixed(int32(float32Decimal))
	case test.Sub:
		res = decimal.NewFromFloat(float64(v)).Sub(v2d).StringFixed(int32(float32Decimal))
	case test.Mul:
		res = decimal.NewFromFloat(float64(v)).Mul(v2d).StringFixed(int32(float32Decimal))
	case test.Div:
		res = decimal.NewFromFloat(float64(v)).Div(v2d).StringFixed(int32(float32Decimal))
	}

	return res, nil
}
