package tron

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var _apikey = os.Getenv("API_KEY")

func TestGrid_GetContractTxsByTs(t *testing.T) {
	grid := getGrid()
	txs, err := grid.GetContractTxsByTs(context.Background(), from, 1698236022000, TO)
	if err != nil {
		panic(err)
	}

	fmt.Println(txs)
}

func TestGrid_GetAccountInfo(t *testing.T) {
	grid := getGrid()
	txs, err := grid.GetAccountInfo(context.Background(), from)
	if err != nil {
		panic(err)
	}

	fmt.Println(txs)
}

func TestGrid_GetUSDTRecentTxsByTs(t *testing.T) {
	grid := getGrid()
	txs, err := grid.GetContractRecentTxsByTs(context.Background(), "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs", from, 0, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(txs)
}

func getGrid() (grid *Grid) {
	client := NewHttpClient("https://api.shasta.trongrid.io")
	grid = NewGrid(client)
	grid.Apikey.Append(_apikey)
	return
}
