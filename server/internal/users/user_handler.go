package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

// create handler function to create
func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// function to cfeate user
func (h *Handler) NewUserCreate(g *gin.Context) {
    var u CreateUserRequest
	if err := g.ShouldBindJSON(&u); err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
		return
	}
	res,err:=h.Service.CreateNewUser(g.Request.Context(),&u)
	if err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"internal server error":err.Error()})
		return
	}
	g.JSON(http.StatusOK,res)
}
//login handler
func (h *Handler) LoginUser(g *gin.Context) {
	var  user LoginUserRequest
if err := g.ShouldBindJSON(&user);err != nil {
	g.JSON(http.StatusBadRequest,gin.H{"internal server error":err.Error()})
}

u,err := h.Service.Login(g.Request.Context(),&user)
if err != nil {
	g.JSON(http.StatusInternalServerError,gin.H{"error loging in":err.Error()})
}
g.SetCookie("jwt",u.accesstoken,3600,"/","localhost",false,true)
resp :=&LoginUserResponse{
	Username: u.Username,
	ID: u.ID,
}
g.JSON(http.StatusOK,resp)
}
//function to logout
func (h *Handler) Logout(g *gin.Context) {
   g.SetCookie("jwt","",-1,"","",false,false)
   g.JSON(http.StatusOK,gin.H{"message": "logout successful"})
}