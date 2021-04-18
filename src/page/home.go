package page

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"thinhlu123/crawl-uniswap/src/model"
)

func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := model.PageVars{
		Title: "Crawl Uniswap",
	}

	// Get All Pair
	pair, err := model.PairCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return
	}
	var out []model.Pair
	err = pair.All(context.TODO(), &out)
	if err != nil {
		return
	} else if len(out) == 0 {
		return
	}

	var listPairInfo []model.PairInfo
	var bundle model.Bundle
	err = model.BundleCollection.FindOne(context.TODO(), bson.M{"type": "ETH_TO_USD"}).Decode(&bundle)
	if err != nil {
		fmt.Println(err)
		return
	}
	ethPrice, err := strconv.ParseFloat(bundle.ETHPrice, 64)
	for _, e := range out {
		var pairInfo model.PairInfo
		pairInfo.PairId = e.PairId
		pairInfo.Name = e.Token0.Symbol + " - " + e.Token1.Symbol
		// TODO: cal total liquidity
		token0ETH, err := strconv.ParseFloat(e.Token0.DerivedETH, 64)
		token1ETH, err := strconv.ParseFloat(e.Token1.DerivedETH, 64)
		token0Liquidity, err := strconv.ParseFloat(e.Token0.TotalLiquidity, 64)
		token1Liquidity, err := strconv.ParseFloat(e.Token1.TotalLiquidity, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		pairInfo.TotalLiquidity = token0Liquidity*token0ETH*ethPrice + token1Liquidity*token1ETH*ethPrice

		listPairInfo = append(listPairInfo, pairInfo)
	}
	pageVars.ListPairInfo = listPairInfo

	Render(w, "home.html", pageVars)
}
