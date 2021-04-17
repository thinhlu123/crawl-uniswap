package page

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"thinhlu123/crawl-uniswap/src/model"
)

func Detail(w http.ResponseWriter, req *http.Request) {
	pageVars := model.PageVars{
		Title: "Detail Uniswap",
	}

	req.ParseForm() //r is url.Values which is a map[string][]string
	var dvalues []string
	for _, values := range req.Form { // range over map
		for _, value := range values { // range over []string
			dvalues = append(dvalues, value) // stick each value in a slice I know the name of
		}
	}

	pairId := dvalues[0]

	////// Get info
	//// 1. Get pair info -> cal total liquidity
	var pair model.Pair
	err := model.PairCollection.FindOne(context.TODO(), bson.M{"pair_id": pairId}).Decode(&pair)
	if err != nil {
		fmt.Println(err)
		return
	}
	pageVars.Pair = pair

	// get token price in usb
	var bundle model.Bundle
	err = model.BundleCollection.FindOne(context.TODO(), bson.M{"type": "ETH_TO_USD"}).Decode(&bundle)
	if err != nil {
		fmt.Println(err)
		return
	}
	ethPrice, err := strconv.ParseFloat(bundle.ETHPrice, 64)
	token0ETH, err := strconv.ParseFloat(pair.Token0.DerivedETH, 64)
	token1ETH, err := strconv.ParseFloat(pair.Token1.DerivedETH, 64)

	if err != nil {
		fmt.Println(err)
		return
	}
	pageVars.Token0USD = token0ETH * ethPrice
	pageVars.Token1USD = token1ETH * ethPrice

	// get total liquidity
	token0Liquidity, err := strconv.ParseFloat(pair.Token0.TotalLiquidity, 64)
	token1Liquidity, err := strconv.ParseFloat(pair.Token1.TotalLiquidity, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	pageVars.TotalLiquidity = token0Liquidity*token0ETH*ethPrice + token1Liquidity*token1ETH*ethPrice

	//// 2. Get list transaction
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(15)
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	tran, err := model.TransactionCollection.Find(context.TODO(), bson.M{"pair_id": pairId}, findOptions)
	if err != nil {
		return
	}
	var transaction []model.TokenTransaction
	err = tran.All(context.TODO(), &transaction)
	if err != nil {
		return
	} else if len(transaction) == 0 {
		return
	}

	pageVars.Transaction = transaction

	Render(w, "detail.html", pageVars)
}
