package auction_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"fullcycle-auction_go/internal/usecase/auction_usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockAuctionRepository struct {
	auctions map[string]*auction_entity.Auction
}

func (m *mockAuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	m.auctions[auction.Id] = auction
	return nil
}

func (m *mockAuctionRepository) SetAuctionTimeOut(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	m.auctions[auction.Id].Status = auction_entity.Completed
	return nil
}

func TestAuctionExpiration(t *testing.T) {
	// Seta duração curta só para o teste
	t.Setenv("AUCTION_DURATION", "2s")

	repo := &mockAuctionRepository{auctions: make(map[string]*auction_entity.Auction)}
	useCase := auction_usecase.NewAuctionUseCase(repo, nil)

	useCase.CreateAuction(context.Background(), AuctionInputDTO{
		ProductName: "Produto X",
		Category:    "Categoria X",
		Description: "Descrição",
		Condition:   0,
	})

	// Pega o ID do leilão criado e verifica status inicial
	auctionId := /* id gerado */
		assert.Equal(t, auction_entity.Active, repo.auctions[auctionId].Status)

	// Aguarda a expiração
	time.Sleep(3 * time.Second)

	assert.Equal(t, auction_entity.Completed, repo.auctions[auctionId].Status)
}
