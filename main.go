package main

import (
	"os"
	"time"

	"github.com/stip-flip/oracle-sdks/data/coingecko"
	"github.com/stip-flip/oracle-sdks/data/cryptocompare"
	"github.com/stip-flip/oracle-sdks/data/fred"
	"github.com/stip-flip/oracle-sdks/ethereum"

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
		log.Fatal(err)
	}

	log.Infof("Starting oracles on network %s", os.Getenv("CHAIN_ID"))

	sender, err := ethereum.Sender()

	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("FRED_ORACLE_ADDRESS") != "" {
		if os.Getenv("FRED_API_KEY") == "" {
			log.Fatal("FRED_API_KEY environment variable not set")
		}
		oracle, err := ethereum.LoadOracle(client, os.Getenv("FRED_ORACLE_ADDRESS"))
		if err != nil {
			log.Fatal(err)
		}
		enoughStake, err := ethereum.HasEnoughStakes(oracle, *sender)
		if err != nil {
			log.Fatal(err)
		}
		if !*enoughStake {
			log.Fatal("Not enough stake to run the FRED oracle")
		}
		log.Infof("Running FRED oracle at address %s...", os.Getenv("FRED_ORACLE_ADDRESS"))
		poll := os.Getenv("POLL_INTERVAL")
		pollDuration, err := time.ParseDuration(poll)
		if err != nil {
			log.Fatal(err)
		}
		fredTicker := time.NewTicker(pollDuration)
		defer fredTicker.Stop()

		go func() {
			for range fredTicker.C {
				fred.FetchAndSubmit(client, os.Getenv("FRED_ORACLE_ADDRESS"))
			}
		}()
	}

	if os.Getenv("CG_ORACLE_ADDRESS") != "" {
		if os.Getenv("CG_API_KEY") == "" {
			log.Fatal("CG_API_KEY environment variable not set")
		}
		oracle, err := ethereum.LoadOracle(client, os.Getenv("CG_ORACLE_ADDRESS"))
		log.Infof("Oracle address: %s", os.Getenv("CG_ORACLE_ADDRESS"))
		if err != nil {
			log.Fatal(err)
		}
		enoughStake, err := ethereum.HasEnoughStakes(oracle, *sender)
		if err != nil {
			log.Fatal(err)
		}
		if !*enoughStake {
			log.Fatal("Not enough stake to run the CoinGecko oracle")
		}
		poll := os.Getenv("POLL_INTERVAL")
		pollDuration, err := time.ParseDuration(poll)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Running CoinGecko oracle at address %s...", os.Getenv("CG_ORACLE_ADDRESS"))
		coingeckoTicker := time.NewTicker(pollDuration)
		defer coingeckoTicker.Stop()

		go func() {
			for range coingeckoTicker.C {
				coingecko.FetchAndSubmit(client, os.Getenv("CG_ORACLE_ADDRESS"))
			}
		}()
	}

	if os.Getenv("CC_ORACLE_ADDRESS") != "" {
		if os.Getenv("CC_API_KEY") == "" {
			log.Fatal("CC_API_KEY environment variable not set")
		}
		oracle, err := ethereum.LoadOracle(client, os.Getenv("CC_ORACLE_ADDRESS"))
		log.Infof("Oracle address: %s", os.Getenv("CC_ORACLE_ADDRESS"))
		if err != nil {
			log.Fatal(err)
		}
		enoughStake, err := ethereum.HasEnoughStakes(oracle, *sender)
		if err != nil {
			log.Fatal(err)
		}
		if !*enoughStake {
			log.Fatal("Not enough stake to run the CoinGecko oracle")
		}
		poll := os.Getenv("POLL_INTERVAL")
		pollDuration, err := time.ParseDuration(poll)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Running CryptoCompare oracle at address %s...", os.Getenv("CC_ORACLE_ADDRESS"))
		cryptoCompareTicker := time.NewTicker(pollDuration)
		defer cryptoCompareTicker.Stop()

		go func() {
			for range cryptoCompareTicker.C {
				cryptocompare.FetchAndSubmit(client, os.Getenv("CC_ORACLE_ADDRESS"))
			}
		}()
	}

	// Keep the main goroutine running
	select {}
}

func checkEnvVars() {
	envVars := []string{"RPC_URL", "PRIVATE_KEY", "CHAIN_ID", "POLL_INTERVAL"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("%s environment variable not set", envVar)
		}
	}
}
