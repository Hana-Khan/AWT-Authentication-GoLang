package controllers
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// we would secure this endpoint so that only the requests having a valid JWT at the Request header will be able to access this. 
// And  it just returns a pong message with a 200 status code.
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}