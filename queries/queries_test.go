package queries

import (
	"fmt"
	"os"
	"testing"

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

func TestClaims(t *testing.T) {
	lastRound := 150
	Claims(int(lastRound - 1))
}

func TestClaimsAndOrder(t *testing.T) {
	ClaimsAndSort(149, "0x05abc9884d19f7a10f0fa94f9ced65c30ff05b92")
}

func TestPools(t *testing.T) {
	Pools()
}
