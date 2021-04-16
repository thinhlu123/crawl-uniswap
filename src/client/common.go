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

var RawPairQuery = `{
  pairs {
    id
    token0{
      id
    }
    token1{
      id
    }
    reserve0
    reserve1
    token0Price
    token1Price
  }
}`

var RawDataTodayQuery = ``

func GetPair() ([]model.Pair, error) {
	graphqlRequest := graphql.NewRequest(RawPairQuery)

	var graphqlResponse map[string][]map[string]interface{}
	err := UniswapClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if err != nil {
		return nil, err
	}

	var result []model.Pair
	mapstructure.Decode(graphqlResponse["pairs"], &result)

	return result, err
}
