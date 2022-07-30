package amountdecimaltest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"testing"
)

func Test333(t *testing.T) {
	//--------------------【1/3 + 1/3 + 1/3 = 1】-------------------------------
	fmt.Println("\n--------------------【1/3 + 1/3 + 1/3 = 1】-------------------------------")
  	a, _ := amountdecimal.New("1").Div("3").ToBigRat() // 1/3
  	b, _ := amountdecimal.New("1").Div("3").ToBigRat() // 1/3
  	c, _ := amountdecimal.New("1").Div("3").ToBigRat() // 1/3
  	fmt.Println("abc:", a, b, c) // abc: 1/3 1/3 1/3

  	strA1, errA1 := amountdecimal.New(0).Add(a).Add(b).Add(c).ToString(nil) // 1/3 + 1/3 + 1/3
  	fmt.Println("sum:", strA1, errA1) // sum: 1 <nil>

	strA2, errA2 := amountdecimal.New(0).Add(a).Add(b).Add(c).ToString(3) // 1/3 + 1/3 + 1/3
	fmt.Println("sum:", strA2, errA2) // sum: 1.000 <nil>
	fmt.Println()


  	//--------------------【0.333 + 0.333 = ?】-------------------------------
	fmt.Println("\n--------------------【0.333 + 0.333 = ?】-------------------------------")

  	val := 0.333
	//val := math.MaxFloat64
	strB1, errB1 := amountdecimal.New(val).Add(val).Add(val).ToString(nil)
	fmt.Println("B1:", strB1, errB1) // B1: 0.999 <nil>

	strB2, errB2 := amountdecimal.New(val).Add(val).Add(val).ToString(2)
	fmt.Println("B2:", strB2, errB2) // B2: 1.00 <nil>

	strB3, errB3 := amountdecimal.New(val).Add(val).Add(val).ToString(3)
	fmt.Println("B3:", strB3, errB3) // B3: 0.999 <nil>


	//--------------------【0.333 * 0.333 = ?】-------------------------------
	fmt.Println("\n--------------------【0.333 * 0.333 = ?】-------------------------------")

	valC := 0.333
	//val := math.MaxFloat64
	strC1, errC1 := amountdecimal.New(valC).Mul(valC).Mul(valC).ToString(nil)
	fmt.Println("C1:", strC1, errC1) // C1: 0.037 <nil>

	strC2, errC2 := amountdecimal.New(valC).Mul(valC).Mul(valC).ToString(2)
	fmt.Println("C2:", strC2, errC2) // C2: 0.04 <nil>

	strC3, errC3 := amountdecimal.New(valC).Mul(valC).Mul(valC).ToString(3)
	fmt.Println("C3:", strC3, errC3) // C3: 0.037 <nil>


	//--------------------【1 / 3 = ?】-------------------------------
	fmt.Println("\n--------------------【1 / 3 = ?】-------------------------------")
	valD := 3
	strD1, errD1 := amountdecimal.New(1).Div(valD).ToString(nil)
	fmt.Println("D1:", strD1, errD1) // C1: 0.037 <nil>

	strD2, errD2 := amountdecimal.New(1).Div(valD).ToString(2)
	fmt.Println("D2:", strD2, errD2) // C2: 0.04 <nil>

	strD3, errD3 := amountdecimal.New(1).Div(valD).ToString(3)
	fmt.Println("D3:", strD3, errD3) // C3: 0.037 <nil>

}
