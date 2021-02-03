package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver      string
	Addr        string
	Port        int
	User        string
	Pass        string
	LibName     string
	MaxIdleConn int
	MaxOpenConn int
}

func InitDatabase(cfg Config) *sql.DB {
	var dsn string
	switch cfg.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=30s", cfg.User, cfg.Pass, cfg.Addr, cfg.Port, cfg.LibName)
	default:
		panic("unsupported db driver: " + cfg.Driver)
	}

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err)
	}
	if cfg.MaxIdleConn > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}
	if cfg.MaxOpenConn > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
