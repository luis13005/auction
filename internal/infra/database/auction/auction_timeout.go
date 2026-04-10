package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"

	"go.mongodb.org/mongo-driver/bson"
)

func (ar *AuctionRepository) SetAuctionTimeOut(
	ctx context.Context,
	auction *auction_entity.Auction) *internal_error.InternalError {

	fmt.Println("auction: ", auction)

	filter := bson.M{"_id": auction.Id}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error trying to set auction timeout", err)
		return internal_error.NewInternalServerError("Error trying to set auction timeout")
	}
	return nil
}
