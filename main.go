package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/blockcypher/gobcy"
	"github.com/labstack/echo"
)

var heights = map[string]int{
	"eth":  0,
	"btc":  0,
	"ltc":  0,
	"doge": 0,
	"dash": 0,
}

func updateHeight(chain string) {
	bc := gobcy.API{"", chain, "main"}
	for {
		data, err := bc.GetChain()
		if err != nil {
			log.Fatal(err)
		}
		heights[chain] = data.Height
		fmt.Println("Set " + chain + " height to: " + strconv.Itoa(data.Height))
		time.Sleep(60 * time.Second)
	}
}

func handler(c echo.Context) error {
	chain := c.Param("chain")
	if _, ok := heights[chain]; ok {
		return c.String(http.StatusOK, strconv.Itoa(heights[chain]))
	}
	return c.String(http.StatusNotFound, "Invalid chain")
}

func main() {
	for k := range heights {
		go updateHeight(k)
	}
	e := echo.New()
	e.GET("/:chain", handler)
	e.Logger.Fatal(e.Start(":8080"))
}
