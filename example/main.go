package main

import (
	"fmt"
	"os"

	"github.com/dbenson24/iris-pgx"
	"github.com/jackc/pgx"
	"github.com/kataras/iris/v12"
)

func main() {
	fmt.Println("vim-go")
	var c irispgx.Config
	c.ConnPoolConfig.Host = os.Getenv("PGX_HOST")
	c.ConnPoolConfig.Database = os.Getenv("PGX_DB")
	c.ConnPoolConfig.User = os.Getenv("PGX_USER")
	c.ConnPoolConfig.Password = os.Getenv("PGX_PASS")
	c.ConnCtxKey = "PGXCONN"
	c.AttachConn = true
	c.AttachPool = false
	ConnMiddleware := irispgx.New(c)
	app := iris.New()
	app.Use(ConnMiddleware.Serve)
	app.Get("/test", func(c iris.Context) {
		conn := c.Values().Get("PGXCONN").(*pgx.Conn)
		c.Writef("%s",  conn.IsAlive())
	})
	app.Run(iris.Addr(":8080"))
}
