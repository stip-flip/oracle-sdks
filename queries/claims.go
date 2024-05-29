package queries

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"sf-peripheries/ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
)

type ClaimStruct struct {
	Claimer string     `json:"claimer"`
	Round   string     `json:"round"`
	Owner   string     `json:"owner"`
	Claimed bool       `json:"claimed"`
	Exit    bool       `json:"exit"`
	Pool    SynthField `json:"pool"`
}

type SynthField struct {
	ID     string `json:"id"`
	Oracle string `json:"oracle"`
}

type ClaimsData struct {
	Claims []ClaimStruct `json:"claims"`
}

func Claims(round int) (*ClaimsData, error) {
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

	var respData ClaimsData
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

type PositionClaimStruct struct {
	Claimer string     `json:"claimer"`
	Round   string     `json:"round"`
	Owner   string     `json:"owner"`
	Tick    int        `json:"tick"`
	Claimed bool       `json:"claimed"`
	Burn    bool       `json:"burn"`
	Pool    SynthField `json:"pool"`
}

type PositionClaimsData struct {
	PositionClaims []PositionClaimStruct `json:"positionClaims"`
}

func PositionClaims(round int) (*PositionClaimsData, error) {
	client := graphql.NewClient(os.Getenv("GQL_URL"))

	req := graphql.NewRequest(`
			query ($claimer: String, $round: Int) {
				positionClaims(where: {claimer: $claimer, round: $round, claimed: false}) {
					id
					owner
					tick
					burn
					amount
					round
					claimer
					claimed
					pool {
						id
						oracle
					}
				}
			}
	`)

	req.Var("claimer", os.Getenv("PERIPHERY_ADDRESS"))
	req.Var("round", round)

	var respData PositionClaimsData
	if err := client.Run(context.Background(), req, &respData); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &respData, nil
}

type EnterExitsPosition struct {
	Enterees [][32]byte `json:"enters"`
	Exitees  [][32]byte `json:"exits"`
}

func PositionClaimsAndSort(round int, oracle string, client *ethclient.Client) (*map[string]EnterExitsPosition, error) {
	resp, err := PositionClaims(round)
	if err != nil {
		return nil, err
	}

	// organise each claims under a mapping with key the pool and values the enters and exits claims
	claimsMap := make(map[string]EnterExitsPosition)

	for _, claim := range resp.PositionClaims {
		if claim.Pool.Oracle != oracle {
			continue
		}

		pool, err := ethereum.LoadPool(client, claim.Pool.ID)

		t := big.NewInt(int64(claim.Tick))

		pack, err := pool.Pack(&bind.CallOpts{}, t, common.HexToAddress(claim.Owner))

		if err != nil {
			return &claimsMap, err
		}

		if claim.Burn {
			claimsMap[claim.Pool.ID] = EnterExitsPosition{Enterees: claimsMap[claim.Pool.ID].Enterees, Exitees: append(claimsMap[claim.Pool.ID].Exitees, pack)}
		} else {
			claimsMap[claim.Pool.ID] = EnterExitsPosition{Enterees: append(claimsMap[claim.Pool.ID].Enterees, pack), Exitees: claimsMap[claim.Pool.ID].Exitees}
		}
	}

	return &claimsMap, nil
}

const (
	MAX_INT24  = 0x7FFFFF
	MAX_UINT24 = 0xFFFFFF
)

func pack(int24 int32, address common.Address) [32]byte {
	int24Bytes := [4]byte{}
	copy(int24Bytes[:], new(big.Int).SetInt64(int64(int24)).Bytes())

	bytes := [32]byte{}
	for i := 0; i < 20; i++ {
		bytes[i] = address[i]
	}
	for i := 0; i < 4; i++ {
		bytes[i+12] = int24Bytes[i]
	}

	return bytes
}
