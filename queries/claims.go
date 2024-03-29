package queries

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
)

type ClaimStruct struct {
	Claimer string    `json:"claimer"`
	Round   string    `json:"round"`
	Owner   string    `json:"owner"`
	Claimed bool      `json:"claimed"`
	Exit    bool      `json:"exit"`
	Pool    PoolField `json:"pool"`
}

type PoolField struct {
	ID     string `json:"id"`
	Oracle string `json:"oracle"`
}

type ResponseData struct {
	Claims []ClaimStruct `json:"claims"`
}

func Claims(round int) (*ResponseData, error) {
	client := graphql.NewClient(os.Getenv("GQL_URL"))

	req := graphql.NewRequest(`
			query ($claimer: String, $round: Int) {
				claims(where: {claimer: $claimer, claimed: false, round: $round}) {
					claimer
					round
					owner
					claimed
					exit
					pool {
						id
						oracle
					}
				}
			}
	`)

	req.Var("claimer", os.Getenv("PERIPHERY_ADDRESS"))
	req.Var("round", round)

	var respData ResponseData
	if err := client.Run(context.Background(), req, &respData); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &respData, nil
}

type EnterExits struct {
	Enterees []common.Address `json:"enters"`
	Exitees  []common.Address `json:"exits"`
}

// ClaimsAndSort returns the claims for a specific round
// organise each claims under a mapping with key the pool and values the enters and exits claims
func ClaimsAndSort(round int, oracle string) (*map[string]EnterExits, error) {
	resp, err := Claims(round)
	if err != nil {
		return nil, err
	}

	// organise each claims under a mapping with key the pool and values the enters and exits claims
	claimsMap := make(map[string]EnterExits)

	for _, claim := range resp.Claims {
		if claim.Pool.Oracle != oracle {
			continue
		}
		if claim.Exit {
			claimsMap[claim.Pool.ID] = EnterExits{Enterees: claimsMap[claim.Pool.ID].Enterees, Exitees: append(claimsMap[claim.Pool.ID].Exitees, common.HexToAddress(claim.Owner))}
		} else {
			claimsMap[claim.Pool.ID] = EnterExits{Enterees: append(claimsMap[claim.Pool.ID].Enterees, common.HexToAddress(claim.Owner)), Exitees: claimsMap[claim.Pool.ID].Exitees}
		}
	}

	return &claimsMap, nil
}
