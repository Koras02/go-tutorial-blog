package main

import "fmt"

func main() {
	sum, i := 0,0

	// for 문에 조건식만 사용
	// 실행 결과 45
	for i < 10 {
		sum += i
		i++
	}

	fmt.Println(sum)
}