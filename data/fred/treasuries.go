package fred

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/stip-flip/oracle-sdks/ethereum"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Observation struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

type Response struct {
	Observations []Observation `json:"observations"`
}

const pricePrecision = 1e6

var log = logrus.New()

var seriesIDs = []string{"DGS1MO", "DGS6MO", "DGS1", "DGS2", "DGS5", "DGS10", "DGS20", "DGS30"} // List of symbols

// QueryFred queries the FRED API for a given list of series ID and returns the observations
// for the last 7 days. The series ID is a string that identifies the data series to fetch.
// The function returns a Response struct containing the observations and an error.
// If an error occurs during the request or processing the response, the error will be
// non-nil and the Response struct will be the zero value.
func QueryFred() ([8]*big.Int, error) {
	log.Infof("fetching data from FRED for series: %v", seriesIDs)
	// Create a channel to collect the responses
	valueChan := make(chan *big.Int, len(seriesIDs))

	// Create a channel to collect any errors
	errorChan := make(chan error, len(seriesIDs))

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Loop over the seriesIDs and start a goroutine for each one
	for _, seriesID := range seriesIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			// Call the API and send the response or error to the appropriate channel
			resp, err := callAPI(id)
			if err != nil {
				errorChan <- err
				return
			}
			valueChan <- resp
		}(seriesID)
	}

	// Start a goroutine to close the channels once all the other goroutines have finished
	go func() {
		wg.Wait()
		close(valueChan)
		close(errorChan)
	}()

	// Collect the responses and errors
	var values [8]*big.Int
	var errors []error
	i := 0
	for v := range valueChan {
		values[i] = v
		i++
		if i >= 8 {
			break
		}
	}
	for err := range errorChan {
		errors = append(errors, err)
	}

	// loop through the values and make sure none is nil, if nil initialize to 0
	for i, v := range values {
		if v == nil {
			values[i] = big.NewInt(0)
		}
	}

	// If there were any errors, return the first one
	if len(errors) > 0 {
		return values, errors[0]
	}

	// Return the responses
	return values, nil
}

// callAPI makes a request to the FRED API for a given series ID and returns the last available value
func callAPI(seriesID string) (*big.Int, error) {
	// Create a struct to hold the response
	var response Response

	var result big.Int

	// fred will update the average price depending on the wall street market results,
	// so it will lag by one day
	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	log.Infof("fetching data for series: %s and date: %s", seriesID, date)

	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.stlouisfed.org/fred/series/observations?series_id=%s&observation_start=%s&api_key=%s&file_type=json", seriesID, date, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	// Defer closing the response body
	defer resp.Body.Close()

	// Read the response body into the variable body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the JSON data into the response variable
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	// If there are no observations, return an error
	if len(response.Observations) == 0 {
		return nil, fmt.Errorf("no observations found")
	}

	// Get the last observation (there should be only one)
	obs := response.Observations[len(response.Observations)-1]

	// Convert the observation value to a float64
	value, err := strconv.ParseFloat(obs.Value, 64)

	value = value * pricePrecision

	if err != nil {
		return nil, fmt.Errorf("failed to parse observation value: %w", err)
	}

	// Convert the float64 to a big.Int
	// result.SetFloat64(value)
	result.SetInt64(int64(value))

	return &result, nil
}

func FetchAndSubmit(client *ethclient.Client, oracleAddress string) {
	// Set the log format to include the timestamp in UTC
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // This is an example, you can use any format you like
	})

	// verify that the round is not 0
	oracle, err := ethereum.LoadOracle(client, oracleAddress)
	if err != nil {
		log.Errorf("Error loading oracle (%s): %s", oracleAddress, err)
		return
	}

	round, err := ethereum.IsReadyToSubmit(oracle)
	if err != nil {
		log.Infof("FRED oracle (%s) not ready to submit: %s", oracleAddress, err)
		return
	}

	// Fetch the data
	result, err := QueryFred()

	if err != nil {
		log.Errorf("Error fetching data: %s", err)
		return
	}

	err = ethereum.SetPrices(client, oracle, result, *round)

	if err != nil {
		log.Errorf("Error setting price: %s", err)
		return
	}
}
