package services

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Event types
const (
	EventProductListed   = "ProductListed"
	EventProductBought   = "ProductBought"
	EventProductAuthenticated = "ProductAuthenticated"
	EventPriceUpdated    = "PriceUpdated"
)

// StartEventListeners starts listening for contract events
func (s *Web3Service) StartEventListeners(ctx context.Context) error {
	// Create event query
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.contractAddr},
	}

	// Subscribe to logs
	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %v", err)
	}

	// Start event processing
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
			case vLog := <-logs:
				// Parse event
				event, err := s.contract.ParseEvent(vLog)
				if err != nil {
					log.Printf("Failed to parse event: %v", err)
					continue
				}

				// Handle event based on type
				switch event.EventType {
				case EventProductListed:
					if e, ok := event.(*ReVibeContractProductListed); ok {
						s.handleProductListed(e)
					}
				case EventProductBought:
					if e, ok := event.(*ReVibeContractProductBought); ok {
						s.handleProductBought(e)
					}
				case EventProductAuthenticated:
					if e, ok := event.(*ReVibeContractProductAuthenticated); ok {
						s.handleProductAuthenticated(e)
					}
				case EventPriceUpdated:
					if e, ok := event.(*ReVibeContractPriceUpdated); ok {
						s.handlePriceUpdated(e)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

// handleProductListed handles the ProductListed event
func (s *Web3Service) handleProductListed(event *ReVibeContractProductListed) {
	log.Printf("Product listed: ID=%s, Name=%s, Price=%s, Seller=%s",
		event.ProductId.String(),
		event.Name,
		event.Price.String(),
		event.Seller.Hex())
}

// handleProductBought handles the ProductBought event
func (s *Web3Service) handleProductBought(event *ReVibeContractProductBought) {
	log.Printf("Product bought: ID=%s, Buyer=%s, Price=%s",
		event.ProductId.String(),
		event.Buyer.Hex(),
		event.Price.String())
}

// handleProductAuthenticated handles the ProductAuthenticated event
func (s *Web3Service) handleProductAuthenticated(event *ReVibeContractProductAuthenticated) {
	log.Printf("Product authenticated: ID=%s, Result=%v",
		event.ProductId.String(),
		event.Result)
}

// handlePriceUpdated handles the PriceUpdated event
func (s *Web3Service) handlePriceUpdated(event *ReVibeContractPriceUpdated) {
	log.Printf("Price updated: ID=%s, NewPrice=%s",
		event.ProductId.String(),
		event.NewPrice.String())
}

// GetPastEvents retrieves past events
func (s *Web3Service) GetPastEvents(ctx context.Context, fromBlock, toBlock *big.Int) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{s.contractAddr},
	}

	return s.client.FilterLogs(ctx, query)
} 