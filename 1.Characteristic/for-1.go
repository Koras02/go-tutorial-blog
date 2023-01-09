// go - 1.for 반복문 1
package main

import "fmt"


func main() {
	sum :=0
	  // for문에 초기화 구문, 조건식, 후속 작업 순으로 정의
	  // 실행 결과 45
	  for i :=0; i < 10; i++ {
		sum += i
	  }
	  fmt.Println(sum)
}