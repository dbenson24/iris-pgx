package irispgx

import (
	"fmt"

	"github.com/jackc/pgx"
	"github.com/kataras/iris/v12"
)

type errorHandler func(iris.Context, string)

type Middleware struct {
	Config Config
	Pool   *pgx.ConnPool
}

func main() {
	fmt.Println("vim-go")
}

func (m *Middleware) Serve(ctx iris.Context) {
	if m.Config.AttachPool {
		ctx.Values().Set(m.Config.PoolCtxKey, m.Pool)
	}

	if m.Config.AttachConn {
		conn, err := m.Pool.Acquire()
		if err != nil {
			fmt.Println("Error acquiring pool in: ", ctx.Path())
		} else {
			fmt.Print("Setting the Connection")
			ctx.Values().Set(m.Config.ConnCtxKey, conn)
		}
	}
	ctx.Next()

	//_, ok := ctx.Get("pgx").(pgx.Conn)
}

func New(cfg ...Config) *Middleware {
	var c Config
	if len(cfg) == 0 {
		c = Config{}
	} else {
		c = cfg[0]
	}
	pool, _ := pgx.NewConnPool(c.ConnPoolConfig)
	return &Middleware{
		Config: c,
		Pool:   pool,
	}
}
