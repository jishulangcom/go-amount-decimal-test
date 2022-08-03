package bigtest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/jishulangcom/go-amount-decimal/test"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

var (
	bigBigMin     int64       = 0
	bigBigMax     int64       = 99
	bigBigNub     int         = 1
	bigBigDecimal int         = amountdecimal.COIN_ETH_USDT_DECIMAL
	bigBigCoin    interface{} = amountdecimal.COIN_ETH_USDT
)

func TestBigInt(t *testing.T) {
	list := test.RandBigIntList(bigBigMin, bigBigMax, bigBigNub)
	for _, v := range list {
		v2 := test.RandBigInt(bigBigMin, bigBigMax)
		BigIntAdd(v, v2)
		BigIntSub(v, v2)
		BigIntMul(v, v2)
		BigIntDiv(v, v2)
	}
}

func BigIntAdd(v *big.Int, v2 *big.Int) {
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)

	res, err := BigIntCalculationAmountDecimal(test.Add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := BigIntCalculationAmountDecimal2(test.Add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := BigIntCalculationDecimal(test.Add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func BigIntSub(v *big.Int, v2 *big.Int) {
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)

	res, err := BigIntCalculationAmountDecimal(test.Add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := BigIntCalculationAmountDecimal2(test.Add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := BigIntCalculationDecimal(test.Add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func BigIntMul(v *big.Int, v2 *big.Int) {
	fmt.Printf("\n\r ----------%v * %v ------------------------\n", v, v2)

	res, err := BigIntCalculationAmountDecimal(test.Mul, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := BigIntCalculationAmountDecimal2(test.Mul, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := BigIntCalculationDecimal(test.Mul, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func BigIntDiv(v *big.Int, v2 *big.Int) {
	fmt.Printf("\n\r ----------%v / %v ------------------------\n", v, v2)

	res, err := BigIntCalculationAmountDecimal(test.Div, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		fmt.Println("calculation_amountdecimal err:", err)
		panic(err)
	}

	res2, err2 := BigIntCalculationAmountDecimal2(test.Div, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		fmt.Println("calculation_amountdecimal2 err:", err2)
		panic(err2)
	}

	res3, err3 := BigIntCalculationDecimal(test.Div, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func BigIntCalculationAmountDecimal(f uint8, v *big.Int, v2 *big.Int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.New(v).Add(v2).ToString(bigBigCoin)
	case test.Sub:
		return amountdecimal.New(v).Sub(v2).ToString(bigBigCoin)
	case test.Mul:
		return amountdecimal.New(v).Mul(v2).ToString(bigBigCoin)
	case test.Div:
		return amountdecimal.New(v).Div(v2).ToString(bigBigCoin)
	}

	return "", nil
}

func BigIntCalculationAmountDecimal2(f uint8, v *big.Int, v2 *big.Int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case test.Add:
		return amountdecimal.NewBigInt(v).AddBigInt(v2).ToString(bigBigCoin)
	case test.Sub:
		return amountdecimal.NewBigInt(v).SubBigInt(v2).ToString(bigBigCoin)
	case test.Mul:
		return amountdecimal.NewBigInt(v).MulBigInt(v2).ToString(bigBigCoin)
	case test.Div:
		return amountdecimal.NewBigInt(v).DivBigInt(v2).ToString(bigBigCoin)
	}

	return "", nil
}

func BigIntCalculationDecimal(f uint8, v *big.Int, v2 *big.Int) (string, error) {
	var res string

	v2d := decimal.NewFromBigInt(v2, int32(bigBigDecimal))

	// 计算方法所对应的操作
	switch f {
	case test.Add:
		res = decimal.NewFromBigInt(v, 0).Add(v2d).StringFixed(int32(bigBigDecimal))
	case test.Sub:
		res = decimal.NewFromBigInt(v, 0).Sub(v2d).StringFixed(int32(bigBigDecimal))
	case test.Mul:
		res = decimal.NewFromBigInt(v, 0).Mul(v2d).StringFixed(int32(bigBigDecimal))
	case test.Div:
		res = decimal.NewFromBigInt(v, 0).Div(v2d).StringFixed(int32(bigBigDecimal))
	}

	return res, nil
}
