package floattest

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal-test/test"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestFloat64TimeConsuming(t *testing.T) {
	list := test.RandFloat64List(float64Min, float64Max, float64Decimal)

	Float64TimeConsuming1(list)
	fmt.Println("")
	Float64TimeConsuming2(list)
	fmt.Println("")
	Float64TimeConsuming3(list)
}

// amountdecimal耗时
func Float64TimeConsuming1(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat64(float64Min, float64Max)
		Float64CalculationAmountDecimal(test.Add, v, v2)
		Float64CalculationAmountDecimal(test.Sub, v, v2)
		Float64CalculationAmountDecimal(test.Mul, v, v2)
		Float64CalculationAmountDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// amountdecimal2耗时
func Float64TimeConsuming2(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat64(float64Min, float64Max)
		Float64CalculationAmountDecimal2(test.Add, v, v2)
		Float64CalculationAmountDecimal2(test.Sub, v, v2)
		Float64CalculationAmountDecimal2(test.Mul, v, v2)
		Float64CalculationAmountDecimal2(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func Float64TimeConsuming3(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat64(float64Min, float64Max)
		Float64CalculationDecimal(test.Add, v, v2)
		Float64CalculationDecimal(test.Sub, v, v2)
		Float64CalculationDecimal(test.Mul, v, v2)
		Float64CalculationDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}
