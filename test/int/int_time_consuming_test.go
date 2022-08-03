package inttest

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal-test/test"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestIntTimeConsuming(t *testing.T) {
	list := test.RandIntList(intMin, intMax, intNub)

	IntTimeConsuming1(list)
	fmt.Println("")
	IntTimeConsuming2(list)
	fmt.Println("")
	IntTimeConsuming3(list)
}

// amountdecimal耗时
func IntTimeConsuming1(list []int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandInt(intMin, intMax)
		IntCalculationAmountDecimal(test.Add, v, v2)
		IntCalculationAmountDecimal(test.Sub, v, v2)
		IntCalculationAmountDecimal(test.Mul, v, v2)
		IntCalculationAmountDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// amountdecimal2耗时
func IntTimeConsuming2(list []int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandInt(intMin, intMax)
		IntCalculationAmountDecimal2(test.Add, v, v2)
		IntCalculationAmountDecimal2(test.Sub, v, v2)
		IntCalculationAmountDecimal2(test.Mul, v, v2)
		IntCalculationAmountDecimal2(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func IntTimeConsuming3(list []int) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandInt(intMin, intMin)
		IntCalculationDecimal(test.Add, v, v2)
		IntCalculationDecimal(test.Sub, v, v2)
		IntCalculationDecimal(test.Mul, v, v2)
		IntCalculationDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}
