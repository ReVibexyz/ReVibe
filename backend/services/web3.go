package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yourusername/revibe/backend/config"
)

// Web3Service handles blockchain interactions
type Web3Service struct {
	client       *ethclient.Client
	contract     *ReVibeContract
	contractAddr common.Address
	chainID      *big.Int
}

// NewWeb3Service creates a new Web3Service instance
func NewWeb3Service() (*Web3Service, error) {
	// Connect to Ethereum network
	client, err := ethclient.Dial(fmt.Sprintf("https://mainnet.infura.io/v3/%s", config.AppConfig.InfuraID))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum network: %v", err)
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Parse contract address
	contractAddr := common.HexToAddress(config.AppConfig.ContractAddress)

	// Create contract instance
	contract, err := NewReVibeContract(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %v", err)
	}

	return &Web3Service{
		client:       client,
		contract:     contract,
		contractAddr: contractAddr,
		chainID:      chainID,
	}, nil
}

// Close closes the Web3Service connection
func (s *Web3Service) Close() {
	if s.client != nil {
		s.client.Close()
	}
}

// GetAuth creates an authentication transactor
func (s *Web3Service) GetAuth(privateKey string) (*bind.TransactOpts, error) {
	// Parse private key
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	// Create auth
	auth, err := bind.NewKeyedTransactorWithChainID(key, s.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %v", err)
	}

	return auth, nil
}

// ListProduct lists a product on the blockchain
func (s *Web3Service) ListProduct(auth *bind.TransactOpts, name string, price *big.Int) (string, error) {
	tx, err := s.contract.ListProduct(auth, name, price)
	if err != nil {
		return "", fmt.Errorf("failed to list product: %v", err)
	}

	return tx.Hash().Hex(), nil
}

// BuyProduct purchases a product
func (s *Web3Service) BuyProduct(auth *bind.TransactOpts, productID *big.Int) (string, error) {
	tx, err := s.contract.BuyProduct(auth, productID)
	if err != nil {
		return "", fmt.Errorf("failed to buy product: %v", err)
	}

	return tx.Hash().Hex(), nil
}

// AuthenticateProduct authenticates a product
func (s *Web3Service) AuthenticateProduct(auth *bind.TransactOpts, productID *big.Int) (string, error) {
	tx, err := s.contract.AuthenticateProduct(auth, productID)
	if err != nil {
		return "", fmt.Errorf("failed to authenticate product: %v", err)
	}

	return tx.Hash().Hex(), nil
}

// UpdatePrice updates a product's price
func (s *Web3Service) UpdatePrice(auth *bind.TransactOpts, productID *big.Int, newPrice *big.Int) (string, error) {
	tx, err := s.contract.UpdatePrice(auth, productID, newPrice)
	if err != nil {
		return "", fmt.Errorf("failed to update price: %v", err)
	}

	return tx.Hash().Hex(), nil
}

// GetProduct retrieves product details
func (s *Web3Service) GetProduct(productID *big.Int) (*Product, error) {
	product, err := s.contract.GetProduct(nil, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	return &Product{
		ID:          productID.String(),
		Name:        product.Name,
		Price:       product.Price.String(),
		Seller:      product.Seller.Hex(),
		IsListed:    product.IsListed,
		IsAuthenticated: product.IsAuthenticated,
	}, nil
}

// GetProductPrice retrieves a product's price
func (s *Web3Service) GetProductPrice(productID *big.Int) (*big.Int, error) {
	price, err := s.contract.GetProductPrice(nil, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product price: %v", err)
	}

	return price, nil
}

// GetProductSeller retrieves a product's seller address
func (s *Web3Service) GetProductSeller(productID *big.Int) (common.Address, error) {
	seller, err := s.contract.GetProductSeller(nil, productID)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get product seller: %v", err)
	}

	return seller, nil
}

// GetProductAuthentication retrieves a product's authentication status
func (s *Web3Service) GetProductAuthentication(productID *big.Int) (bool, error) {
	isAuthenticated, err := s.contract.GetProductAuthentication(nil, productID)
	if err != nil {
		return false, fmt.Errorf("failed to get product authentication: %v", err)
	}

	return isAuthenticated, nil
} 