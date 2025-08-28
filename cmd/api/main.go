package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandrowiemesfilho/login-api/internal/configs"
	"github.com/leandrowiemesfilho/login-api/internal/db"
	"github.com/leandrowiemesfilho/login-api/internal/routes"
)

func main() {
	log.Println("Starting application...")

	log.Println("Loading configurations...")
	conf := configs.LoadConfigs()

	log.Println("Connecting to database...")
	mongoDB := db.NewMongoDB(conf)
	defer mongoDB.Close()

	r := gin.Default()

	routes.Setup(r)

	if err := r.Run(); err != nil {
		log.Fatal("Application did not started with error: ", err)
	}

	log.Println("Application finished successfully")
}
