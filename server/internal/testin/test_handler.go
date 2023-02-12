package testin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHanlder struct {
	testserv TestService
}

func NewTestHanlder(th TestService) *TestHanlder {
	return &TestHanlder{
		testserv: th,
	}
}
func (th *TestHanlder) NewTestCreate(g *gin.Context)  {
   var tres TestRequest
   if err := g.ShouldBindJSON(&tres); err != nil {
	g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
	return
}
res,err:=th.testserv.CreateTest(g.Request.Context(),&tres)
if err != nil {
	g.JSON(http.StatusInternalServerError,gin.H{"internal server error":err.Error()})
	return
}
g.JSON(http.StatusOK,res)
}