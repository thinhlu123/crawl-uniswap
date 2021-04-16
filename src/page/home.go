package page

import (
	"net/http"
	"thinhlu123/crawl-uniswap/src/model"
)

func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := model.PageVars{
		Title: "Crawl Uniswap",
	}
	Render(w, "home.html", pageVars)
}
