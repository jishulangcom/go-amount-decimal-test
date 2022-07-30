package floattest

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal/test"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestFloat32TimeConsuming(t *testing.T) {
	list := test.RandFloat32List(float32Min, float32Max, float32Decimal)

	Float32TimeConsuming1(list)
	fmt.Println("")
	Float32TimeConsuming2(list)
	fmt.Println("")
	Float32TimeConsuming3(list)
}

// amountdecimal耗时
func Float32TimeConsuming1(list []float32) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat32(float32Min, float32Max)
		Float32CalculationAmountDecimal(test.Add, v, v2)
		Float32CalculationAmountDecimal(test.Sub, v, v2)
		Float32CalculationAmountDecimal(test.Mul, v, v2)
		Float32CalculationAmountDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// amountdecimal2耗时
func Float32TimeConsuming2(list []float32) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat32(float32Min, float32Max)
		Float32CalculationAmountDecimal2(test.Add, v, v2)
		Float32CalculationAmountDecimal2(test.Sub, v, v2)
		Float32CalculationAmountDecimal2(test.Mul, v, v2)
		Float32CalculationAmountDecimal2(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func Float32TimeConsuming3(list []float32) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := test.RandFloat32(float32Min, float32Max)
		Float32CalculationDecimal(test.Add, v, v2)
		Float32CalculationDecimal(test.Sub, v, v2)
		Float32CalculationDecimal(test.Mul, v, v2)
		Float32CalculationDecimal(test.Div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}
