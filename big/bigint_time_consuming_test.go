package bigtest

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal/test"
	"math/big"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestBigIntTimeConsuming(t *testing.T) {
	list := test.RandBigIntList(bigBigMin, bigBigMax, bigBigNub)

	BigIntTimeConsuming1(list)
	fmt.Println("")
	BigIntTimeConsuming2(list)
	fmt.Println("")
	BigIntTimeConsuming3(list)
}

// amountdecimal耗时
func BigIntTimeConsuming1(list []*big.Int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandBigInt(bigBigMin, bigBigMax)
		BigIntCalculationAmountDecimal(test.Add, v, v2)
		BigIntCalculationAmountDecimal(test.Sub, v, v2)
		BigIntCalculationAmountDecimal(test.Mul, v, v2)
		BigIntCalculationAmountDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// amountdecimal2耗时
func BigIntTimeConsuming2(list []*big.Int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandBigInt(bigBigMin, bigBigMax)
		BigIntCalculationAmountDecimal2(test.Add, v, v2)
		BigIntCalculationAmountDecimal2(test.Sub, v, v2)
		BigIntCalculationAmountDecimal2(test.Mul, v, v2)
		BigIntCalculationAmountDecimal2(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func BigIntTimeConsuming3(list []*big.Int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandBigInt(bigBigMin, bigBigMax)
		BigIntCalculationDecimal(test.Add, v, v2)
		BigIntCalculationDecimal(test.Sub, v, v2)
		BigIntCalculationDecimal(test.Mul, v, v2)
		BigIntCalculationDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}
