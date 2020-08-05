package utils

import (
	"log"

	env "github.com/Netflix/go-env"
	vault "github.com/hashicorp/vault/api"
)

//Environment 환경변수에서 받아올 값들
// "" -> default value
type Environment struct {
	// "dev" or pord
	Mode         string `env:"MODE"`
	SMTPAddress  string `env:"SMTP_ADRESS" json:"smtp_address"`
	SMTPPassword string `env:"SMTP_PASSWORD" json:"smtp_password"`
	MYSQLUri     string `env:"MYSQL_URI" json:"mysql_uri"`
	// 기타등등
	Extras env.EnvSet
}

//Config 환경변수 전역변수
var Config Environment = Environment{
	Mode:         "dev",
	SMTPAddress:  "",
	SMTPPassword: "",
	MYSQLUri:     "",
}

//LoadSettings 환경변수 로드하는 함수
func LoadSettings() {

	client, err := vault.NewClient(&vault.Config{
		Address: "https://vault.hcail.ml",
	})
	if err != nil {
		es, err := env.UnmarshalFromEnviron(&Config)
		if err != nil {
			log.Fatal(err)
		}
		Config.Extras = es
		return
	}

	c := client.Logical()
	secret, err := c.Read("/csuos/rabums")
	Config.SMTPAddress = secret.Data["smtp_address"].(string)
	Config.SMTPPassword = secret.Data["smtp_password"].(string)
	Config.MYSQLUri = secret.Data["mysql_uri"].(string)
}
