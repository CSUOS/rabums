package server

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/CSUOS/rabums/ent"
	"github.com/CSUOS/rabums/ent/clientserver"
	"github.com/CSUOS/rabums/ent/event"
	"github.com/CSUOS/rabums/ent/user"
	"github.com/CSUOS/rabums/pkg/database"
	"github.com/rs/zerolog/log"
)

//Server base server struct
type Server struct{}

// GetClientList get client info
// (GET /v1/client)
func (s Server) GetClientList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := GetUser(ctx)
	db := database.GetClient()

	clientServers, err := db.User.Query().
		Where(user.UserID(userID)).
		QueryOwns().
		All(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason("internal server error:%s", err.Error()))
		return
	}

	// use make to print empty list
	res := make([]ClientInfo, 0)
	for _, c := range clientServers {
		res = append(res, ClientInfo{
			ClientID:    &c.ID,
			ClientName:  c.ClientName,
			Description: c.Description,
			Link:        c.Link,
			Token:       &c.Token,
			Valid:       c.Available,
		})
	}

	send(w, http.StatusOK, res)
	return
}

// CreateClient request for registration
// (PUT /v1/client)
func (s Server) CreateClient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &ClientInfo{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}

	db := database.GetClient()

	c, err := db.ClientServer.Query().Where(clientserver.ClientName(req.ClientName)).All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query clientserver")
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}
	if len(c) != 0 {
		send(w, http.StatusNotAcceptable, GetReason("%q is already exists", req.ClientName))
		return
	}

	u, err := GetUserObject(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason("%q fail to get userinfo", GetUser(ctx)))
		return
	}

	server, err := db.ClientServer.Create().
		SetClientName(req.ClientName).
		SetDescription(req.Description).
		SetLink(req.Link).
		SetAvailable(true).
		SetToken(GenerateNewToken()).
		AddOwner(u).
		AddOwnerIDs(1). // admin user
		Save(ctx)

	res := ClientInfo{
		ClientID:    &server.ID,
		ClientName:  server.ClientName,
		Description: server.Description,
		Link:        server.Link,
		Token:       &server.Token,
		Valid:       server.Available,
	}

	_, err = db.Event.Create().SetEvent("rabums_client_created").SetUser(u).SetClientserver(server).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	send(w, http.StatusAccepted, res)
	return
}

// GetUserInfoByClientToken get user info with client token
// (POST /v1/client/user)
func (s Server) GetUserInfoByClientToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &GetUserInfoByClientTokenJSONBody{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}
	db := database.GetClient()
	c, err := db.ClientServer.Query().Where(clientserver.Token(req.Token)).First(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason("bad token"))
		return
	}
	u, err := db.User.Query().Where(user.UserID(req.UserID)).First(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason("bad user id or user pw"))
		return
	}
	if !database.IsEqualPassword(u.UserPw, req.UserPW) {
		send(w, http.StatusBadRequest, GetReason("bad user id or user pw"))
		return
	}
	_, err = db.Event.Create().SetEvent("logged in").SetClientserver(c).SetUser(u).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}
	send(w, http.StatusAccepted, UserInfo{
		Id:         &u.ID,
		UserEmail:  u.Email,
		UserID:     u.UserID,
		UserName:   u.UserName,
		UserNumber: u.UserNumber,
	})
	return
}

// DeleteClient delete client
// (DELETE /v1/client/{clientID})
func (s Server) DeleteClient(w http.ResponseWriter, r *http.Request, clientID string) {
	ctx := r.Context()
	db := database.GetClient()

	server, err := db.ClientServer.Query().Where(clientserver.HasOwnerWith(user.UserID(GetUser(ctx)))).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			send(w, http.StatusNotFound, GetReason("requested client not found"))
		}
		send(w, http.StatusBadRequest, GetReason("%q, fail to get userinfo", GetUser(ctx)))
		return
	}
	err = db.ClientServer.DeleteOne(server).Exec(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	u, err := GetUserObject(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}
	_, err = db.Event.Create().SetEvent("rabums_cliend_deleted").SetMessage(server.ClientName).SetUser(u).SetClientserver(server).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	send(w, http.StatusAccepted, nil)
	return
}

// LoginUser Logs in user
// (POST /v1/login)
func (s Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &LoginRequest{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}

	db := database.GetClient()
	u, err := db.User.Query().Where(
		user.UserID(req.UserID),
	).First(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Interface("request", req).Msg("fail to found user")
		if ent.IsNotFound(err) {
			send(w, http.StatusBadRequest, GetReason("id or pw not correct"))
			return
		}
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	if u.UserPw != database.EncryptPassword(req.UserPW) {
		send(w, http.StatusBadRequest, GetReason("id or pw not correct"))
		return
	}

	cookie, err := GenerateTokenCookie(u.UserID)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Interface("request", req).Msg("fail to write jwt")
		send(w, http.StatusInternalServerError, GetReason("signing error"))
		return
	}

	_, err = db.Event.Create().SetEvent("rabums_logged_in").SetUser(u).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	http.SetCookie(w, cookie)
	send(w, http.StatusOK, GetReason("Welcome %q :-)", req.UserID))

	log.Ctx(ctx).Info().Str("user_id", req.UserID).Msg("logged in")

	return
}

// LogoutUser Log out user by clear cookie
// (GET /v1/logout)
func (s Server) LogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "jwt",
		Value:  "",
		MaxAge: -1,
	})
	send(w, http.StatusOK, GetReason("you are logged out. :-)"))
	return
}

// GetUserLogs get login history for user
// (GET /v1/logs)
func (s Server) GetUserLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	u, err := GetUserObject(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}
	events, err := u.QueryEvents().Order(ent.Desc(event.FieldCreatedAt)).Limit(100).All(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
	}
	var res Logs
	for _, e := range events {
		res = append(res, struct {
			Event string "json:\"event\""
			Time  string "json:\"time\""
		}{
			Event: e.Event,
			Time:  e.CreatedAt.Local().String(),
		})
	}
	send(w, http.StatusOK, res)
	return
}

// PingPong ping check
// (GET /v1/ping)
func (s Server) PingPong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := database.GetClient()
	_, err := db.User.Query().First(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
	}
	send(w, http.StatusOK, GetReason("good. :-)"))
	return
}

// RequestToken request token for registration
// (POST /v1/token)
func (s Server) RequestToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &RequestToken{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}
	tkn, err := GenerateRegisterToken(req.UserID, req.Email)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	// need to send token by email
	log.Ctx(ctx).Info().Str("token", tkn).Msg("new token")
	//
	send(w, http.StatusAccepted, nil)
	return
}

// GetUser get user info
// (GET /v1/user)
func (s Server) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	u, err := GetUserObject(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason(err.Error()))
	}
	send(w, http.StatusAccepted, UserInfo{
		Id:         &u.ID,
		UserEmail:  u.Email,
		UserID:     u.UserID,
		UserName:   u.UserName,
		UserNumber: u.UserNumber,
	})
	return
}

// UpdateUser update user info
// (POST /v1/user)
func (s Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &UpdateUserRequest{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}
	u, err := GetUserObject(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}

	update := u.Update()

	if req.UserName != "" {
		update = update.SetUserName(req.UserName)
	}
	if req.UserPW != nil {
		if !database.IsEqualPassword(u.UserPw, req.PreviousPassword) {
			send(w, http.StatusBadRequest, GetReason("wrong password"))
			return
		}
		update = update.SetUserPw(database.EncryptPassword(*req.UserPW))
	}
	u, err = update.Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	db := database.GetClient()
	_, err = db.Event.Create().SetEvent("rabums_user_info_updated").SetUser(u).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	send(w, http.StatusAccepted, UserInfo{
		Id:         &u.ID,
		UserEmail:  u.Email,
		UserID:     u.UserID,
		UserName:   u.UserName,
		UserNumber: u.UserNumber,
	})
	return
}

// CreateUser request for registration
// (PUT /v1/user)
func (s Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &CreateUserRequest{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, GetReason(err.Error()))
		return
	}

	id, email, err := DecryptRegisterToken(req.Token)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Interface("request", req).Msg("invalid token")
		send(w, http.StatusBadRequest, GetReason("invalid token"))
		return
	}
	if id != req.UserID || email != req.UserEmail {
		send(w, http.StatusBadRequest, GetReason("invalid token"))
		return
	}

	db := database.GetClient()
	us, err := db.User.Query().Where(
		user.Or(
			user.UserID(req.UserID),
			user.Email(req.UserEmail),
			user.UserNumber(req.UserNumber),
		),
	).All(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}
	if len(us) != 0 {
		send(w, http.StatusConflict, GetReason("%q, %q, %q already exists.", req.UserID, req.UserEmail, req.UserNumber))
		return
	}

	u, err := db.User.Create().
		SetUserID(req.UserID).
		SetUserPw(database.EncryptPassword(*req.UserPW)).
		SetUserNumber(req.UserNumber).
		SetEmail(req.UserEmail).
		SetUserName(req.UserName).Save(ctx)

	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	_, err = db.Event.Create().SetEvent("rabums_user_created").SetUser(u).Save(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, GetReason(err.Error()))
		return
	}

	send(w, http.StatusAccepted, UserInfo{
		Id:         &u.ID,
		UserEmail:  u.Email,
		UserID:     u.UserID,
		UserName:   u.UserName,
		UserNumber: u.UserNumber,
	})
	return
}
