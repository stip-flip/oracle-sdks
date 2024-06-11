package cryptocompare

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Warn("Could not load .env file")
	}

	// Run the tests
	os.Exit(m.Run())
}

func TestQueryCryptocompare(t *testing.T) {
	// Call the function
	prices, err := QueryCryptocompare(os.Getenv("CC_API_KEY"))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fmt.Println(prices)
}
