package floattest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal" // go:lazyinit
	"testing"
)

func Test01(t *testing.T) {
	var forNub = 10
	var fee float64 = 0.1
	var amountF float64 = 1.00

	fmt.Println("\n//---------------------【add】-----------------------------------------")
	ad := amountdecimal.New(0)
	for i := 1; i <= forNub; i++ {
		ad = ad.Add(fee)
		amount, err := ad.ToString(1)
		fmt.Println("buy：", i, " balance：", amount, err)
	}

	fmt.Println("\n//---------------------【sub】-----------------------------------------")
	ad2 := amountdecimal.New(amountF)
	for i := 1; i <= forNub; i++ {
		ad2 = ad2.Sub(fee)
		amount2, err2 := ad2.ToString(1)
		fmt.Println("buy：", i, " balance：", amount2, err2)
	}

	fmt.Println("\n//---------------------【mul】------------------------------------------")
	ad3 := amountdecimal.New(fee)
	for i := 1; i <= forNub; i++ {
		ad3 = ad3.Mul(fee)
		amount3, err3 := ad3.ToString(10)
		fmt.Println("buy：", i, " balance：", amount3, err3)
	}

	fmt.Println("\n//---------------------【div】-------------------------------------------")
	ad4 := amountdecimal.New(fee)
	for i := 1; i <= forNub; i++ {
		ad4 = ad4.Div(fee)
		amount4, err4 := ad4.ToString(10)
		fmt.Println("buy：", i, " balance：", amount4, err4)
	}
}
