package rest

import (
	"database/sql"

	"github.com/MrBrown969/restCrud/internal/boundary/model/adapter/rest"
)

type MobController struct {
	sql.Conn
}

func Spawn(request rest.MobSpawnRequest) *rest.MobSpawnResponse {
	return nil
}
