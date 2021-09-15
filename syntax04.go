package main

import (
	"fmt"
	"math/rand"
	"time"
)

//난수 추출된 수의 소수 판정 프로그램 v0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	check := false

	number := rand.Intn(150) + 2 // 0과 1제외, 2~151 사이의 수
	//fmt.Println(number)
	//number = 47
	for i := 2; i < number; i++ {
		if (number % i) == 0 {
			check = true
		}
	}
	if check == false {
		fmt.Print(number, "는(은) 솟수입니다!")
	} else {
		fmt.Print(number, "는(은) 솟수가 아닙니다!")
	}
}
