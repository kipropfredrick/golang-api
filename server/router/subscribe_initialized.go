package router

import (
	"log"
	"server/db"
	"server/internal/repository"
	"server/internal/services"
	"server/internal/testin"
	"server/internal/users"
	
)
type SetupStruct interface {
   SetupRoutes()
}

func setupRoute(){
	routers:=NewGinRouter()
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Couldnt connect to db %s",err)
	}
	userRepo:=repository.NewUserRepository(dbConn.GetDB())
	testRepo:=testin.NewTestrepo(dbConn.GetDB())
	userService:=services.NewUserService(userRepo)
	testService:=testin.NewTestservice(testRepo)
	testHnadl:=testin.NewTestHanlder(testService)
	userHandler:=users.NewHandler(userService)
    userPost:=NewUserRoute(*userHandler,*testHnadl,routers)
	userPost.Setup()
	routers.Gin.Run("0.0.0.0:8080")
}