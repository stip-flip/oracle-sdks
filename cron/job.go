package main

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stip-flip/oracle-sdks/data/cryptocompare"
	"github.com/stip-flip/oracle-sdks/ethereum"
)

var log = logrus.New()

func checkEnvVars() {
	envVars := []string{"RPC_URL", "PRIVATE_KEY", "CHAIN_ID"}

	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("%s environment variable not set", envVar)
		}
	}
}

func main() {
	// Set the log format to include the timestamp in UTC
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // This is an example, you can use any format you like
	})

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}

	err = godotenv.Load()
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

	if os.Getenv("CC_ORACLE_ADDRESS") != "" {
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

		log.Infof("Running CryptoCompare oracle at address %s...", os.Getenv("CC_ORACLE_ADDRESS"))

		cryptocompare.FetchAndSubmit(client, os.Getenv("CC_ORACLE_ADDRESS"))
	}
}
