package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var blockHeight int = 0

func checkExpire() {
	for {
		blockHeight++
		fmt.Printf("Set height to: " + strconv.Itoa(blockHeight))
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go checkExpire()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ETH Blockheight: %s!", strconv.Itoa(blockHeight))
	})
	http.ListenAndServe(":8080", nil)
}
