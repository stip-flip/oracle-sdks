package main

import (
	"context"
	"os"
	"sf-peripheries/ethereum"
	"sf-peripheries/queries"
	"time"

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

	// every 1 minutes, try to perform a claiming
	ticker := time.NewTicker(2 * time.Minute)

	defer ticker.Stop()

	go func() {
		for range ticker.C {
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
	}()

	// Keep the main goroutine running
	select {}
}

func performClaiming(client *ethclient.Client, oraclePools map[string][]string) {
	log.Info("Checking for claims")
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

		log.Infof("Found %d claims", len(*claimsMap))
		// for each pool in the oracle, perform a claimAll
		for pool, enterExits := range *claimsMap {
			// perform the claimAll
			err = ethereum.ClaimAll(client, enterExits.Enterees, enterExits.Exitees, lastRound-1, pool)
			if err != nil {
				bn, _ := client.BlockNumber(context.TODO())
				log.Errorf("Failed to claimAll for pool %s at block number %d: %s", pool, bn, err.Error())
			}
		}
	}
}

func checkEnvVars() {
	envVars := []string{"PERIPHERY_ADDRESS", "PRIVATE_KEY", "CHAIN_ID", "RPC_URL", "GQL_URL"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("%s environment variable not set", envVar)
		}
	}
}
