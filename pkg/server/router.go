package server

import (
	"context"
	"net/http"
	"time"

	"github.com/CSUOS/rabums/pkg/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

var tokenAuth *jwtauth.JWTAuth

//Init initialize router
func Init(ctx context.Context) *chi.Mux {
	tokenAuth = jwtauth.New("HS256", []byte(config.SecretKey), nil)

	var si Server
	r := chi.NewRouter()

	options := ChiServerOptions{
		BaseURL:     "/api",
		BaseRouter:  r,
		Middlewares: []MiddlewareFunc{},
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Use(middleware.RealIP)

	// Logging
	r.Use(hlog.NewHandler(log.Logger))
	r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		log.Ctx(r.Context()).Info().
			Str("method", r.Method).
			Str("user_agent", r.Header.Get("User-Agent")).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	// Panic Recover
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// Server
	r.Get(options.BaseURL+"/v1/ping", wrapper.PingPong)
	r.Post(options.BaseURL+"/v1/login", wrapper.LoginUser)
	r.Get(options.BaseURL+"/v1/logout", wrapper.LogoutUser)
	r.Post(options.BaseURL+"/v1/token", wrapper.RequestToken)
	r.Put(options.BaseURL+"/v1/user", wrapper.CreateUser)

	// Need authorization
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(AuthorizeUser)

		r.Get(options.BaseURL+"/v1/logs", wrapper.GetUserLogs)
		r.Get(options.BaseURL+"/v1/user", wrapper.GetUser)
		r.Post(options.BaseURL+"/v1/user", wrapper.UpdateUser)

		r.Get(options.BaseURL+"/v1/client", wrapper.GetClientList)
		r.Put(options.BaseURL+"/v1/client", wrapper.CreateClient)
		r.Post(options.BaseURL+"/v1/client/user", wrapper.GetUserInfoByClientToken)
		r.Delete(options.BaseURL+"/v1/client/{clientID}", wrapper.DeleteClient)
	})
	return r
}
