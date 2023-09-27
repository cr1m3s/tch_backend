package repositories

import "github.com/cr1m3s/tch_backend/queries"

func NewAppRepository() *queries.Queries {
	conn := ConnectDataBase()
	db := queries.New(conn)

	return db
}
