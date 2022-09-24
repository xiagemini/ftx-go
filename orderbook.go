package ftx

import (
	"log"

	"github.com/xiagemini/ftx-go/structs"
)

type Orderbook structs.Orderbook

func (client *FtxClient) GetOrderBook(market string) (Orderbook, error) {
	var orderBook Orderbook
	resp, err := client._get("markets/"+market+"/orderbook?depth=5", []byte(""))
	if err != nil {
		log.Fatal(err)
		return orderBook, err
	}
	err = _processResponse(resp, &orderBook)
	return orderBook, err
}

func (client *FtxClient) GetBestPrice(market string) (map[string]float32, error) {
	var orderBook Orderbook
	resp, err := client._get("markets/"+market+"/orderbook?depth=1", []byte(""))
	if err != nil {
		log.Fatal(err)
		return map[string]float32{"ask": 0, "bid": 0}, err
	}
	err = _processResponse(resp, &orderBook)
	return map[string]float32{"ask": orderBook.Result.Asks[0][0], "bid": orderBook.Result.Bids[0][0]}, err
}
