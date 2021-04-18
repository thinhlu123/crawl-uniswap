package page

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"thinhlu123/crawl-uniswap/src/model"
	"time"
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
	token0Per1, err := strconv.ParseFloat(pair.Token0Price, 64)
	token1Per0, err := strconv.ParseFloat(pair.Token1Price, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	pair.Token0Price = fmt.Sprintf("%.3f", token0Per1)
	pair.Token1Price = fmt.Sprintf("%.3f", token1Per0)
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
	pageVars.Token0USD = fmt.Sprintf("%.3f", token0ETH*ethPrice)
	pageVars.Token1USD = fmt.Sprintf("%.3f", token1ETH*ethPrice)

	// get total liquidity
	token0Liquidity, err := strconv.ParseFloat(pair.Token0.TotalLiquidity, 64)
	token1Liquidity, err := strconv.ParseFloat(pair.Token1.TotalLiquidity, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	pageVars.TotalLiquidity = fmt.Sprintf("%.3f", token0Liquidity*token0ETH*ethPrice+token1Liquidity*token1ETH*ethPrice)

	//// 2. Get list transaction
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(15)
	findOptions.SetSort(bson.D{{"time_transaction", -1}})
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

	now := time.Now()
	for i := 0; i < len(transaction); i++ {
		if transaction[i].Type == "SWAP" {
			if transaction[i].Amount0In != "0" {
				transaction[i].Type += " " + pair.Token1.Symbol + "-" + pair.Token0.Symbol
				amount0, err := strconv.ParseFloat(transaction[i].Amount1Out, 64)
				amount1, err := strconv.ParseFloat(transaction[i].Amount0In, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				transaction[i].Amount0 = fmt.Sprintf("%.3f ", amount0) + pair.Token1.Symbol
				transaction[i].Amount1 = fmt.Sprintf("%.3f ", amount1) + pair.Token0.Symbol
			} else {
				amount0, err := strconv.ParseFloat(transaction[i].Amount0Out, 64)
				amount1, err := strconv.ParseFloat(transaction[i].Amount1In, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				transaction[i].Type += " " + pair.Token0.Symbol + "-" + pair.Token1.Symbol
				transaction[i].Amount0 = fmt.Sprintf("%.3f ", amount0) + pair.Token0.Symbol
				transaction[i].Amount1 = fmt.Sprintf("%.3f ", amount1) + pair.Token1.Symbol
			}
		} else {
			amount0, err := strconv.ParseFloat(transaction[i].Amount0, 64)
			amount1, err := strconv.ParseFloat(transaction[i].Amount1, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			transaction[i].Amount0 = fmt.Sprintf("%.3f ", amount0) + pair.Token0.Symbol
			transaction[i].Amount1 = fmt.Sprintf("%.3f ", amount1) + pair.Token1.Symbol
		}

		// convert time to stirng
		tm := time.Unix(transaction[i].TimeTransaction, 0)
		diff := now.Sub(tm)
		out := time.Time{}.Add(diff)
		transaction[i].TransactionTimeString = out.Format("15 hours 04 minutes 05 seconds")
	}

	pageVars.Transaction = transaction

	Render(w, "detail.html", pageVars)
}
