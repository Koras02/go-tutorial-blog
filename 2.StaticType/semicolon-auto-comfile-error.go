package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}

	for j := 5; j >= 0; j-- // for문 끝에 세미콜론이 삽입되 컴파일 오류가 발생한다.
	{
		fmt.Print(j)
	}
}