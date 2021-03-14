// +build tools

package ent

import (
	_ "entgo.io/ent/cmd/ent" // generating binary
)

//go:generate go run entgo.io/ent/cmd/ent generate ./schema
//go:generate go run entgo.io/ent/cmd/ent describe ./schema
