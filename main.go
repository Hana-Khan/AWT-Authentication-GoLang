// Implemented JWT Authentication in Golang REST APIs and secured it with Authentication Middleware. 
// Built a simple, yet neatly organized Golang REST API with packages like Gin for Routing (mostly), GORM for persisting user data to a MySQL Database, and so on. 
// Also built a Middleware that can secure Endpoints and allow requests that have a valid JWT in the Request’s Authorization Header.
package main
import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
	"github.com/gin-gonic/gin"
)
func main() {
	// Initialize Database
	database.Connect("root:@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
	// Initialize Router=> method wrote below
	router := initRouter()
	// we run the API server at Port 8080.
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	// Creates a new Gin Router instance.
	router := gin.Default()
	// we need a couple of routes as below.
	// api/user/register
	// api/token
	// api/secured/ping
	// api/secured/something-else
	api := router.Group("/api")
	{
		// we grouped everything under /api. 
		// Then, we routed the api/token to the GenerateToken function that we wrote in the tokencontroller. 
		// Similarly, for the user registration endpoint too.
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)

		// Now, we need to secure all the endpoints that will come under the api/secured/ routes. 
		// Here is where we tell GIN to use the middleware that we created. 
		//  we use the Auth middleware that will be attached to this particular set of endpoints. 
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}