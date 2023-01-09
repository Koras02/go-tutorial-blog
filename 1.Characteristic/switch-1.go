// switch 문 case 조건식 사용예
package main

import "fmt"

func main() {
	c := 'H'
	switch {
		// case 에 조건식을 사용
	   case '0' <= c && c <= '9':
		 fmt.Printf("%c은(는) 숫자입니다.",c)
	   case 'a' <= c && c <= 'z':
		 fmt.Printf("%c은(는) 소문자입니다.", c)
	   case 'A' <= c && c <= 'Z':
		  fmt.Printf("%c은(는) 대문자 입니다.", c)
	}
}