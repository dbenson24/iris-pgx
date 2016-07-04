package irispgx

import "github.com/jackc/pgx"

type Config struct {
	ConnPoolConfig pgx.ConnPoolConfig
	AttachPool     bool
	AttachConn     bool
	PoolCtxKey     string
	ConnCtxKey     string
}
