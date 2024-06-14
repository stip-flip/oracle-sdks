package cryptocompare

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

var coinIDs = []string{"etc", "btc", "eth", "doge", "xmr", "sol", "bnb", "ada"}

var decimals = map[string]float64{
	"etc":  7, // EthereumClassic will be used as USD, so we want the price of USD in ETC, probably below 0
	"btc":  2,
	"eth":  3,
	"doge": 8,
	"xmr":  5,
	"sol":  5,
	"bnb":  4,
	"ada":  7,
}

var baseURL = "https://min-api.cryptocompare.com/data/v2/histoday"

var log = logrus.New()

func QueryCryptocompare(apiKey string) (*[8]*big.Int, error) {
	log.Info("Fetching CryptoCompare data...")

	prices := [8]*big.Int{}
	// first find out the etc price
	etcPrice, err := fetchCoinPrice("etc", apiKey)
	if err != nil {
		return nil, err
	}

	for i, coinID := range coinIDs {
		if i == 0 {
			prices[0] = etcPrice
			continue
		}

		// wait 1 minute to avoid rate limiting
		// time.Sleep(1 * time.Minute)

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
	params.Set("fsym", coinID)
	params.Set("tsym", "usd")
	params.Set("limit", "1")
	params.Set("api_key", apiKey)
	// make sure the time is between 00:00 and 1am UTC
	// if time.Now().UTC().Hour() > 1 {
	// 	return nil, fmt.Errorf("time is not between 00:00 and 1am UTC")
	// }

	// Construct the URL for the specific coin
	url := baseURL + "?" + params.Encode()
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
	price, ok := coinData["Data"].(map[string]interface{})["Data"].([]interface{})[1].(map[string]interface{})["open"].(float64)

	if !ok {
		return nil, fmt.Errorf("failed to extract price for coin %s", coinID)
	}

	timestamp, ok := coinData["Data"].(map[string]interface{})["Data"].([]interface{})[1].(map[string]interface{})["time"].(float64)

	if !ok {
		return nil, fmt.Errorf("failed to extract timestamp for coin %s", coinID)
	}

	// Convert the Unix timestamp to a time.Time object
	dateFromTimestamp := time.Unix(int64(timestamp), 0).UTC().Format("2006-01-02")

	date := time.Now().UTC().Format("2006-01-02")

	if dateFromTimestamp != date {
		return nil, fmt.Errorf("date is not today")
	}

	// Print the price for debugging purposes
	log.Infof("Price of %s: %f", coinID, price)

	if coinID == "etc" {
		return big.NewInt(int64(math.Pow(10, decimals[coinID]) / price)), nil
	} else {
		return big.NewInt(int64(price * math.Pow(10, decimals[coinID]) / etcPrice[0])), nil
	}
}

func FetchAndSubmit(client *ethclient.Client, oracleAddress string) {
	// get the coingecko api key
	apiKey := os.Getenv("CC_API_KEY")
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
	result, err := QueryCryptocompare(apiKey)

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
