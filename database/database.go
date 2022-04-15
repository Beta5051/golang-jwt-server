package database

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func InitDB() error {
	engine, err := xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		return nil
	}
	if err := engine.Ping(); err != nil {
		return err
	}
	if err := engine.Sync2(new(User)); err != nil {
		return err
	}
	DB = engine
	return nil
}

type User struct {
	Id           int64  `json:"-"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
