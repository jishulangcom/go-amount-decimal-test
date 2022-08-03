package inttest

import (
	"errors"
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/jishulangcom/go-amount-decimal-test/test"
	"github.com/shopspring/decimal"
	"testing"
)

var (
	intMin     int         = -200 // math.MinInt64
	intMax     int         = 10  // math.MaxInt64
	intNub     int         = 10000
	intDecimal int         = amountdecimal.COIN_ETH_USDT_DECIMAL
	intCoin    interface{} = amountdecimal.COIN_ETH_USDT
)

func TestInt(t *testing.T) {
	list := test.RandIntList(intMin, intMax, intNub)
	for _, v := range list {
		v2 := test.RandInt(intMin, intMax)
		IntAdd(v, v2)
		IntSub(v, v2)
		IntMul(v, v2)
		IntDiv(v, v2)
	}
}

func IntAdd(v int, v2 int) {
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)

	res, err := IntCalculationAmountDecimal(test.Add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := IntCalculationAmountDecimal2(test.Add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := IntCalculationDecimal(test.Add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func IntSub(v int, v2 int) {
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)

	res, err := IntCalculationAmountDecimal(test.Add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := IntCalculationAmountDecimal2(test.Add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := IntCalculationDecimal(test.Add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func IntMul(v int, v2 int) {
	fmt.Printf("\n\r ----------%v * %v ------------------------\n", v, v2)

	res, err := IntCalculationAmountDecimal(test.Mul, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := IntCalculationAmountDecimal2(test.Mul, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := IntCalculationDecimal(test.Mul, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func IntDiv(v int, v2 int) {
	fmt.Printf("\n\r ----------%v / %v ------------------------\n", v, v2)

	res, err := IntCalculationAmountDecimal(test.Div, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		fmt.Println("calculation_amountdecimal err:", err)
		panic(err)
	}

	res2, err2 := IntCalculationAmountDecimal2(test.Div, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		fmt.Println("calculation_amountdecimal2 err:", err2)
		panic(err2)
	}

	res3, err3 := IntCalculationDecimal(test.Div, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func IntCalculationAmountDecimal(f uint8, v int, v2 int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.New(v).Add(v2).ToString(intCoin)
	case test.Sub:
		return amountdecimal.New(v).Sub(v2).ToString(intCoin)
	case test.Mul:
		return amountdecimal.New(v).Mul(v2).ToString(intCoin)
	case test.Div:
		return amountdecimal.New(v).Div(v2).ToString(intCoin)
	}

	return "", nil
}

func IntCalculationAmountDecimal2(f uint8, v int, v2 int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.NewInt(v).AddInt(v2).ToString(intCoin)
	case test.Sub:
		return amountdecimal.NewInt(v).SubInt(v2).ToString(intCoin)
	case test.Mul:
		return amountdecimal.NewInt(v).MulInt(v2).ToString(intCoin)
	case test.Div:
		return amountdecimal.NewInt(v).DivInt(v2).ToString(intCoin)
	}

	return "", nil
}

func IntCalculationDecimal(f uint8, v int, v2 int) (string, error) {
	var res string

	v2d := decimal.NewFromInt(int64(v2))

	// 计算方法所对应的操作
	switch f {
	case test.Add:
		res = decimal.NewFromInt(int64(v)).Add(v2d).StringFixed(int32(intDecimal))
	case test.Sub:
		res = decimal.NewFromInt(int64(v)).Sub(v2d).StringFixed(int32(intDecimal))
	case test.Mul:
		res = decimal.NewFromInt(int64(v)).Mul(v2d).StringFixed(int32(intDecimal))
	case test.Div:
		if v2 == 0 {
			return "", errors.New("被除数不能为0")
		}
		res = decimal.NewFromInt(int64(v)).Div(v2d).StringFixed(int32(intDecimal))
	}

	return res, nil
}
