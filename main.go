package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"thinhlu123/crawl-uniswap/src/client"
	"thinhlu123/crawl-uniswap/src/model"
	"thinhlu123/crawl-uniswap/src/page"
	"thinhlu123/crawl-uniswap/src/worker"
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

	// init worker
	var wg = sync.WaitGroup{}
	var workerCrawl worker.AppWorker
	workerCrawl.SetTask(worker.CrawlUniswap)
	workerCrawl.SetRepeatPeriod(60)
	// Do first time
	workerCrawl.Task()
	wg.Add(1)
	go workerCrawl.Execute()

	///// http
	// serve css folder
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// route home page
	http.HandleFunc("/", page.Home)
	http.HandleFunc("/detail", page.Detail)

	// serve
	port := getPort()
	fmt.Println("Serve at PORT" + port)
	http.ListenAndServe(port, nil)

	wg.Wait()
}

//
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}
