package fred

import (
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

func TestCallAPI(t *testing.T) {
	// Call the function
	_, err := callAPI("DGS1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
