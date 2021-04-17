package page

import (
	"net/http"
	"thinhlu123/crawl-uniswap/src/model"
)

func Detail(w http.ResponseWriter, req *http.Request) {
	pageVars := model.PageVars{
		Title: "Detail Uniswap",
	}

	// Get All Pair

	Render(w, "detail.html", pageVars)
}
