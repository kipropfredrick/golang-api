package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/testin"
	"server/internal/repository"
	"server/internal/services"
	"server/internal/users"
	"server/router"
)

func main() {
	routers:=router.NewGinRouter()
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Couldnt connect to db %s",err)
	}
    // Create
	userRepo:=repository.NewUserRepository(dbConn.GetDB())
	testRepo:=testin.NewTestrepo(dbConn.GetDB())
	userService:=services.NewUserService(userRepo)
	testService:=testin.NewTestservice(testRepo)
	testHnadl:=testin.NewTestHanlder(testService)
	userHandler:=users.NewHandler(userService)
    userPost:=router.NewUserRoute(*userHandler,*testHnadl,routers)
	userPost.Setup()
	// initialize the router and
    // router.InitRouter(userHandler)
	// router.Start("0.0.0.0:8080")
	routers.Gin.Run("http://34.219.53.184/")
	fmt.Println("Connecting to database")
}