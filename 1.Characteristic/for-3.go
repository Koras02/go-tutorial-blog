// for 문에 조건식을 생략한 방법
package main


import "fmt"


func main() {
	sum, i := 0,0

	// for문에 조건식 생략
	// 실행 결과 45
	for {
		if i >= 10 {
			// i가 10보다 크거나 같을 경우 멈춤
			break 
		} 
		sum += i
		i++
	}
	fmt.Println(sum)
}