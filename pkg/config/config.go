package config

import "context"

type ModeStatus string

const (
	ModeLocal  ModeStatus = "LOCAL"
	ModeDeploy ModeStatus = "DEPLOY"
	ModeProd   ModeStatus = "PROD"
)

var (
	Mode      = ModeLocal
	SecretKey = "1q2w3e4r"
	Database  = struct {
		URI  string
		Salt string
	}{
		URI:  "root:root@tcp(localhost:3306)/rabums",
		Salt: "e109hjq0dn",
	}
)

func Init(ctx context.Context) {
}
