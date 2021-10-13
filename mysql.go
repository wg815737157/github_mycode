package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {

	db, err := sql.Open("mysql", "")
	if err != nil {
		//日志
		fmt.Println(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		//日志
		fmt.Println(err.Error())
		return
	}

	db.SetMaxIdleConns(16)
	//db.SetMaxOpenConns(32)
	//db.SetConnMaxLifetime(8 * time.Minute)

	_, err = db.Query("select 1")
	if err != nil {
		//日志
		fmt.Println(err.Error())
		return
	}
	<-time.After(61 * time.Second)
	db.Query("select 1")

}
