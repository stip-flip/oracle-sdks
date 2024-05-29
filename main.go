package main

import (
	"context"
	"os"
	"sf-peripheries/ethereum"
	"sf-peripheries/queries"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// Set the log format to include the timestamp in UTC
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // This is an example, you can use any format you like
	})

	err := godotenv.Load()

	if err != nil {
		log.Warn("Could not load .env file")
	}

	checkEnvVars()

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))

	if err != nil {
		log.Fatalf("Error connecting ethereum client %s", err.Error())
	}

	log.Infof("Connected to ethereum client %s", os.Getenv("RPC_URL"))
	log.Infof("Starting the claiming service on network %s", os.Getenv("CHAIN_ID"))
	log.Infof("Using periphery address %s", os.Getenv("PERIPHERY_ADDRESS"))
	// every 1 minutes, try to perform a claiming
	ticker := time.NewTicker(10 * time.Second)

	defer ticker.Stop()

	var lastTransaction *types.Transaction

	go func() {
		for range ticker.C {
			// fetch all the pools
			synths, err := queries.Synths()

			if err != nil {
				log.Fatalf("Error querying the pools %s", err.Error())
			}

			// join the pools in a mapping under their common oracle
			oraclePools := make(map[string][]string)
			for _, synth := range synths.Synths {
				oraclePools[synth.Oracle] = append(oraclePools[synth.Oracle], synth.ID)
			}

			swap := claimSwaps(client, oraclePools)
			if swap != nil {
				lastTransaction = swap
			}

			position := claimPositions(client, oraclePools)
			if position != nil {
				lastTransaction = position
			}

			if lastTransaction != nil {
				bind.WaitMined(context.Background(), client, lastTransaction)
			}
		}
	}()

	// Keep the main goroutine running
	select {}
}

func claimSwaps(client *ethclient.Client, oraclePools map[string][]string) *types.Transaction {
	log.Info("Checking for claims")
	var lastTransaction *types.Transaction
	// for each oracle key in the mapping, fetch the last round
	for oracle := range oraclePools {
		oracleContract, err := ethereum.LoadOracle(client, oracle)
		if err != nil {
			log.Errorf("Error loading oracle %s", err.Error())
		}

		// get the last round
		lastRound, err := oracleContract.GetLastRound(nil)
		if err != nil {
			log.Errorf("Error getting oracle last round %s", err.Error())
		}

		log.Infof("Last round for oracle %s: %d", oracle, lastRound) // Convert lastRound to string
		// fetch all claims that qualify for this round (lastRound - 1)
		claimsMap, err := queries.ClaimsAndSort(int(lastRound-1), oracle)

		if err != nil {
			log.Errorf("Error getting active claims %s", err)
		}

		log.Infof("Found %d swap claims", len(*claimsMap))
		// for each pool in the oracle, perform a claimAll
		for pool, enterExits := range *claimsMap {
			// perform the claimAll
			t, err := ethereum.ClaimAllSwap(client, enterExits.Enterees, enterExits.Exitees, lastRound-1, pool)
			if err != nil {
				bn, _ := client.BlockNumber(context.TODO())
				log.Errorf("Failed to claimAllSwap for pool %s at block number %d: %s", pool, bn, err.Error())
			}
			if t != nil {
				lastTransaction = t
			}
		}
	}
	return lastTransaction
}

func claimPositions(client *ethclient.Client, oraclePools map[string][]string) *types.Transaction {
	log.Info("Checking for positions")
	var lastTransaction *types.Transaction
	// for each oracle key in the mapping, fetch the last round
	for oracle := range oraclePools {
		oracleContract, err := ethereum.LoadOracle(client, oracle)
		if err != nil {
			log.Errorf("Error loading oracle %s", err.Error())
		}

		// get the last round
		lastRound, err := oracleContract.GetLastRound(nil)
		if err != nil {
			log.Errorf("Error getting oracle last round %s", err.Error())
		}

		log.Infof("Last round for oracle %s: %d", oracle, lastRound) // Convert lastRound to string
		// fetch all claims that qualify for this round (lastRound - 1)
		claimsMap, err := queries.PositionClaimsAndSort(int(lastRound-1), oracle, client)

		if err != nil {
			log.Errorf("Error getting active claims %s", err)
		}

		log.Infof("Found %d position claims", len(*claimsMap))

		// for each pool in the oracle, perform a claimAll
		for pool, enterExits := range *claimsMap {
			// perform the claimAll
			t, err := ethereum.ClaimAllPosition(client, enterExits.Enterees, enterExits.Exitees, lastRound-1, pool)
			if err != nil {
				bn, _ := client.BlockNumber(context.TODO())
				log.Errorf("Failed to claimAllPosition for pool %s at block number %d: %s", pool, bn, err.Error())
			}
			if t != nil {
				lastTransaction = t
			}
		}
	}
	return lastTransaction
}

func checkEnvVars() {
	envVars := []string{"PERIPHERY_ADDRESS", "PRIVATE_KEY", "CHAIN_ID", "RPC_URL", "GQL_URL"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("%s environment variable not set", envVar)
		}
	}
}
