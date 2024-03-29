package ethereum

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func TestClaimAll(t *testing.T) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// PERIPHERY := "0x10DFC6EA1f62E4639EEE08Ea2034B927e4915324"

	POOL := "0xaC73C4A3BF4B9f5a98dbf0De7E0D83E885d1D291"

	// transform PERIPHERY and POOL addresses into common.Address
	err = ClaimAll(
		client,
		[]common.Address{common.HexToAddress("0x62310C887d21bAc7e3E3F7c080e5b517f495c1f5")},
		[]common.Address{},
		894,
		POOL,
	)

	if err != nil {
		t.Fatalf("Failed to claim all: %v", err)
	}
}
