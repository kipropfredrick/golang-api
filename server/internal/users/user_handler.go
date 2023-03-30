package users

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"server/internal/models"
)

type Handler struct {
	models.Service
}

// create handler function to create
func NewHandler(s models.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// function to cfeate user
func (h *Handler) NewUserCreate(g *gin.Context) {
    var u models.CreateUserRequest
	
	if err := g.ShouldBindJSON(&u); err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
		return
	}
	// g.JSON(http.StatusAccepted,u);
	// return;
	
	res,err:=h.Service.CreateNewUser(g.Request.Context(),&u)
	if err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"internal server error":err.Error()})
		return
	}
	// fmt.Println("done",u);
	g.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"succefuly created","data":res})
}
//login handler
func (h *Handler) LoginUser(g *gin.Context) {
	var  user models.LoginUserRequest
if err := g.ShouldBindJSON(&user);err != nil {
	g.JSON(http.StatusBadRequest,gin.H{"internal server error":err.Error()})
}

u,err := h.Service.Login(g.Request.Context(),&user)
if err != nil {
	g.JSON(http.StatusInternalServerError,gin.H{"error loging in":err.Error()})
}

g.SetCookie("jwt",u.Accesstoken,3600,"/","localhost",false,true)
resp :=&models.LoginUserResponse{
	Accesstoken: u.Accesstoken,
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
func (h *Handler) Home(g *gin.Context) {
	// g.SetCookie("jwt","",-1,"","",false,false)
	g.JSON(http.StatusOK,gin.H{"message": "endpoint working"})
 }