package main

import "fmt"

func main(){
const convertUsdtToEur float64 = 0.85
const convertUsdtToRub float64 = 75.63
convertUsdrInRub := convertUsdtToRub/convertUsdtToEur
fmt.Printf("%.2f\n" ,convertUsdrInRub)
}