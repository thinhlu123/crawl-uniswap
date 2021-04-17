package page

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
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
	for _, e := range out {
		var pairInfo model.PairInfo
		pairInfo.Name = e.Token0.Symbol + " - " + e.Token1.Symbol
		// TODO: cal total liquidity

		listPairInfo = append(listPairInfo, pairInfo)
	}
	pageVars.ListPairInfo = listPairInfo

	Render(w, "home.html", pageVars)
}
