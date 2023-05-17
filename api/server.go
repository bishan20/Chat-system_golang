package api

import (
	db "Chat-system_golang/db/sqlc"
	"Chat-system_golang/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}

	router.Use(cors.New(config))

	routerVersionOne := router.Group("/v1")
	{
		routerVersionOne.GET("/ws", server.handleWebSocket)
		routerVersionOne.POST("/users", server.createUser)
		routerVersionOne.GET("/hello", server.hello)

		messageRouter := routerVersionOne.Group("/messages")
		{
			messageRouter.POST("/", server.listMessages)
			messageRouter.PATCH("/", server.updateMessage)
			messageRouter.DELETE("/", server.deleteMessage)
		}
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type Err struct {
	Error string `json:"error"`
}

// error response function
func errorResponse(err error) Err {

	return Err{Error: err.Error()}
}

// type SuccessResponse struct {
// 	Message string `json:"message"`
// }

// // success response function
// func successResponse(msg string) SuccessResponse {

// 	return SuccessResponse{Message: msg}
// }
