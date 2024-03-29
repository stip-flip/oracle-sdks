package ethereum

import (
	"context"
	"fmt"
	"os"
	"sf-peripheries/store"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var loggr = logrus.New()

func LoadTrader(client *ethclient.Client, address string) (*store.Trader, error) {
	// Load the trader periphery contract
	formattedAddress := common.HexToAddress(address)

	instance, err := store.NewTrader(formattedAddress, client)

	if err != nil {
		return nil, fmt.Errorf("failed to load Trader periphery contract: %w", err)
	}

	return instance, nil
}

func ClaimAll(client *ethclient.Client, enterees []common.Address, exitees []common.Address, round uint64, pool string) error {
	auth, err := Auth(client)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	trader, err := LoadTrader(client, os.Getenv("PERIPHERY_ADDRESS"))

	if err != nil {
		return fmt.Errorf("failed to load trader: %w", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	poolAddress := common.HexToAddress(pool)

	auth.NoSend = true

	transaction, err := trader.ClaimAll(auth, enterees, exitees, round, poolAddress)

	if err != nil {
		return fmt.Errorf("transaction incorrect: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: auth.From,
		To:   transaction.To(),
		Data: transaction.Data(),
	}

	_, err = client.EstimateGas(ctx, callMsg)

	if err != nil {
		// Log the list of addresses
		entereeStrings := make([]string, len(enterees))
		for i, addr := range enterees {
			entereeStrings[i] = addr.Hex()
		}
		exiteeStrings := make([]string, len(exitees))
		for i, addr := range exitees {
			exiteeStrings[i] = addr.Hex()
		}

		return fmt.Errorf(
			"failed to estimate gas for claimAll with enterees: %v, exitees, %v, round: %d: %w",
			entereeStrings,
			exiteeStrings,
			round,
			err,
		)
	}

	auth.NoSend = false

	loggr.Info("Claiming all rewards")

	// Claim all rewards
	transaction, err = trader.ClaimAll(auth, enterees, exitees, round, poolAddress)

	if err != nil {
		return fmt.Errorf("transaction incorrect: %w", err)
	}

	_, err = bind.WaitMined(ctx, client, transaction)
	if err != nil {
		return fmt.Errorf("failed to claim: %w", err)
	}

	return nil
}
