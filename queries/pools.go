package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

type PoolStruct struct {
	ID     string `json:"id"`
	Oracle string `json:"oracle"`
}

type PoolsData struct {
	Pools []PoolStruct `json:"pools"`
}

func Pools() (*PoolsData, error) {
	client := graphql.NewClient(os.Getenv("GQL_URL"))

	req := graphql.NewRequest(`
			query {
				pools {
					id
					oracle
				}
			}
	`)

	var poolsData PoolsData
	if err := client.Run(context.Background(), req, &poolsData); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &poolsData, nil
}
