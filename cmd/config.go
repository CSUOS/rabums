package main

import (
	"log"

	env "github.com/Netflix/go-env"
)

//Environment 환경변수에서 받아올 값들
// "" -> default value
type Environment struct {
	// "dev" or pord
	Mode string `env:"MODE"`
	// 기타등등
	Extras env.EnvSet
}

//Env 환경변수 전역변수
var Env Environment

//LoadSettings 환경변수 로드하는 함수
func LoadSettings() {
	es, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
	Env.Extras = es
}
