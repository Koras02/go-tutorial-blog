package main

import "fmt"

func main() {
	sum, i := 0, 0
	for i < 10 {
		sum += i++ // 증감 연산자는 반환값이 없다, 컴파일 에러가 발생
	}
	fmt.Println(sum)
}