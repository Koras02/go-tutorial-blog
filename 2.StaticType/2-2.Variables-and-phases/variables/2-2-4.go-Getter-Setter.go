// 특정 객체를 반호나하는 함수나 메서드의 이름은 
// 명사형으로 짓고 Get 접수어를 붙이지 않는다.

func Connection() *Conn { 
	// GetConnection()으로 사용하지 않음
	// ...
	return conn
}

// 특정 개체를 변형하거나 설정하는 함수의 이름 앞에는
// Set 포함

func SetId(i int) {...}