package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/stip-flip/oracle-sdks/store"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func HasEnoughStakes(oracle *store.Oracle, sender common.Address) (*bool, error) {
	// check the stakes
	stakes, err := oracle.Stakes(nil, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get stakes: %w", err)
	}
	minStake, err := oracle.MinStake(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get min stake: %w", err)
	}
	boolValue := false
	// check if the stakes are less than the minimum
	if minStake.Cmp(stakes) > 0 {
		return &boolValue, nil
	}
	boolValue = true
	return &boolValue, nil
}

func IsReadyToSubmit(oracle *store.Oracle) (*uint64, error) {
	round, err := oracle.GetCurrentRound(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get current round: %w", err)
	}
	if round == 0 {
		return nil, fmt.Errorf("round is 0")
	}

	sender, err := Sender()
	if err != nil {
		return nil, fmt.Errorf("failed to get sender: %w", err)
	}
	// verify that wa have not already submitted for this round
	// if we have, return
	p, err := oracle.Submitters(&bind.CallOpts{}, round, *sender)
	if err != nil {
		return nil, fmt.Errorf("failed to get submitter: %w", err)
	}
	// check if p is zero
	if p.Cmp(big.NewInt(0)) != 0 {
		return nil, fmt.Errorf("already submitted for this round")
	}
	return &round, nil
}

func SetPrices(client *ethclient.Client, oracle *store.Oracle, values [8]*big.Int, round uint64) error {
	auth, err := Auth(client)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	slotValues, err := oracle.SetSlots(nil, values)
	if err != nil {
		return fmt.Errorf("failed to set slots: %w", err)
	}

	// estimate the gas or if the transaction would fail
	auth.NoSend = true

	transaction, _ := oracle.SetPrices(auth, slotValues, round)

	callMsg := ethereum.CallMsg{
		From: auth.From,
		To:   transaction.To(),
		Data: transaction.Data(),
	}

	_, err = client.EstimateGas(context.Background(), callMsg)

	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w from %s", err, auth.From.String())
	}

	auth.NoSend = false

	transaction, err = oracle.SetPrices(auth, slotValues, round)
	if err != nil {
		return fmt.Errorf("failed to set prices: %w", err)
	}

	_, err = bind.WaitMined(ctx, client, transaction)
	if err != nil {
		return fmt.Errorf("failed to set prices: %w", err)
	}

	return nil
}
