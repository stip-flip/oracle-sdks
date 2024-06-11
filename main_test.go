package main

import (
	"fmt"
	"os"
	"testing"

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
	main()
}
