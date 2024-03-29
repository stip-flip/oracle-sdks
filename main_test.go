package main

import (
	"fmt"
	"os"
	"sf-peripheries/queries"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file")
	}

	// Run the tests
	os.Exit(m.Run())
}

func TestOnlyMain(t *testing.T) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))

	if err != nil {
		log.Fatalf("Error connecting ethereum client %s", err.Error())
	}

	// fetch all the pools
	pools, err := queries.Pools()

	if err != nil {
		log.Fatalf("Error querying the pools %s", err.Error())
	}

	// join the pools in a mapping under their common oracle
	oraclePools := make(map[string][]string)
	for _, pool := range pools.Pools {
		oraclePools[pool.Oracle] = append(oraclePools[pool.Oracle], pool.ID)
	}

	performClaiming(client, oraclePools)
}
