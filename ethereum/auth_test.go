package ethereum

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Could not load .env file")
	}

	// Run the tests
	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatalf("Could not load the eth client, got %v", err)
	}
	// Call the function
	_, err = Auth(client)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSetPrice(t *testing.T) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatalf("Could not load the eth client, got %v", err)
	}
	// Load the oracle
	oracle, _ := LoadOracle(client, os.Getenv("CG_ORACLE_ADDRESS"))
	// Call the function
	_ = SetPrices(client, oracle, [8]*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6), big.NewInt(7), big.NewInt(8)}, 1)
}
