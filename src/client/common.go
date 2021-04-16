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

/////// STEP
// Choose 10 pair -> save to db
// 1. Get ETH to USD
// 2. Get Coin to ETH 0x956f47f50a910163d8bf957cf5846d573e7f87ca  0x1405c709d6bed996d046cd94d174af7ec0c39f43
// 3. Compare coin use data in 1,2
// 4. Get 5 transaction of each pair

//// Get total totalLiquidity
/// get totalLiquidity of 2 token in pair and convert to usd

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
var RawSwapQuery = `query ($pairId: String!){
  swaps(where: {pair:$pairId} ,first: 5 , orderBy: timestamp, orderDirection: desc) {
    id
    transaction{
      	blockNumber
    }
    pair {
      	id
    }
    sender
    amount0In
    amount1In
    amount0Out
    amount1Out
    to
    amountUSD
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
  bundle(id: 1) {
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

func GetSwapById(id string) ([]model.Swap, error) {
	graphqlRequest := graphql.NewRequest(RawSwapQuery)
	// set any variables
	graphqlRequest.Var("pairId", id)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result []model.Swap
	mapstructure.Decode(graphqlResponse["swaps"], &result)

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

func GetPriceETH() ([]model.Bundle, error) {
	graphqlRequest := graphql.NewRequest(RawTokenQuery)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result []model.Bundle
	mapstructure.Decode(graphqlResponse["bundles"], &result)

	return result, err
}
