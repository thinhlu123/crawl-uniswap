package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"thinhlu123/crawl-uniswap/src/client"
	"thinhlu123/crawl-uniswap/src/model"
	"thinhlu123/crawl-uniswap/src/page"
)

func main() {
	// init db
	clientDB, err := model.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		// Close the connection once no longer needed
		err = clientDB.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection to MongoDB closed.")
		}
	}()

	// init client
	client.InitUniSwapClient()

	///// http
	// serve css folder
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// route home page
	http.HandleFunc("/", page.Home)

	// serve
	http.ListenAndServe(getPort(), nil)
}

//
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}
