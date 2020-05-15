package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func getBase64From(user, password string) string {
	s := fmt.Sprintf("%s:%s", user, password)
	e := base64.StdEncoding.EncodeToString([]byte(s))
	return e
}

func getStringFrom(enc string) string {
	d, e := base64.StdEncoding.DecodeString(enc)
	if e != nil {
		panic(e)
	}

	return string(d)
}

type Auths struct {
	Registries RegistriesStruct `json:"auths"`
}

type RegistriesStruct map[string]struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Auth     string `json:"auth"`
}

func main() {
	a := &Auths{
		Registries: RegistriesStruct{
			".dockerconfigjson": struct {
				Username string `json:"username"`
				Password string `json:"password"`
				Email    string `json:"email"`
				Auth     string `json:"auth"`
			}{
				Username: "user123",
				Password: "123445555",
				Email:    "a@a.com",
				Auth:     getBase64From("user123", "123445555"),
			},
		},
	}

	fmt.Println(a.Registries[".dockerconfigjson"].Auth)
	fmt.Println(getStringFrom(a.Registries[".dockerconfigjson"].Auth))

	b, e := json.Marshal(a)
	if e != nil {
		panic(e)
	}

	fmt.Println(string(b))
}
