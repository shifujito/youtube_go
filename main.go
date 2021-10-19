package main

import (
	"youtube/my"

	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var dbDriver = "sqlite3"
var dbName = "data.sqlite3"

var sesName = "youtube-session"
var cs = sessions.NewCookieStore([]byte("secret-key-1234"))

func checkLogin() {
	my.Migrate()
}
