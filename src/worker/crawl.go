package worker

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"thinhlu123/crawl-uniswap/src/client"
	"thinhlu123/crawl-uniswap/src/model"
	"time"
)

// List 10 uniswap
// WETH - USDT, FEI - WETH, FEI - TRIBE, WBTC - WETH, USDC - WETH, RAI - DAI, DAI - HOPR, BOND - USDC, ORN - WETH, DAO - WETH
var pairArr = []string{"0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852", "0x94b0a3d511b6ecdb17ebf877278ab030acb0a878",
	"0x9928e4046d7c6513326ccea028cd3e7a91c7590a", "0xbb2b8038a1640196fbe3e38816f3e67cba72d940",
	"0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc", "0x4a0ea6ad985f6526de7d1ade562e1007e9c5d757",
	"0x92c2fc5f306405eab0ff0958f6d85d7f8892cf4d", "0x6591c4bcd6d7a1eb4e537da8b78676c1576ba244",
	"0x6c8b0dee9e90ea9f790da5daf6f5b20d23b39689", "0x7dd3f5705504002dc946aeafe6629b9481b72272",
}

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
		now := time.Now()
		newObjs := make([]interface{}, 0, len(transaction))
		for i := range transaction {
			newObj, err := model.ToBsonDoc(transaction[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			newObj["created_time"] = now

			timeTransaction, err := strconv.ParseInt(transaction[i].Timestamp, 10, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			newObj["time_transaction"] = timeTransaction
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
