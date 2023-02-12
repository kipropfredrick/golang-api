package router

import (
	"server/internal/testin"
	"server/internal/users"
)

type PostRoute struct {
	UserHandler users.Handler
	TestHandl testin.TestHanlder
	handler GinRouter
}
//func to hanlde all routes

func NewUserRoute(
  uHandler users.Handler,
  TestHanlder testin.TestHanlder,
  handler GinRouter,
)PostRoute {
	return PostRoute{
		UserHandler: uHandler,
		TestHandl:  TestHanlder,
		handler:handler,
	}
}
//set up route function
func (p PostRoute)Setup(){
   allroutes:=p.handler.Gin.Group("create/user")
   {
	allroutes.POST("/create/u",p.UserHandler.NewUserCreate)
	allroutes.POST("/login/",p.UserHandler.LoginUser)
	allroutes.GET("/logout/",p.UserHandler.Logout)
	allroutes.POST("/create/test",p.TestHandl.NewTestCreate)
   }
}