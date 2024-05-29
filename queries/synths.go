package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

type SynthStruct struct {
	ID     string `json:"id"`
	Oracle string `json:"oracle"`
}

type SynthsData struct {
	Synths []SynthStruct `json:"synths"`
}

func Synths() (*SynthsData, error) {
	client := graphql.NewClient(os.Getenv("GQL_URL"))

	req := graphql.NewRequest(`
			query {
				synths {
					id
					oracle
				}
			}
	`)

	var synthsData SynthsData
	if err := client.Run(context.Background(), req, &synthsData); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &synthsData, nil
}
