// 변수 선언
/*
   변수 이름과 타입 표기
  
   var 변수명 타입
   */
var a int 
var b string 

// 변수 여러개 한번에 선언
var name, id, address string 

// 타입이 다른 변수를 한꺼번에 선언시 소괄호
var (
   name string 
   age int
   weight float 
)

// 변수 타입 생략 
var c = true

// 특정 타입으로 변환된 값 할당
var size = unit16(1024)

// 짧은 선언 
// 변수 선언과 동시에 값을 할당 시 var을 생략 가능

// var 생략시 := 연산자 사용
start := 1

// go 코드 컨벤션
if v := getValue() {
	fmt.Println(v)
}