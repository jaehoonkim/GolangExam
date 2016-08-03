package bcrypt_test

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 생성된 해쉬와 원문과 비교가 가능한가?
// 결과: 비교가 정상적으로 잘 된다.
func ExampleCreateHash() {
	pw := "user_password"

	hp, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	log.Println(string(hp))

	err = bcrypt.CompareHashAndPassword(hp, []byte(pw))
	fmt.Println(err)
	// Output:
	// <nil>
}

// 동일한 원문으로 두번 생성된 해쉬들과 원문이 비교가 가능한가?
// 결과: 동일한 원문으로 매번 발생이 되는 해쉬들은 모두 비교가 가능했다.
// 그래서 salt가 필요함...
func ExampleTwiceHash() {
	pw := "user_password"

	hp, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	log.Println(string(hp))

	hp2, err2 := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err2)
	}
	log.Println(string(hp2))

	err = bcrypt.CompareHashAndPassword(hp, []byte(pw))
	fmt.Println(err)

	err = bcrypt.CompareHashAndPassword(hp2, []byte(pw))
	fmt.Println(err)
	// Output:
	// <nil>
	// <nil>
}
