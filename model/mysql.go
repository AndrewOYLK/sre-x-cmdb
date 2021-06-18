package model

import (
	"fmt"
	"sync"

	"github.com/AndrewOYLK/ou-cmdb/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var mydb *DB

// store business info
var (
	DateFormat = "%Y-%m-%d %H:%i:%s"
	TimeFormat = "2006-01-02 15:04:05"
)

type DB struct {
	*sqlx.DB
}

func init() {
	if mydb != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		conn, err := sqlx.Connect(
			"mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
				config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Host, config.Config.Mysql.Port, config.Config.Mysql.DBName),
		)
		if err != nil {
			panic(err)
		}

		conn.SetMaxIdleConns(config.Config.Mysql.MaxIdleConn)
		conn.SetMaxOpenConns(config.Config.Mysql.MaxConn)
		mydb = &DB{conn}
	})
}
