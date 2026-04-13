package main

import (
	"fullcycle-auction_go/cmd/dependencies"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	_, _, _, userController, bidController, auctionsController := dependencies.InitDependencies("cmd/auction/.env")

	router.GET("/auction", auctionsController.FindAuctions)
	router.GET("/auction/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auction", auctionsController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}
