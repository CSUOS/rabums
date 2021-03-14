package database

import (
	"context"
	"fmt"

	"github.com/CSUOS/rabums/ent"
	"github.com/CSUOS/rabums/ent/user"
	"github.com/CSUOS/rabums/pkg/config"
	_ "github.com/go-sql-driver/mysql" //mysql plugin
	"github.com/rs/zerolog/log"
)

//DBType ent support mysql, sqlite and PostgreSQL but this project will use MySQL
const DBType = "mysql"

var client = &ent.Client{}

//Init initialize database
func Init(ctx context.Context) {

	var err error
	log.Ctx(ctx).Info().Msg("connecting to db.....")
	client, err = ent.Open(DBType, fmt.Sprintf("%s?parseTime=true", config.Database.URI))
	if err != nil {
		log.Fatal().Msgf("failed opening connection to mysql: %v", err)
	}

	log.Ctx(ctx).Info().Msg("migrating db.....")
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	createAdminUser(ctx)
}

func createAdminUser(ctx context.Context) {
	num, err := client.User.Query().
		Where(user.UserIDEQ("admin")).
		Count(ctx)
	if err != nil {
		log.Fatal().Msgf("fail to query admin user to db")
	}
	if num != 0 {
		log.Ctx(ctx).Info().Msg("admin user already exsists")
		return
	}
	log.Ctx(ctx).Info().Msg("admin user not detected creating one.....")
	user := client.User.Create().
		SetUserID("admin").
		SetUserName("admin").
		SetEmail("admin@rabums.csuos.ml").
		SetUserNumber(2020920000). // rabums created at 2020 :-D
		SetUserPw(EncryptPassword("admin")).
		SaveX(ctx)

	log.Ctx(ctx).Info().Interface("admin", user).Msg("admin user created")
}

//Close close all connections
func Close(ctx context.Context) {
	if client != nil {
		client.Close()
	}
}
