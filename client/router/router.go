package router

import (
	"voting/client/handler"

	"github.com/gin-gonic/gin"
)

// Init returns a new router
func Init() *gin.Engine {
	router := gin.Default()

	router.POST("/vote/create", handler.CreateVote)
	router.DELETE("/vote", handler.DeleteVote)
	router.POST("/vote/update", handler.UpdateVote)
	router.GET("/vote/:voteID", handler.GetVote)
	// router.GET("/vote", handler.ListVote)
	router.GET("/vote/:voteID/result", handler.GetResult)

	router.POST("/ballot/add", handler.AddBallot)
	router.POST("/ballot/cast", handler.CastBallot)
	router.POST("/ballot/count", handler.CountBallot)
	// router.GET("/ballot/:voteID/:code", handler.GetBallot)
	// router.GET("/ballot/:voteID", handler.ListBallot)

	return router
}
