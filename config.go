package irispgx

import "github.com/jackc/pgx"

type Config struct {
	ConnPoolConfig pgx.ConnPoolConfig
}
