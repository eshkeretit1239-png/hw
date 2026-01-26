package main

import ("fmt"
		"strings"
		"slices")

func main(){
arr :=[]float32{}
slc := arr[:]
sum := float32(0)
sum1 := float32(0)
for{ 
	slcm := input()
	if slcm == 0{
		break
	}
	slc = append(slc, slcm)
}
fmt.Println(slc)
var operation string
for{
	fmt.Println("Введите операцию AVG - среднее, SUM - сумму, MED - медиану")
	fmt.Scanln(&operation)
	if strings.ToLower(operation) == "avg"{
		for _, v := range slc{
			sum += v
		}
		result := sum/float32(len(slc))
		fmt.Println("AVG :", result)	
	} else if strings.ToLower(operation) == "sum"{
		for _, v1 := range slc{
			sum1 += v1
		}
		fmt.Println("SUM :", sum1)
	} else if strings.ToLower(operation) == "med"{
			slices.Sort(slc)
		if len(slc) %2 == 0{
				result:= (slc[len(slc)/2-1] + slc[len(slc)/2])/2
				fmt.Println("MED: ", result)
		} else if len(slc)%2 == 1{
			result := slc[len(slc)/2]
			fmt.Println("MED: ", result)
		}
	}
			fmt.Println("Хотите ли продолжит?")
		var answer bool = cicle()
		if answer == false {
			break
}
}
}
func input()float32{
	for{
	fmt.Println("Ведиете ваши числа, что бы закончить введите нуль")
	var number float32
	fmt.Scanln(&number)
	return number
	}
}
func cicle() bool {
	var answer string
	for {
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "yes" {
			return true
		} else if strings.ToLower(answer) == "no" {
			return false
		} else {
			fmt.Println("Введите yes или no")
		}
	}
}
