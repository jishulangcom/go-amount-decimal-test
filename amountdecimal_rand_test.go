package amountdecimaltest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"testing"
)

func Test(t *testing.T) {
	//var i float32 = 1
	//i := json.Number("14")
	//i := &struct {}{}
	//var i interface{}
	//amountBigRat := new(big.Rat)
	//amountBigRat.SetString("1.222")
	//
	//amountBigInt := new(big.Int)
	//amountBigInt.SetString("111", 0)
	//
	////i := &kk{}
	//
	//
	//amount := amountdecimal.New(amountBigInt, 8)
	//result, err := amount.Div("2.1").String()
	//fmt.Println(result, err)
	//
	//result2, err2 := amount.Div("2.1").JsonNumber()
	//fmt.Println(result2, err2)
	//
	//amountBigInt, err3 := amount.Div("2.1").BigInt()
	//fmt.Println(amountBigInt, err3)
	//
	//ss, sok := amountdecimal.BigIntActualAmount(amountBigInt, 8)
	//fmt.Println(ss, sok)
	//


	d := amountdecimal.New(1)
	str,err := amountdecimal.New(d).ToString(-1)
	fmt.Println(str, err)

}
