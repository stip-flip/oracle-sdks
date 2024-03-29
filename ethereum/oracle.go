package ethereum

import (
	"fmt"
	"sf-peripheries/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadOracle(client *ethclient.Client, address string) (*store.Oracle, error) {
	// Load the oracle
	formattedAddress := common.HexToAddress(address)

	instance, err := store.NewOracle(formattedAddress, client)

	if err != nil {
		return nil, fmt.Errorf("failed to load Oracle contract: %w", err)
	}

	return instance, nil
}
