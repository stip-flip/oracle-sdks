package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Sender() (*common.Address, error) {
	pk := os.Getenv("PRIVATE_KEY")

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &fromAddress, nil
}

func Auth(client *ethclient.Client) (*bind.TransactOpts, error) {
	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, err := strconv.ParseInt(chainIDStr, 10, 64)

	if err != nil {
		return nil, err
	}

	// load the private key from environment variable
	pk := os.Getenv("PRIVATE_KEY")

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
