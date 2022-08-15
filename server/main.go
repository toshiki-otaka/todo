package main

import (
	"fmt"
	"net/http"
	"todo/server/config"
	"todo/server/domain/model"
	"todo/server/infrastructure/dao/item"
	"todo/server/interface/handler"
	"todo/server/interface/registry"
	"todo/server/interface/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectDB() *gorm.DB {
	user := "root"
	pass := "zaqroot"
	protocol := "tcp"
	host := config.Config.Db.Host
	port := config.Config.Db.Port
	dbName := config.Config.Db.Name
	option := "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" + dbName + option

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := connectDB()
	db.LogMode(true)
	db.AutoMigrate(&model.Item{})
	defer db.Close()

	itemRepo := item.NewRepository(db)
	itemHandler := handler.NewItemHandler(itemRepo)

	handlerRegistry := registry.NewRegistry(itemHandler)

	mux := router.Init(handlerRegistry)

	// webサーバー起動
	http.ListenAndServe(":8080", mux)
}

func printLog(msg string) {
	fmt.Println(msg)
}
