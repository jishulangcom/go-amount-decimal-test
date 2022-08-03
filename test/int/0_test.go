package inttest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"testing"
)

func Test0(t *testing.T) {
	fmt.Println("\n//----------------0/0=?---------------------------")
	str1, err1 := amountdecimal.New(0).Div(0).ToString(10) // 被除数为0
	fmt.Println(str1, err1) // 0 <nil>

	fmt.Println("\n//----------------1/0=?---------------------------")
	str2, err2 := amountdecimal.New(1).Div(0).ToString(10) // 除数为0报错
	fmt.Println(str2, err2) //  divisor cannot be 0
}
