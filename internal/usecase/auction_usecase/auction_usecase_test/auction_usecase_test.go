package auctionusecasetest

import (
	"context"
	"fullcycle-auction_go/cmd/dependencies"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/usecase/auction_usecase"
	"fullcycle-auction_go/internal/usecase/bid_usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuctionExpiration(t *testing.T) {
	_, auctionUsecase, bidUsecase, _, _, _ := dependencies.InitDependencies("../../../../cmd/auction/.env")

	ctx := context.Background()

	input := auction_usecase.AuctionInputDTO{
		ProductName: "Produto Teste",
		Category:    "Eletronicos",
		Description: "Descrição longa ",
		Condition:   auction_usecase.ProductCondition(auction_entity.New),
	}

	err := auctionUsecase.CreateAuction(ctx, input)
	assert.Nil(t, err, "erro ao criar auction")

	auctions, err := auctionUsecase.FindAuctions(ctx, auction_usecase.AuctionStatus(auction_entity.Active), "Eletronicos", "Produto Teste")

	assert.Nil(t, err, "erro ao buscar auctions")
	assert.NotEmpty(t, auctions, "nenhum auction encontrado")
	auction := auctions[len(auctions)-1]

	assert.Equal(t, auction_entity.Active, auction_entity.AuctionStatus(auction.Status), "Auction should be active.")

	err = bidUsecase.CreateBid(ctx, bid_usecase.BidInputDTO{
		UserId:    uuid.New().String(),
		AuctionId: auction.Id,
		Amount:    1500,
	})

	assert.Nil(t, err)

	time.Sleep(auction_usecase.GetAuctionDuration())

	expiredAuction, err := auctionUsecase.FindAuctionById(ctx, auction.Id)
	assert.Nil(t, err, "Error finding auction by id")
	assert.Equal(t, auction_entity.Completed, auction_entity.AuctionStatus(expiredAuction.Status), "auction should be Completed")

	err = bidUsecase.CreateBid(ctx, bid_usecase.BidInputDTO{
		UserId:    uuid.New().String(),
		AuctionId: auction.Id,
		Amount:    2000,
	})

	assert.NotNil(t, err)
}
