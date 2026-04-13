package dependencies

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/infra/api/web/controller/auction_controller"
	"fullcycle-auction_go/internal/infra/api/web/controller/bid_controller"
	"fullcycle-auction_go/internal/infra/api/web/controller/user_controller"
	"fullcycle-auction_go/internal/infra/database/auction"
	"fullcycle-auction_go/internal/infra/database/bid"
	"fullcycle-auction_go/internal/infra/database/user"
	"fullcycle-auction_go/internal/usecase/auction_usecase"
	"fullcycle-auction_go/internal/usecase/bid_usecase"
	"fullcycle-auction_go/internal/usecase/user_usecase"
	"log"

	"github.com/joho/godotenv"
)

func InitDependencies(envPath string) (userUsecase user_usecase.UserUseCaseInterface,
	auctionUsecase auction_usecase.AuctionUseCaseInterface,
	bidUsecase bid_usecase.BidUseCaseInterface,
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	ctx := context.Background()

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	database, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userUsecase = user_usecase.NewUserUseCase(userRepository)
	auctionUsecase = auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository)
	bidUsecase = bid_usecase.NewBidUseCase(bidRepository, auctionRepository)

	userController = user_controller.NewUserController(userUsecase)
	auctionController = auction_controller.NewAuctionController(auctionUsecase)
	bidController = bid_controller.NewBidController(bidUsecase)

	return userUsecase, auctionUsecase, bidUsecase, userController, bidController, auctionController
}
