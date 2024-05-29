package ethereum

import (
	"fmt"
	"sf-peripheries/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LoadPool loads the pool contract
func LoadPool(client *ethclient.Client, address string) (*store.Pool, error) {
	// Load the pool periphery contract
	formattedAddress := common.HexToAddress(address)

	instance, err := store.NewPool(formattedAddress, client)

	if err != nil {
		return nil, fmt.Errorf("failed to load Pool periphery contract: %w", err)
	}

	return instance, nil
}
