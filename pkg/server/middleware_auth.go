package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/CSUOS/rabums/ent"
	"github.com/CSUOS/rabums/ent/user"
	"github.com/CSUOS/rabums/pkg/database"
	"github.com/go-chi/jwtauth"

	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var userKey = &struct {
	key string
}{"user_id"}

//AuthorizeUser middleware for validating user token and embed to userinfo to context
func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			send(w, http.StatusUnauthorized, GetReason("fail to get token"))
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			send(w, http.StatusUnauthorized, GetReason("fail to validate token"))
			return
		}

		user, ok := claims["user_id"].(string)
		if !ok {
			send(w, http.StatusUnauthorized, GetReason("fail to parse user_id from token"))
			return
		}
		if user == "" {
			send(w, http.StatusUnauthorized, GetReason("invalid user_id"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, user)

		log.Ctx(ctx).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("user_id", user)
		})

		r = r.WithContext(ctx)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

//GetUser get user from context
func GetUser(ctx context.Context) string {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return ""
	}
	user, ok := claims["user_id"].(string)
	if user == "" || !ok {
		return ""
	}
	return user
}

func GetUserObject(ctx context.Context) (*ent.User, error) {
	userID := GetUser(ctx)
	if userID == "" {
		return nil, fmt.Errorf("no user in context")
	}
	db := database.GetClient()
	return db.User.Query().Where(user.UserID(userID)).First(ctx)
}

func GenerateRegisterToken(userID, eamil string) (string, error) {
	claims := make(map[string]interface{})

	claims["user_id"] = userID
	claims["email"] = eamil
	jwtauth.SetExpiryIn(claims, time.Duration(time.Minute*30))
	jwtauth.SetIssuedNow(claims)

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecryptRegisterToken(token string) (userID, email string, err error) {

	tkn, err := tokenAuth.Decode(token)
	if err != nil {
		return
	}
	if err = jwt.Validate(tkn); err != nil {
		return
	}
	id, ok := tkn.Get("user_id")
	if !ok {
		err = fmt.Errorf("failed to get user_id from token")
		return
	}
	userID = id.(string)
	m, ok := tkn.Get("email")
	if !ok {
		err = fmt.Errorf("failed to get email from token")
		return
	}
	email = m.(string)
	return
}

//GenerateTokenCookie generate cookie which includes JWT
func GenerateTokenCookie(userID string) (*http.Cookie, error) {
	token, err := GenerateToken(userID)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
	}
	return cookie, nil
}

//GenerateToken generate token for user auth
func GenerateToken(userID string) (string, error) {
	claims := make(map[string]interface{})

	claims["user_id"] = userID
	jwtauth.SetExpiryIn(claims, time.Duration(time.Hour*1))
	jwtauth.SetIssuedNow(claims)

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
