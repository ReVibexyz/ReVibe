package utils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ParseAddress parses an Ethereum address from a string
func ParseAddress(address string) (common.Address, error) {
	if !common.IsHexAddress(address) {
		return common.Address{}, fmt.Errorf("invalid Ethereum address: %s", address)
	}
	return common.HexToAddress(address), nil
}

// ParseBigInt parses a big.Int from a string
func ParseBigInt(value string) (*big.Int, error) {
	value = strings.TrimPrefix(value, "0x")
	bigInt := new(big.Int)
	if _, ok := bigInt.SetString(value, 16); !ok {
		return nil, fmt.Errorf("invalid big.Int value: %s", value)
	}
	return bigInt, nil
}

// GenerateSignature generates a signature for a message
func GenerateSignature(privateKey *ecdsa.PrivateKey, message []byte) (string, error) {
	hash := crypto.Keccak256(message)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// Adjust V value
	signature[64] += 27

	return common.Bytes2Hex(signature), nil
}

// VerifySignature verifies a signature
func VerifySignature(address common.Address, message []byte, signature string) (bool, error) {
	sig := common.FromHex(signature)
	if len(sig) != 65 {
		return false, fmt.Errorf("invalid signature length")
	}

	// Adjust V value
	sig[64] -= 27

	hash := crypto.Keccak256(message)
	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr == address, nil
}

// CreateAuth creates a new auth transactor
func CreateAuth(privateKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %v", err)
	}

	// Set gas price and limit
	auth.GasPrice = big.NewInt(20000000000) // 20 Gwei
	auth.GasLimit = uint64(3000000)

	return auth, nil
}

// FormatEther formats Wei to Ether
func FormatEther(wei *big.Int) string {
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, big.NewFloat(1e18))
	return ether.Text('f', 18)
}

// ParseEther parses Ether to Wei
func ParseEther(ether string) (*big.Int, error) {
	value := new(big.Float)
	if _, ok := value.SetString(ether); !ok {
		return nil, fmt.Errorf("invalid ether value: %s", ether)
	}

	wei := new(big.Float).Mul(value, big.NewFloat(1e18))
	weiInt := new(big.Int)
	wei.Int(weiInt)
	return weiInt, nil
}

// IsValidAddress checks if an address is valid
func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// ShortenAddress shortens an Ethereum address
func ShortenAddress(address string) string {
	if !IsValidAddress(address) {
		return address
	}
	return address[:6] + "..." + address[len(address)-4:]
} 