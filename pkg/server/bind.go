package server

import (
	"fmt"
	"net/http"
)

func (v *ClientInfo) Bind(r *http.Request) error {
	if v.ClientName == "" || v.Link == "" {
		return fmt.Errorf("empty client name and link is not allowed")
	}
	return nil
}

func (v *LoginRequest) Bind(r *http.Request) error {
	if v.UserID == "" || v.UserPW == "" {
		return fmt.Errorf("id or pw should not be empty")
	}
	return nil
}

func (v *RequestToken) Bind(r *http.Request) error {
	if v.UserID == "" || v.Email == "" {
		return fmt.Errorf(("userid and email should not be empty"))
	}
	return nil
}

func (v *CreateUserRequest) Bind(r *http.Request) error {
	if v.UserID == "" || v.UserEmail == "" || v.Token == "" || v.UserName == "" ||
		v.UserPW == nil ||
		v.UserNumber < 1000000000 || v.UserNumber > 2030000000 {
		return fmt.Errorf(("invalid form"))
	}
	return nil
}

func (v *UpdateUserRequest) Bind(r *http.Request) error {
	if v.UserPW != nil && v.PreviousPassword == "" {
		return fmt.Errorf("to change pw, you must specify previous one")
	}
	return nil
}

func (v *GetUserInfoByClientTokenJSONBody) Bind(r *http.Request) error {
	if v.Token == "" || v.UserID == "" || v.UserPW == "" {
		return fmt.Errorf("invalid format")
	}
	return nil
}
