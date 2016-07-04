package irispgx

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/kataras/iris"
)

type errorHandler func(*iris.Context, string)

type Middleware struct {
	Config Config
	Pool   *pgx.ConnPool
}

func main() {
	fmt.Println("vim-go")
}

func (m *Middleware) Serve(ctx *iris.Context) {
	conn, err := m.Pool.Acquire()
	if err != nil {
		fmt.Println("Error acquiring pool in: ", ctx.Path())
	} else {
		ctx.Set("pgxConn", conn)
		ctx.Set("pgxConnPool", m.Pool)
		ctx.Next()
	}
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
