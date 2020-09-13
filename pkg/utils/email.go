package utils

import (
	"encoding/base64"
	"log"
	"net/smtp"
)

//SendTokenViaEmail 주어진 이메일 주소로 토큰을 보낸다.
func SendTokenViaEmail(address string, token string) error {
	return send(address, "[CSUOS] 이메일을 인증해주세요!",
		`<body><h2>메일인증 안내입니다.</h2>
<p>CSUOS 프로젝트 중 하나인 RABUMS를 통해 회원가입 하시는것을 환영합니다.</p>
<p>RABUMS는 CSUOS에서 진행되는 프로젝트들의 통합 회원 관리 시스템입니다.</p>
<p>CSUOS에서 진행되는 프로젝트는 모두 오픈소스로 진행되며, 프로젝트 참여에 관심이 있으시다면</p>
<p><a href='https://github.com/CSUOS' target='_blank' class='url'>https://github.com/CSUOS</a></p>
<p>위 링크를 참고해주세요.</p>
<p>&nbsp;</p>
<p>회원가입 진행을 위해서는 다음 토큰을 복사해서 RABUMS에 등록시켜주시면 됩니다.</p>
<blockquote style="background: whitesmoke;padding: 5px;"><p>`+token+`</p>
</blockquote>
<p>감사합니다. :-)</p>
<p>&nbsp;</p>
<p><b>Team. CSUOS</b></p>
</body>
	`)
}

func send(address string, title string, body string) error {
	from := Config.SMTPAddress
	pass := Config.SMTPPassword
	to := address

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: =?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(title)) + "?=\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n" +
		"Content-Transfer-Encoding: base64\n\n" +
		base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return nil
}
