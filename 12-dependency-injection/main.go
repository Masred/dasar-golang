package main

import (
	"github.com/Masred/dasar-golang/golang-dependency-injection/helper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
