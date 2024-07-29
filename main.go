package main

import (
	"glinboy.com/gin-crud-rest-api-demo/db"
	"glinboy.com/gin-crud-rest-api-demo/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
