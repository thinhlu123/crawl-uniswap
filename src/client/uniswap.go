package client

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/mitchellh/mapstructure"
	"thinhlu123/crawl-uniswap/src/model"
)

var UniswapClient *graphql.Client

func InitUniSwapClient() {
	//src := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	//)
	//httpClient := oauth2.NewClient(context.Background(), src)

	UniswapClient = graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")
}

var RawPairQuery = `query ($id: String!){
	pairs (where:{id: $id}) {
      	id
		token0{
			symbol
			name
			totalLiquidity
			derivedETH
		}
		token1{
			symbol
			name
			totalLiquidity
			derivedETH
		}
		token0Price
		token1Price
    }
}`

// RawSwapQuery Get transaction of Pair by id
var RawTokenTransactionQuery = `query($allPairs: String!) {
 mints(first: 5, where: { pair: $allPairs }, orderBy: timestamp, orderDirection: desc) {
   transaction {
     id
     timestamp
   }
   timestamp
   to
   liquidity
   amount0
   amount1
   amountUSD
 }
 burns(first: 5, where: { pair: $allPairs }, orderBy: timestamp, orderDirection: desc) {
   transaction {
     id
     timestamp
   }
   timestamp
   to
   liquidity
   amount0
   amount1
   amountUSD
 }
 swaps(first: 5, where: { pair: $allPairs }, orderBy: timestamp, orderDirection: desc) {
   transaction {
     id
     timestamp
   }
   timestamp
   amount0In
   amount0Out
   amount1In
   amount1Out
   amountUSD
   to
 }
}`

// RawTokenQuery get price of coin to eth
var RawTokenQuery = `query ($id: String!){
  tokens (where:{id:$id}) {
    derivedETH
  }
}`

// GetETHPriceQuery use to convert eth to usd
var GetETHPriceQuery = `{
  bundle(id: "1") {
    ethPrice 
  } 
}`

func GetPairById(id string) ([]model.Pair, error) {
	graphqlRequest := graphql.NewRequest(RawPairQuery)
	// set any variables
	graphqlRequest.Var("id", id)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result []model.Pair
	mapstructure.Decode(graphqlResponse["pairs"], &result)

	return result, err
}

func GetTokenTransactionById(id string) ([]model.TokenTransaction, error) {
	graphqlRequest := graphql.NewRequest(RawTokenTransactionQuery)
	// set any variables
	graphqlRequest.Var("allPairs", id)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var swaps []model.TokenTransaction
	mapstructure.Decode(graphqlResponse["swaps"], &swaps)
	for i := 0; i < len(swaps); i++ {
		swaps[i].Type = "SWAP"
		swaps[i].PairId = id
	}

	var burns []model.TokenTransaction
	mapstructure.Decode(graphqlResponse["burns"], &burns)
	for i := 0; i < len(burns); i++ {
		burns[i].Type = "BURN"
		burns[i].PairId = id
	}

	var mints []model.TokenTransaction
	mapstructure.Decode(graphqlResponse["mints"], &mints)
	for i := 0; i < len(mints); i++ {
		mints[i].Type = "MINT"
		mints[i].PairId = id
	}

	var result = append(swaps, burns...)
	result = append(result, mints...)

	//sort.Slice(result, func(i, j int) bool {
	//	a, _ := strconv.Atoi(result[i].Timestamp)
	//	b, _ := strconv.Atoi(result[j].Timestamp)
	//	return a > b
	//})

	return result, err
}

func GetTokenById(id string) ([]model.Token, error) {
	graphqlRequest := graphql.NewRequest(RawTokenQuery)
	// set any variables
	graphqlRequest.Var("id", id)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result []model.Token
	mapstructure.Decode(graphqlResponse["tokens"], &result)

	return result, err
}

func GetPriceETH() (*model.Bundle, error) {
	graphqlRequest := graphql.NewRequest(GetETHPriceQuery)

	var graphqlResponse map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result model.Bundle
	mapstructure.Decode(graphqlResponse["bundle"], &result)

	return &result, err
}
