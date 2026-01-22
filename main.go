package main

import "fmt"

func main(){
const convertUsdtToEur float64 = 0.85
const convertUsdtToRub float64 = 75.63
convertUsdrInRub := convertUsdtToRub/convertUsdtToEur
fmt.Printf("%.2f\n" ,convertUsdrInRub)
}
func scanInput()(float64, float64){
var a, b float64
fmt.Println("Введите ваши числа")
fmt.Scan(&a, &b)
return a,b
}
func calculate(float64, float64, float64) (float64){
	
}