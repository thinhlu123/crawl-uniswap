package worker

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"thinhlu123/crawl-uniswap/src/client"
	"thinhlu123/crawl-uniswap/src/model"
)

/////// STEP
// Choose 10 pair -> save to db
// 1. Get ETH to USD
// 2. Get Coin to ETH
// 3. Compare coin use data in 1,2
// 4. Get 5 transaction of each pair

//// Get total totalLiquidity
/// get totalLiquidity of 2 token in pair and convert to usd

var pairArr = []string{"0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"}

func CrawlUniswap() {
	fmt.Println("Worker crawl start...")
	ctx := context.TODO()

	// Step 1: Get bundle
	bundle, err := client.GetPriceETH()
	if err != nil {
		fmt.Println(err)
		return
	}

	bundle.Type = "ETH_TO_USD"
	opts := options.Update().SetUpsert(true)
	filter := model.Bundle{Type: "ETH_TO_USD"}
	update := bson.M{"$set": bundle}
	_, err = model.BundleCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: Get Pair & Transaction info
	for _, pairId := range pairArr {

		// pair
		pair, err := client.GetPairById(pairId)
		if err != nil {
			fmt.Println(err)
			return
		}

		filterPair := bson.M{"pair_id": pairId}
		updatePair := bson.M{"$set": pair[0]}
		_, err = model.PairCollection.UpdateOne(ctx, filterPair, updatePair, opts)
		if err != nil {
			fmt.Println(err)
			return
		}

		/// TODO: calculate convert to coin

		// transaction
		transaction, err := client.GetTokenTransactionById(pairId)
		if err != nil {
			fmt.Println(err)
			return
		}
		//swapDoc, err := model.ToBsonDoc(swap)
		newObjs := make([]interface{}, 0, len(transaction))
		for i := range transaction {
			newObj, err := model.ToBsonDoc(transaction[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			newObjs = append(newObjs, newObj)
		}
		_, err = model.TransactionCollection.InsertMany(ctx, newObjs)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Worker crawl end...")
}
