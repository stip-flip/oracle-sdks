package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/stip-flip/oracle-sdks/ethereum"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var coinIDs = []string{"ethereum-classic", "bitcoin", "ethereum", "doge", "monero", "solana", "binancecoin", "cardano"}

var decimals = map[string]float64{
	"ethereum-classic": 6, // EthereumClassic will be used as USD, so we want the price of USD in ETC, probably below 0
	"bitcoin":          4,
	"ethereum":         5,
	"dogecoin":         7,
	"monero":           6,
	"solana":           6,
	"binancecoin":      5,
	"cardano":          7,
}

var baseURL = "https://api.coingecko.com/api/v3/coins/"

var log = logrus.New()

// QueryCoinGecko queries the CoinGecko API for the current price of a list of coins, and returns the prices in ETC and in a list, as integers
// each price is multiplied by 1e(decimals) to convert it to an integer
func QueryCoingecko(apiKey string) (*[8]*big.Int, error) {
	log.Info("Fetching CoinGecko data...")

	prices := [8]*big.Int{}
	// first find out the etc price
	etcPrice, err := fetchCoinPrice("ethereum-classic", apiKey)
	if err != nil {
		return nil, err
	}

	for i, coinID := range coinIDs {
		if i == 0 {
			prices[0] = etcPrice
			continue
		}
		// wait 1 minute to avoid rate limiting
		time.Sleep(1 * time.Minute)

		price, err := fetchCoinPrice(coinID, apiKey, float64(math.Pow(10, decimals["etc"]))/float64(etcPrice.Int64()))

		if err != nil {
			return nil, err
		}
		prices[i] = price
	}

	for _, price := range prices {
		println(price)
	}

	return &prices, nil
}

func fetchCoinPrice(coinID string, apiKey string, etcPrice ...float64) (*big.Int, error) {
	params := url.Values{}
	// params.Set("id", coinID)
	params.Set("vs_currencies", "usd")
	params.Set("precision", "12")
	params.Set("key", apiKey)
	// make sure the time is between 00:00 and 1am UTC
	if time.Now().UTC().Hour() > 1 {
		return nil, fmt.Errorf("time is not between 00:00 and 1am UTC")
	}
	// date := time.Now().UTC().AddDate(0, 0, -1)
	date := time.Now().UTC()

	fmt.Println("date", date.Format("02-01-2006"))

	params.Set("date", date.Format("02-01-2006"))
	params.Set("localization", "false")
	// Construct the URL for the specific coin
	url := baseURL + coinID + "/history?" + params.Encode()
	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price for coin %s: %s", coinID, err)
	}
	defer resp.Body.Close()

	// Read the response body into the variable body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for coin %s: %s", coinID, err)
	}

	// Unmarshal the JSON data into a map
	var coinData map[string]interface{}
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response for coin %s: %s", coinID, err)
	}

	coinDataStr := fmt.Sprintf("%v", coinData)
	println("coindata", coinDataStr)

	// Extract the price from the coin data
	price, ok := coinData["market_data"].(map[string]interface{})["current_price"].(map[string]interface{})["usd"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to extract price for coin %s", coinID)
	}

	// Print the price for debugging purposes
	log.Infof("Price of %s: %f", coinID, price)

	if coinID == "ethereum-classic" {
		fmt.Println(1e6/price, big.NewInt(int64(math.Pow(10, decimals[coinID])/price)))
		return big.NewInt(int64(math.Pow(10, decimals[coinID]) / price)), nil
	} else {
		fmt.Println(big.NewInt(int64(price * 1e9 / etcPrice[0])))
		return big.NewInt(int64(price / etcPrice[0] * math.Pow(10, decimals[coinID]))), nil
	}
}

func FetchAndSubmit(client *ethclient.Client, oracleAddress string) {
	// get the coingecko api key
	apiKey := os.Getenv("OG_API_KEY")
	// Set the log format to include the timestamp in UTC
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // This is an example, you can use any format you like
	})

	// verify that the round is not 0
	oracle, err := ethereum.LoadOracle(client, oracleAddress)
	if err != nil {
		log.Errorf("Error loading Oracle (%s): %s", oracleAddress, err)
		return
	}

	round, err := ethereum.IsReadyToSubmit(oracle)
	if err != nil {
		log.Infof("Coingecko oracle (%s) not ready to submit: %s", oracleAddress, err)
		return
	}

	// Fetch the data
	result, err := QueryCoingecko(apiKey)

	if err != nil {
		log.Errorf("Error fetching data: %s", err)
		return
	}

	err = ethereum.SetPrices(client, oracle, *result, *round)

	if err != nil {
		bn, _ := client.BlockNumber(context.TODO())
		log.Errorf("Error setting price at block %d: %s", bn, err)
		return
	}
	log.Info("Prices submitted successfully")
}
