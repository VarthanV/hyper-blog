package main

import (
	"log"

	"github.com/VarthanV/hyper-todo/controllers"
	"github.com/VarthanV/hyper-todo/models"
	hyper "github.com/VarthanV/hyper/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	h := hyper.New()
	h.GET("/ping", func(w hyper.ResponseWriter, request *hyper.Request) {
		w.WriteJSON(200, map[string]string{"message": "PONG"})
	})

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("error in auto migrating models ", err)
	}

	c := controllers.Controller{DB: db}

	h.POST("/posts", c.CreatePost)
	h.GET("/posts", c.GetAllPosts)

	h.ListenAndServe("localhost", "6060", `
 _  _  _  _  ____  ____  ____      ____  __     __    ___ 
/ )( \( \/ )(  _ \(  __)(  _ \ ___(  _ \(  )   /  \  / __)
) __ ( )  /  ) __/ ) _)  )   /(___)) _ (/ (_/\(  O )( (_ \
\_)(_/(__/  (__)  (____)(__\_)    (____/\____/ \__/  \___/

`)
}
