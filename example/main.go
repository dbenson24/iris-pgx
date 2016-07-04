package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dbenson24/iris-pgx"
	"github.com/jackc/pgx"
	"github.com/kataras/iris"
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
	iris.Use(ConnMiddleware)
	iris.Get("/test", func(c *iris.Context) {
		conn := c.Get("PGXCONN").(*pgx.Conn)
		c.SetBodyString(strconv.FormatBool(conn.IsAlive()))
	})
	iris.Listen(":8080")
}
